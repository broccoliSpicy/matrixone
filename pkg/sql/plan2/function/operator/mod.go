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

package operator

import (
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/errno"
	"github.com/matrixorigin/matrixone/pkg/sql/errors"
	"github.com/matrixorigin/matrixone/pkg/vectorize/mod"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"golang.org/x/exp/constraints"
)

// ErrModByZero is reported when computing the rest of a division by zero.
var ErrModByZero = errors.New(errno.SyntaxErrororAccessRuleViolation, "zero modulus")

func ModInt[T constraints.Integer](vectors []*vector.Vector, proc *process.Process) (*vector.Vector, error) {
	lv, rv := vectors[0], vectors[1]
	lvs, rvs := vector.MustTCols[T](lv), vector.MustTCols[T](rv)
	rtl := lv.Typ.Oid.FixedLength()

	if lv.IsScalarNull() || rv.IsScalarNull() {
		return proc.AllocScalarNullVector(lv.Typ), nil
	}

	switch {
	case lv.IsScalar() && rv.IsScalar():
		vec := proc.AllocScalarVector(lv.Typ)
		rs := make([]T, 1)
		nulls.Or(lv.Nsp, rv.Nsp, vec.Nsp)
		if !nulls.Any(rv.Nsp) {
			for _, v := range rvs {
				if v == 0 {
					return nil, ErrModByZero
				}
			}
			vector.SetCol(vec, mod.IntMod[T](lvs, rvs, rs))
			return vec, nil
		}
		sels := process.GetSels(proc)
		defer process.PutSels(sels, proc)
		for i, j := uint64(0), uint64(len(rvs)); i < j; i++ {
			if nulls.Contains(rv.Nsp, i) {
				continue
			}
			if rvs[i] == 0 {
				return nil, ErrModByZero
			}
			sels = append(sels, int64(i))
		}
		vector.SetCol(vec, mod.IntModSels[T](lvs, rvs, rs, sels))
		return vec, nil
	case lv.IsScalar() && !rv.IsScalar():
		if !nulls.Any(rv.Nsp) {
			for _, v := range rvs {
				if v == 0 {
					return nil, ErrModByZero
				}
			}
			vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(rvs)))
			if err != nil {
				return nil, err
			}
			rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
			nulls.Set(vec.Nsp, rv.Nsp)
			vector.SetCol(vec, mod.IntModScalar[T](lvs[0], rvs, rs))
			return vec, nil
		}
		sels := process.GetSels(proc)
		defer process.PutSels(sels, proc)
		for i, j := uint64(0), uint64(len(rvs)); i < j; i++ {
			if nulls.Contains(rv.Nsp, i) {
				continue
			}
			if rvs[i] == 0 {
				return nil, ErrModByZero
			}
			sels = append(sels, int64(i))
		}
		vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(rvs)))
		if err != nil {
			return nil, err
		}
		rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
		nulls.Set(vec.Nsp, rv.Nsp)
		vector.SetCol(vec, mod.IntModByScalarSels[T](lvs[0], rvs, rs, sels))
		return vec, nil
	case !lv.IsScalar() && rv.IsScalar():
		if rvs[0] == 0 {
			return nil, ErrModByZero
		}
		vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(lvs)))
		if err != nil {
			return nil, err
		}
		rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
		nulls.Set(vec.Nsp, lv.Nsp)
		vector.SetCol(vec, mod.IntModByScalar[T](rvs[0], lvs, rs))
		return vec, nil
	}
	vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(lvs)))
	if err != nil {
		return nil, err
	}
	rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
	nulls.Or(lv.Nsp, rv.Nsp, vec.Nsp)
	if !nulls.Any(rv.Nsp) {
		for _, v := range rvs {
			if v == 0 {
				return nil, ErrModByZero
			}
		}
		vector.SetCol(vec, mod.IntMod[T](lvs, rvs, rs))
		return vec, nil
	}
	sels := process.GetSels(proc)
	defer process.PutSels(sels, proc)
	for i, j := uint64(0), uint64(len(rvs)); i < j; i++ {
		if nulls.Contains(rv.Nsp, i) {
			continue
		}
		if rvs[i] == 0 {
			return nil, ErrModByZero
		}
		sels = append(sels, int64(i))
	}
	vector.SetCol(vec, mod.IntModSels[T](lvs, rvs, rs, sels))
	return vec, nil
}

