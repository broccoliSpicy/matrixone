// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/matrixorigin/matrixone/pkg/vm/engine/tae/common"
	"github.com/matrixorigin/matrixone/pkg/vm/engine/tae/logstore/entry"
)

type syncBase struct {
	*sync.RWMutex
	groupLSN                     map[uint32]uint64 // for alloc
	lsnmu                        sync.RWMutex
	checkpointing                map[uint32]*checkpointInfo
	ckpmu                        sync.RWMutex
	syncing                      map[uint32]uint64
	checkpointed, synced, ckpCnt *syncMap
	uncommits                    map[uint32][]uint64
	addrs                        map[uint32]map[int]common.ClosedIntervals //group-version-glsn range
	addrmu                       sync.RWMutex
}

type checkpointInfo struct {
	ranges  *common.ClosedIntervals
	partial map[uint64]*partialCkpInfo
}

func newCheckpointInfo() *checkpointInfo {
	return &checkpointInfo{
		ranges:  common.NewClosedIntervals(),
		partial: make(map[uint64]*partialCkpInfo),
	}
}

func (info *checkpointInfo) UpdateWtihRanges(intervals *common.ClosedIntervals) {
	info.ranges.TryMerge(*intervals)
}

func (info *checkpointInfo) UpdateWtihPartialCheckpoint(lsn uint64, ckps *partialCkpInfo) {
	if info.ranges.Contains(*common.NewClosedIntervalsByInt(lsn)) {
		return
	}
	partialInfo, ok := info.partial[lsn]
	if !ok {
		partialInfo = newPartialCkpInfo(ckps.size)
		info.partial[lsn] = partialInfo
	}
	partialInfo.MergePartialCkpInfo(ckps)
	if partialInfo.IsAllCheckpointed() {
		info.ranges.TryMerge(*common.NewClosedIntervalsByInt(lsn))
		delete(info.partial, lsn)
	}
}

func (info *checkpointInfo) UpdateWithCommandInfo(lsn uint64, cmds *entry.CommandInfo) {
	if info.ranges.Contains(*common.NewClosedIntervalsByInt(lsn)) {
		return
	}
	partialInfo, ok := info.partial[lsn]
	if !ok {
		partialInfo = newPartialCkpInfo(cmds.Size)
		info.partial[lsn] = partialInfo
	}
	partialInfo.MergeCommandInfos(cmds)
	if partialInfo.IsAllCheckpointed() {
		info.ranges.TryMerge(*common.NewClosedIntervalsByInt(lsn))
		delete(info.partial, lsn)
	}
}
func (info *checkpointInfo) MergeCheckpointInfo(ockp *checkpointInfo) {
	info.ranges.TryMerge(*ockp.ranges)
	for lsn, ockpinfo := range ockp.partial {
		ckpinfo, ok := info.partial[lsn]
		if !ok {
			info.partial[lsn] = ockpinfo
		} else {
			if ckpinfo.size != ockpinfo.size {
				panic("logic err")
			}
			ckpinfo.ckps.Or(ockpinfo.ckps)
		}
	}
}
func (info *checkpointInfo) GetCheckpointed() uint64 {
	if info.ranges == nil || len(info.ranges.Intervals) == 0 {
		return 0
	}
	if info.ranges.Intervals[0].Start != 1 {
		return 0
	}
	return info.ranges.Intervals[0].End
}

func (info *checkpointInfo) String() string {
	s := fmt.Sprintf("range %v, partial ", info.ranges)
	for lsn, partial := range info.partial {
		s = fmt.Sprintf("%s[%d-%v]", s, lsn, partial)
	}
	return s
}

func (info *checkpointInfo) GetCkpCnt() uint64 {
	cnt := uint64(0)
	cnt += uint64(info.ranges.GetCardinality())
	// cnt += uint64(len(info.partial))
	return cnt
}

func (info *checkpointInfo) WriteTo(w io.Writer) (n int64, err error) {
	sn, err := info.ranges.WriteTo(w)
	n += sn
	if err != nil {
		return
	}
	length := uint64(len(info.partial))
	if err = binary.Write(w, binary.BigEndian, length); err != nil {
		return
	}
	n += 8
	for lsn, partialInfo := range info.partial {
		if err = binary.Write(w, binary.BigEndian, lsn); err != nil {
			return
		}
		n += 8
		sn, err = partialInfo.WriteTo(w)
		n += sn
		if err != nil {
			return
		}
	}
	return
}

func (info *checkpointInfo) ReadFrom(r io.Reader) (n int64, err error) {
	info.ranges = common.NewClosedIntervals()
	sn, err := info.ranges.ReadFrom(r)
	n += sn
	if err != nil {
		return
	}
	length := uint64(0)
	if err = binary.Read(r, binary.BigEndian, &length); err != nil {
		return
	}
	n += 8
	for i := 0; i < int(length); i++ {
		lsn := uint64(0)
		if err = binary.Read(r, binary.BigEndian, &lsn); err != nil {
			return
		}
		n += 8
		partial := newPartialCkpInfo(0)
		sn, err = partial.ReadFrom(r)
		n += sn
		if err != nil {
			return
		}
		info.partial[lsn] = partial
	}
	return
}

