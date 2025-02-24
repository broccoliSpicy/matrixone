package timestamp

import (
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/matrixorigin/matrixone/pkg/container/types"
)

func TestDateToTimestamp(t *testing.T) {
	testCases := []struct {
		name    string
		args    []types.Date
		want    []types.Timestamp
		success bool
	}{
		{
			name:    "normal test cases",
			args:    []types.Date{types.FromCalendar(2022, 3, 30)},
			want:    []types.Timestamp{types.FromClockUTC(2022, 3, 30, 0, 0, 0, 0)},
			success: true,
		},
	}

	for _, v := range testCases {
		reply := make([]types.Timestamp, len(v.args))
		ns := &nulls.Nulls{}
		reply = DateToTimestamp(v.args, ns, reply)
		require.Equal(t, reply, v.want)
		require.Equal(t, !nulls.Contains(ns, 0), v.success)
	}
}

func TestDatetimeToTimestamp(t *testing.T) {
	testCases := []struct {
		name    string
		args    []types.Datetime
		want    []types.Timestamp
		success bool
	}{
		{
			name:    "normal test cases",
			args:    []types.Datetime{types.FromClock(2022, 3, 30, 0, 0, 0, 0)},
			want:    []types.Timestamp{types.FromClockUTC(2022, 3, 30, 0, 0, 0, 0)},
			success: true,
		},
	}

	for _, v := range testCases {
		reply := make([]types.Timestamp, len(v.args))
		ns := &nulls.Nulls{}
		reply = DatetimeToTimestamp(v.args, ns, reply)
		require.Equal(t, reply, v.want)
		require.Equal(t, !nulls.Contains(ns, 0), v.success)
	}
}

func TestDateStringToTimestamp(t *testing.T) {
	testCases := []struct {
		name    string
		args    *types.Bytes
		want    []types.Timestamp
		success bool
	}{
		{
			name:    "normal test cases",
			args:    &types.Bytes{Data: []byte("2022-03-30 00:00:00"), Offsets: []uint32{0}, Lengths: []uint32{19}},
			want:    []types.Timestamp{types.FromClockUTC(2022, 3, 30, 0, 0, 0, 0)},
			success: true,
		},
	}

	for _, v := range testCases {
		reply := make([]types.Timestamp, len(v.args.Lengths))
		ns := &nulls.Nulls{}
		reply = DateStringToTimestamp(v.args, ns, reply)
		require.Equal(t, reply, v.want)
		require.Equal(t, !nulls.Contains(ns, 0), v.success)
	}
}
