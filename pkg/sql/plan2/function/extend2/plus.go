package extend2

import (
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/vectorize/add"
	"github.com/matrixorigin/matrixone/pkg/vm/mheap"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"golang.org/x/exp/constraints"
)

// this is just template code with only the purpose of discussion
func Plus[T constraints.Integer | constraints.Float](vectors []*vector.Vector, proc *process.Process) (*vector.Vector, error) {
	// should we differentiate the function signature for binary/unary/variadic operators?
	left, right := vectors[0], vectors[1]
	leftValues, rightValues := left.Col.([]T), right.Col.([]T)
	resultElementSize := left.Typ.Oid.FixedLength()
	switch {
	case left.IsConst && right.IsConst:
		// in the case where the result is a const, I chose to return only one row containing the const,
		// this may require corresponding changes in the frontend, is this okay?
		resultVector := vector.New(left.Typ)
		resultValues := make([]T, 1)
		nulls.Or(left.Nsp, right.Nsp, resultVector.Nsp)
		vector.SetCol(resultVector, add.NumericAdd[T](leftValues, rightValues, resultValues)) // if our input contains null, this step may be redundant,
		resultVector.IsConst = true
		resultVector.Length = left.Length
		return resultVector, nil
	case left.IsConst && !right.IsConst:
		data, err := mheap.Alloc(proc.Mp, int64(resultElementSize*len(rightValues)))
		if err != nil {
			return nil, err
		}
		resultVector := vector.New(left.Typ)
		resultVector.Data = data
		resultValues := encoding.DecodeFixedSlice[T](resultVector.Data, resultElementSize)
		// resultValues := resultValues[:len(rightValues)] // I think this line of code is redundant
		nulls.Or(left.Nsp, right.Nsp, resultVector.Nsp)
		vector.SetCol(resultVector, add.NumericAddScalar[T](leftValues[0], rightValues, resultValues))
		return resultVector, nil
	case !left.IsConst && right.IsConst:
		data, err := mheap.Alloc(proc.Mp, int64(resultElementSize*len(leftValues)))
		if err != nil {
			return nil, err
		}
		resultVector := vector.New(left.Typ)
		resultVector.Data = data
		resultValues := encoding.DecodeFixedSlice[T](resultVector.Data, resultElementSize)
		nulls.Or(left.Nsp, right.Nsp, resultVector.Nsp)
		vector.SetCol(resultVector, add.NumericAddScalar[T](rightValues[0], leftValues, resultValues))
		return resultVector, nil
	}
	data, err := mheap.Alloc(proc.Mp, int64(resultElementSize*len(rightValues)))
	if err != nil {
		return nil, err
	}
	resultVector := vector.New(left.Typ)
	resultVector.Data = data
	resultValues := encoding.DecodeFixedSlice[T](resultVector.Data, resultElementSize)
	nulls.Or(left.Nsp, right.Nsp, resultVector.Nsp)
	vector.SetCol(resultVector, add.NumericAdd[T](leftValues, rightValues, resultValues))
	return resultVector, nil
}
