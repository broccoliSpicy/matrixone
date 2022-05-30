package multifunc

import (
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/vectorize/pi"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func Pi(_ []*vector.Vector, _ *process.Process) (*vector.Vector, error) {
	resultType := types.Type{Oid: types.T_float64, Size: 8}
	resultVector := vector.NewConst(resultType)
	result := make([]float64, 1)
	result[0] = pi.GetPi()
	vector.SetCol(resultVector, result)
	return resultVector, nil
}