func ModFloat[T constraints.Float](vectors []*vector.Vector, proc *process.Process) (*vector.Vector, error) {
	lv, rv := vectors[0], vectors[1]
	lvs, rvs := vector.MustTCols[T](lv), vector.MustTCols[T](rv)
	rtl := lv.Typ.Oid.FixedLength()

	if lv.IsScalarNull() || rv.IsScalarNull() {
		return proc.AllocScalarNullVector(lv.Typ), nil
	}

	switch {
	case lv.IsScalar() && rv.IsScalar():
		vec := proc.AllocScalarVector(lv.Typ)
		rs := make([]T, 1)
		nulls.Or(lv.Nsp, rv.Nsp, vec.Nsp)
		if !nulls.Any(rv.Nsp) {
			for _, v := range rvs {
				if v == 0 {
					return nil, ErrModByZero
				}
			}
			vector.SetCol(vec, mod.FloatMod[T](lvs, rvs, rs))
			return vec, nil
		}
		sels := process.GetSels(proc)
		defer process.PutSels(sels, proc)
		for i, j := uint64(0), uint64(len(rvs)); i < j; i++ {
			if nulls.Contains(rv.Nsp, i) {
				continue
			}
			if rvs[i] == 0 {
				return nil, ErrModByZero
			}
			sels = append(sels, int64(i))
		}
		vector.SetCol(vec, mod.FloatModSels[T](lvs, rvs, rs, sels))
		return vec, nil
	case lv.IsScalar() && !rv.IsScalar():
		if !nulls.Any(rv.Nsp) {
			for _, v := range rvs {
				if v == 0 {
					return nil, ErrModByZero
				}
			}
			vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(rvs)))
			if err != nil {
				return nil, err
			}
			rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
			nulls.Set(vec.Nsp, rv.Nsp)
			vector.SetCol(vec, mod.FloatModScalar[T](lvs[0], rvs, rs))
			return vec, nil
		}
		sels := process.GetSels(proc)
		defer process.PutSels(sels, proc)
		for i, j := uint64(0), uint64(len(rvs)); i < j; i++ {
			if nulls.Contains(rv.Nsp, i) {
				continue
			}
			if rvs[i] == 0 {
				return nil, ErrModByZero
			}
			sels = append(sels, int64(i))
		}
		vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(rvs)))
		if err != nil {
			return nil, err
		}
		rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
		nulls.Set(vec.Nsp, rv.Nsp)
		vector.SetCol(vec, mod.FloatModByScalarSels[T](lvs[0], rvs, rs, sels))
		return vec, nil
	case !lv.IsScalar() && rv.IsScalar():
		if rvs[0] == 0 {
			return nil, ErrModByZero
		}
		vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(lvs)))
		if err != nil {
			return nil, err
		}
		rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
		nulls.Set(vec.Nsp, lv.Nsp)
		vector.SetCol(vec, mod.FloatModByScalar[T](rvs[0], lvs, rs))
		return vec, nil
	}
	vec, err := proc.AllocVector(lv.Typ, int64(rtl)*int64(len(lvs)))
	if err != nil {
		return nil, err
	}
	rs := encoding.DecodeFixedSlice[T](vec.Data, rtl)
	nulls.Or(lv.Nsp, rv.Nsp, vec.Nsp)
	if !nulls.Any(rv.Nsp) {
		for _, v := range rvs {
			if v == 0 {
				return nil, ErrModByZero
			}
		}
		vector.SetCol(vec, mod.FloatMod[T](lvs, rvs, rs))
		return vec, nil
	}
	sels := process.GetSels(proc)
	defer process.PutSels(sels, proc)
	for i, j := uint64(0), uint64(len(rvs)); i < j; i++ {
		if nulls.Contains(rv.Nsp, i) {
			continue
		}
		if rvs[i] == 0 {
			return nil, ErrModByZero
		}
		sels = append(sels, int64(i))
	}
	vector.SetCol(vec, mod.FloatModSels[T](lvs, rvs, rs, sels))
	return vec, nil
}
