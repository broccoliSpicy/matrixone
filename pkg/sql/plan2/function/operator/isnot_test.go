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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either acosress or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package operator

import (
	"testing"

	"github.com/matrixorigin/matrixone/pkg/sql/testutil"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"github.com/stretchr/testify/require"
)

func TestIsNot(t *testing.T) {
	procs := testutil.NewProc()
	cases := []struct {
		name     string
		proc     *process.Process
		left     []bool
		nsp      []uint64
		right    bool
		isScalar bool
		expected []bool
	}{
		{
			name:     "01 - normal test",
			proc:     procs,
			left:     []bool{true, false, true, false, false, true},
			right:    true,
			expected: []bool{false, true, false, true, true, false},
			isScalar: false,
		},
		{
			name:     "02 - normal test",
			proc:     procs,
			left:     []bool{true, false, true, false, false, true},
			right:    false,
			expected: []bool{true, false, true, false, false, true},
			isScalar: false,
		},
		{
			name:     "03 - scalar test",
			proc:     procs,
			left:     []bool{true},
			right:    false,
			expected: []bool{true},
			isScalar: true,
		},
		{
			name:     "04 - null test",
			proc:     procs,
			right:    true,
			expected: []bool{false},
			isScalar: true,
		},
		{
			name:     "05 - null test",
			proc:     procs,
			right:    false,
			expected: []bool{false},
			isScalar: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			vecs := makeIsAndIsNotTestVectors(c.left, c.right, c.isScalar)
			result, err := IsNot(vecs, c.proc)
			if err != nil {
				t.Fatal(err)
			}
			col := result.Col.([]bool)
			require.Equal(t, c.expected, col)
			require.Equal(t, c.isScalar, result.IsScalar())
		})
	}
}
