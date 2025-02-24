// Copyright 2022 Matrix Origin
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

package binary

import (
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/vectorize/startswith"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func Startswith(vectors []*vector.Vector, proc *process.Process) (*vector.Vector, error) {
	left, right := vectors[0], vectors[1]
	resultType := types.Type{Oid: types.T_uint8, Size: 1}
	resultElementSize := int(resultType.Size)
	leftValues, rightValues := vector.MustBytesCols(left), vector.MustBytesCols(right)
	switch {
	case left.IsScalar() && right.IsScalar():
		if left.ConstVectorIsNull() || right.ConstVectorIsNull() {
			return proc.AllocScalarNullVector(resultType), nil
		}
		resultVector := vector.NewConst(resultType)
		resultValues := make([]uint8, 1)
		vector.SetCol(resultVector, startswith.StartsWithAllConst(leftValues, rightValues, resultValues))
		return resultVector, nil
	case left.IsConst && !right.IsConst:
		if left.ConstVectorIsNull() {
			return proc.AllocScalarNullVector(resultType), nil
		}
		resultVector, err := proc.AllocVector(resultType, int64(resultElementSize*len(rightValues.Lengths)))
		if err != nil {
			return nil, err
		}
		resultValues := encoding.DecodeUint8Slice(resultVector.Data)
		resultValues = resultValues[:len(rightValues.Lengths)]
		nulls.Set(resultVector.Nsp, right.Nsp)
		vector.SetCol(resultVector, startswith.StartsWithLeftConst(leftValues, rightValues, resultValues))
		return resultVector, nil
	case !left.IsConst && right.IsConst:
		if right.ConstVectorIsNull() {
			return proc.AllocScalarNullVector(resultType), nil
		}
		resultVector, err := proc.AllocVector(resultType, int64(resultElementSize*len(leftValues.Lengths)))
		if err != nil {
			return nil, err
		}
		resultValues := encoding.DecodeUint8Slice(resultVector.Data)
		resultValues = resultValues[:len(leftValues.Lengths)]
		nulls.Set(resultVector.Nsp, left.Nsp)
		vector.SetCol(resultVector, startswith.StartsWithRightConst(leftValues, rightValues, resultValues))
		return resultVector, nil
	}
	resultVector, err := proc.AllocVector(resultType, int64(resultElementSize*len(rightValues.Lengths)))
	if err != nil {
		return nil, err
	}
	resultValues := encoding.DecodeUint8Slice(resultVector.Data)
	resultValues = resultValues[:len(rightValues.Lengths)]
	nulls.Or(left.Nsp, right.Nsp, resultVector.Nsp)
	vector.SetCol(resultVector, startswith.StartsWith(leftValues, rightValues, resultValues))
	return resultVector, nil
}
