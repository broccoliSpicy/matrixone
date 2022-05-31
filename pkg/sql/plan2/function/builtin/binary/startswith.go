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
	leftValues, rightValues := left.Col.(*types.Bytes), right.Col.(*types.Bytes)
	resultType := types.Type{Oid: types.T_uint8, Size: 1}
	resultElementSize := int(resultType.Size)
	switch {
	case left.IsConst && right.IsConst:
		if left.ConstVectorIsNull() || right.ConstVectorIsNull() {
			return proc.AllocScalarNullVector(resultType), nil
		}
		resultVector := vector.NewConst(left.Typ)
		resultValues := make([]uint8, 1)
		vector.SetCol(resultVector, startswith.StartsWithAllConst(leftValues, rightValues, resultValues))
		return resultVector, nil
	case left.IsConst && !right.IsConst:
		if left.ConstVectorIsNull() {
			return proc.AllocScalarNullVector(left.Typ), nil
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
			return proc.AllocScalarNullVector(left.Typ), nil
		}
		resultVector, err := proc.AllocVector(left.Typ, int64(resultElementSize*len(leftValues.Lengths)))
		if err != nil {
			return nil, err
		}
		resultValues := encoding.DecodeUint8Slice(resultVector.Data)
		resultValues = resultValues[:len(leftValues.Lengths)]
		nulls.Set(resultVector.Nsp, left.Nsp)
		vector.SetCol(resultVector, startswith.StartsWithRightConst(leftValues, rightValues, resultValues))
		return resultVector, nil
	}
	resultVector, err := proc.AllocVector(left.Typ, int64(resultElementSize*len(rightValues.Lengths)))
	if err != nil {
		return nil, err
	}
	resultValues := encoding.DecodeUint8Slice(resultVector.Data)
	resultValues = resultValues[:len(rightValues.Lengths)]
	nulls.Or(left.Nsp, right.Nsp, resultVector.Nsp)
	vector.SetCol(resultVector, startswith.StartsWith(leftValues, rightValues, resultValues))
	return resultVector, nil
}