type syncMap struct {
	*sync.RWMutex
	ids map[uint32]uint64
}

func newSyncMap() *syncMap {
	return &syncMap{
		RWMutex: new(sync.RWMutex),
		ids:     make(map[uint32]uint64),
	}
}
func newSyncBase() *syncBase {
	return &syncBase{
		groupLSN:      make(map[uint32]uint64),
		lsnmu:         sync.RWMutex{},
		checkpointing: make(map[uint32]*checkpointInfo),
		syncing:       make(map[uint32]uint64),
		checkpointed:  newSyncMap(),
		synced:        newSyncMap(),
		ckpCnt:        newSyncMap(),
		uncommits:     make(map[uint32][]uint64),
		addrs:         make(map[uint32]map[int]common.ClosedIntervals),
		addrmu:        sync.RWMutex{},
		ckpmu:         sync.RWMutex{},
	}
}

func (base *syncBase) WritePostCommitEntry(w io.Writer) (n int64, err error) {
	base.ckpmu.RLock()
	defer base.ckpmu.RUnlock()
	//checkpointing
	length := uint32(len(base.checkpointing))
	if err = binary.Write(w, binary.BigEndian, length); err != nil {
		return
	}
	n += 4
	for groupID, ckpInfo := range base.checkpointing {
		if err = binary.Write(w, binary.BigEndian, groupID); err != nil {
			return
		}
		n += 4
		sn, err := ckpInfo.WriteTo(w)
		n += sn
		if err != nil {
			return n, err
		}
	}
	return
}

func (base *syncBase) ReadPostCommitEntry(r io.Reader) (n int64, err error) {
	//checkpointing
	length := uint32(0)
	if err = binary.Read(r, binary.BigEndian, &length); err != nil {
		return
	}
	n += 4
	for i := 0; i < int(length); i++ {
		groupID := uint32(0)
		if err = binary.Read(r, binary.BigEndian, &groupID); err != nil {
			return
		}
		n += 4
		ckpInfo := newCheckpointInfo()
		sn, err := ckpInfo.ReadFrom(r)
		n += sn
		if err != nil {
			return n, err
		}
	}
	return
}
func (base *syncBase) MarshalPostCommitEntry() (buf []byte, err error) {
	var bbuf bytes.Buffer
	if _, err = base.WritePostCommitEntry(&bbuf); err != nil {
		return
	}
	buf = bbuf.Bytes()
	return
}

func (base *syncBase) UnarshalPostCommitEntry(buf []byte) error {
	bbuf := bytes.NewBuffer(buf)
	_, err := base.ReadPostCommitEntry(bbuf)
	return err
}

func (base *syncBase) MakePostCommitEntry(id int) entry.Entry {
	e := entry.GetBase()
	e.SetType(entry.ETPostCommit)
	buf, err := base.MarshalPostCommitEntry()
	if err != nil {
		panic(err)
	}
	err = e.Unmarshal(buf)
	if err != nil {
		panic(err)
	}
	info := &entry.Info{}
	info.PostCommitVersion = id
	info.Group = entry.GTInternal
	e.SetInfo(info)
	return e
}

func (base *syncBase) OnReplay(r *replayer) {
	base.addrs = r.addrs
	base.groupLSN = r.groupLSN
	for k, v := range r.groupLSN {
		base.synced.ids[k] = v
	}
	if r.ckpEntry != nil {
		err := base.UnarshalPostCommitEntry(r.ckpEntry.payload)
		if err != nil {
			panic(err)
		}
	}
	for groupId, ckps := range r.checkpointrange {
		ckpInfo, ok := base.checkpointing[groupId]
		if !ok {
			base.checkpointing[groupId] = ckps
		} else {
			ckpInfo.MergeCheckpointInfo(ckps)
		}
		base.checkpointed.ids[groupId] = base.checkpointing[groupId].GetCheckpointed()
	}
}
func (base *syncBase) GetVersionByGLSN(groupId uint32, lsn uint64) (int, error) {
	base.addrmu.RLock()
	defer base.addrmu.RUnlock()
	versionsMap, ok := base.addrs[groupId]
	if !ok {
		return 0, errors.New("group not existed")
	}
	for ver, interval := range versionsMap {
		if interval.Contains(*common.NewClosedIntervalsByInt(lsn)) {
			return ver, nil
		}
	}
	fmt.Printf("versionsMap is %v\n", versionsMap)
	return 0, errors.New("lsn not existed")
}

//TODO
func (base *syncBase) GetLastAddr(groupName uint32, tid uint64) *VFileAddress {
	// tidMap, ok := base.uncommits[groupName]
	// if !ok {
	// 	return nil
	// }
	return nil
}

func (base *syncBase) OnEntryReceived(v *entry.Info) error {
	switch v.Group {
	case entry.GTCKp:
		for _, intervals := range v.Checkpoints {
			base.ckpmu.Lock()
			ckpInfo, ok := base.checkpointing[intervals.Group]
			if !ok {
				ckpInfo = newCheckpointInfo()
				base.checkpointing[intervals.Group] = ckpInfo
			}
			if intervals.Ranges != nil && len(intervals.Ranges.Intervals) > 0 {
				ckpInfo.UpdateWtihRanges(intervals.Ranges)
			}
			if intervals.Command != nil {
				for lsn, cmds := range intervals.Command {
					ckpInfo.UpdateWithCommandInfo(lsn, &cmds)
				}
			}
			base.ckpmu.Unlock()
		}
	case entry.GTUncommit:
		// addr := v.Addr.(*VFileAddress)
		for _, tid := range v.Uncommits {
			tids, ok := base.uncommits[tid.Group]
			if !ok {
				tids = make([]uint64, 0)
			}
			existed := false
			for _, id := range tids {
				if id == tid.Tid {
					existed = true
					break
				}
			}
			if !existed {
				tids = append(tids, tid.Tid)
			}
			base.uncommits[tid.Group] = tids
		}
		// fmt.Printf("receive uncommit %d-%d\n", v.Group, v.GroupLSN)
	default:
	}
	base.syncing[v.Group] = v.GroupLSN
	base.addrmu.Lock()
	defer base.addrmu.Unlock()
	addr := v.Info.(*VFileAddress)
	versionRanges, ok := base.addrs[addr.Group]
	if !ok {
		versionRanges = make(map[int]common.ClosedIntervals)
	}
	interval, ok := versionRanges[addr.Version]
	if !ok {
		interval = *common.NewClosedIntervals()
	}
	interval.TryMerge(*common.NewClosedIntervalsByInt(addr.LSN))
	versionRanges[addr.Version] = interval
	base.addrs[addr.Group] = versionRanges
	return nil
}

func (base *syncBase) GetPenddingCnt(groupId uint32) uint64 {
	ckp := base.GetCKpCnt(groupId)
	commit := base.GetSynced(groupId)
	return commit - ckp
}

func (base *syncBase) GetCheckpointed(groupId uint32) uint64 {
	base.checkpointed.RLock()
	defer base.checkpointed.RUnlock()
	return base.checkpointed.ids[groupId]
}

func (base *syncBase) SetCheckpointed(groupId uint32, id uint64) {
	base.checkpointed.Lock()
	base.checkpointed.ids[groupId] = id
	base.checkpointed.Unlock()
}

func (base *syncBase) GetSynced(groupId uint32) uint64 {
	base.synced.RLock()
	defer base.synced.RUnlock()
	return base.synced.ids[groupId]
}

func (base *syncBase) SetSynced(groupId uint32, id uint64) {
	base.synced.Lock()
	base.synced.ids[groupId] = id
	base.synced.Unlock()
}

func (base *syncBase) GetCKpCnt(groupId uint32) uint64 {
	base.ckpCnt.RLock()
	defer base.ckpCnt.RUnlock()
	return base.ckpCnt.ids[groupId]
}

func (base *syncBase) SetCKpCnt(groupId uint32, id uint64) {
	base.ckpCnt.Lock()
	base.ckpCnt.ids[groupId] = id
	base.ckpCnt.Unlock()
}

func (base *syncBase) OnCommit() {
	for group, checkpointing := range base.checkpointing {
		checkpointingId := checkpointing.GetCheckpointed()
		ckpcnt := checkpointing.GetCkpCnt()
		// logutil.Infof("G%d-%v",group,checkpointing)
		checkpointedId := base.GetCheckpointed(group)
		if checkpointingId > checkpointedId {
			base.SetCheckpointed(group, checkpointingId)
		}
		preCnt := base.GetCKpCnt(group)
		if ckpcnt > preCnt {
			base.SetCKpCnt(group, ckpcnt)
		}
	}

	for group, syncingId := range base.syncing {
		syncedId := base.GetSynced(group)
		if syncingId > syncedId {
			base.SetSynced(group, syncingId)
		}
	}
}

func (base *syncBase) AllocateLsn(groupID uint32) uint64 {
	base.lsnmu.Lock()
	defer base.lsnmu.Unlock()
	lsn, ok := base.groupLSN[groupID]
	if !ok {
		base.groupLSN[groupID] = 1
		return 1
	}
	lsn++
	base.groupLSN[groupID] = lsn
	return lsn
}

func (base *syncBase) GetCurrSeqNum(groupID uint32) uint64 {
	base.lsnmu.RLock()
	defer base.lsnmu.RUnlock()
	lsn, ok := base.groupLSN[groupID]
	if !ok {
		return 0
	}
	return lsn
}
