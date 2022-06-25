// Copyright 2021 - 2022 Matrix Origin
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

package function

import (
	"github.com/matrixorigin/matrixone/pkg/container/types"
)

const (
	maxTypeNumber = types.T_tuple + 10

	tooManyFunctionsMatched = -2
	wrongFunctionParameters = -1

	matchedDirectly = iota
	matchedByConvert
	matchedFailed
)

type binaryTargetTypes struct {
	convert     bool
	left, right types.T
}

// castTable indicates whether a type can be automatically converted to another type
var castTable [][]bool

// binaryTable is a cast rule table for some binary-operators' parameters
// e.g. PLUS, MINUS, GT and so on.
// Format is `binaryTable[LeftInput][RightInput] = {LeftTarget, RightTarget}`
var binaryTable [][]binaryTargetTypes

// binaryTable2 is a cast rule table for DIV and INTEGER_DIV
// Format is `binaryTable[LeftInput][RightInput] = {LeftTarget, RightTarget}`
var binaryTable2 [][]binaryTargetTypes

// init binaryTable and castTable
func init() {
	all := []types.T{
		types.T_any,
		types.T_bool,
		types.T_int8, types.T_int16, types.T_int32, types.T_int64,
		types.T_uint8, types.T_uint16, types.T_uint32, types.T_uint64,
		types.T_float32, types.T_float64,
		types.T_date, types.T_datetime, types.T_timestamp,
		types.T_char, types.T_varchar,
		types.T_decimal64, types.T_decimal128,
	}
	numbers := []types.T{ // numbers without decimal
		types.T_int8, types.T_int16, types.T_int32, types.T_int64,
		types.T_uint8, types.T_uint16, types.T_uint32, types.T_uint64,
	}
	ints := []types.T{types.T_int8, types.T_int16, types.T_int32, types.T_int64}
	uints := []types.T{types.T_uint8, types.T_uint16, types.T_uint32, types.T_uint64}
	floats := []types.T{types.T_float32, types.T_float64}
	strings := []types.T{types.T_char, types.T_varchar}
	decimals := []types.T{types.T_decimal64, types.T_decimal128}

	maxTypes := 255

	// init binaryTable
	var rules [][4]types.T // left-input, right-input, left-target, right-target
	{
		for _, typ1 := range numbers {
			for _, typ2 := range floats {
				rules = append(rules, [4]types.T{typ1, typ2, types.T_float64, types.T_float64})
				rules = append(rules, [4]types.T{typ2, typ1, types.T_float64, types.T_float64})
			}
		}
		for i := 0; i < len(ints); i++ {
			for j := i + 1; j < len(ints); j++ {
				rules = append(rules, [4]types.T{ints[i], ints[j], ints[j], ints[j]})
				rules = append(rules, [4]types.T{ints[j], ints[i], ints[j], ints[j]})
			}
		}
		for i := 0; i < len(uints); i++ {
			for j := i + 1; j < len(uints); j++ {
				rules = append(rules, [4]types.T{uints[i], uints[j], uints[j], uints[j]})
				rules = append(rules, [4]types.T{uints[j], uints[i], uints[j], uints[j]})
			}
		}
		rules = append(rules, [4]types.T{types.T_float32, types.T_float64, types.T_float64, types.T_float64})
		rules = append(rules, [4]types.T{types.T_float64, types.T_float32, types.T_float64, types.T_float64})
		rules = append(rules, [4]types.T{types.T_decimal64, types.T_decimal128, types.T_decimal128, types.T_decimal128})
		rules = append(rules, [4]types.T{types.T_decimal128, types.T_decimal64, types.T_decimal128, types.T_decimal128})
		for _, typ1 := range decimals {
			for _, typ2 := range numbers {
				rules = append(rules, [4]types.T{typ1, typ2, typ1, typ1})
				rules = append(rules, [4]types.T{typ2, typ1, typ1, typ1})
			}
			for _, typ2 := range floats {
				rules = append(rules, [4]types.T{typ1, typ2, types.T_float64, types.T_float64})
				rules = append(rules, [4]types.T{typ2, typ1, types.T_float64, types.T_float64})
			}
		}
		for i := 0; i < len(ints)-1; i++ {
			for j := 0; j < len(uints)-1; j++ {
				rules = append(rules, [4]types.T{ints[i], uints[j], ints[i+1], ints[i+1]})
				rules = append(rules, [4]types.T{uints[j], ints[i], ints[i+1], ints[i+1]})
			}
		}
		for i := range ints {
			rules = append(rules, [4]types.T{ints[i], types.T_uint64, types.T_int64, types.T_int64})
			rules = append(rules, [4]types.T{types.T_uint64, ints[i], types.T_int64, types.T_int64})
		}
		for i := range uints {
			rules = append(rules, [4]types.T{uints[i], types.T_int64, types.T_int64, types.T_int64})
			rules = append(rules, [4]types.T{types.T_int64, uints[i], types.T_int64, types.T_int64})
		}
		rules = append(rules, [4]types.T{types.T_date, types.T_datetime, types.T_datetime, types.T_datetime})
		rules = append(rules, [4]types.T{types.T_datetime, types.T_date, types.T_datetime, types.T_datetime})
		for _, t1 := range strings {
			for _, t2 := range all {
				if t1 == t2 || t2 == types.T_any {
					continue
				}
				rules = append(rules, [4]types.T{t1, t2, t2, t2})
				rules = append(rules, [4]types.T{t2, t1, t2, t2})
			}
		}
	}

	binaryTable = make([][]binaryTargetTypes, maxTypes)
	for i := range binaryTable {
		binaryTable[i] = make([]binaryTargetTypes, maxTypes)
	}
	for _, r := range rules {
		binaryTable[r[0]][r[1]] = binaryTargetTypes{
			convert: true,
			left:    r[2],
			right:   r[3],
		}
	}

	// init binaryTable2
	var rules2 [][4]types.T
	{
		for i := range numbers {
			for j := range numbers {
				rules2 = append(rules2, [4]types.T{numbers[i], numbers[j], types.T_float64, types.T_float64})
			}
			for j := range decimals {
				rules2 = append(rules2, [4]types.T{numbers[i], decimals[j], types.T_float64, types.T_float64})
				rules2 = append(rules2, [4]types.T{decimals[j], numbers[i], types.T_float64, types.T_float64})
			}
			for j := range floats {
				rules2 = append(rules2, [4]types.T{numbers[i], floats[j], types.T_float64, types.T_float64})
				rules2 = append(rules2, [4]types.T{floats[j], numbers[i], types.T_float64, types.T_float64})
			}
			for j := range strings {
				rules2 = append(rules2, [4]types.T{numbers[i], strings[j], types.T_float64, types.T_float64})
				rules2 = append(rules2, [4]types.T{strings[j], numbers[i], types.T_float64, types.T_float64})
			}
		}
		for i := range floats {
			for j := range decimals {
				rules2 = append(rules2, [4]types.T{floats[i], decimals[j], types.T_float64, types.T_float64})
				rules2 = append(rules2, [4]types.T{decimals[j], floats[i], types.T_float64, types.T_float64})
			}
			for j := range strings {
				rules2 = append(rules2, [4]types.T{floats[i], strings[j], types.T_float64, types.T_float64})
				rules2 = append(rules2, [4]types.T{strings[j], floats[i], types.T_float64, types.T_float64})
			}
		}
		rules2 = append(rules2, [4]types.T{types.T_decimal64, types.T_decimal128, types.T_decimal128, types.T_decimal128})
		rules2 = append(rules2, [4]types.T{types.T_decimal128, types.T_decimal64, types.T_decimal128, types.T_decimal128})
		for i := range decimals {
			for j := range strings {
				rules2 = append(rules2, [4]types.T{strings[j], decimals[i], decimals[i], decimals[i]})
				rules2 = append(rules2, [4]types.T{decimals[i], strings[j], decimals[i], decimals[i]})
			}
		}
	}

	binaryTable2 = make([][]binaryTargetTypes, maxTypes)
	for i := range binaryTable2 {
		binaryTable2[i] = make([]binaryTargetTypes, maxTypes)
	}
	for _, r := range rules2 {
		binaryTable2[r[0]][r[1]] = binaryTargetTypes{
			convert: true,
			left:    r[2],
			right:   r[3],
		}
	}

	// init castTable
	castTable = make([][]bool, maxTypes)
	{ // bool
		castTable[types.T_bool] = make([]bool, maxTypes)
		castTable[types.T_bool][types.T_bool] = true
		for _, typ := range strings {
			castTable[types.T_bool][typ] = true
		}
	}
	{ // date
		castTable[types.T_date] = make([]bool, maxTypes)
		castTable[types.T_date][types.T_date] = true
		castTable[types.T_date][types.T_timestamp] = true
		castTable[types.T_date][types.T_datetime] = true
		for _, typ := range strings {
			castTable[types.T_date][typ] = true
		}
	}
	{ // datetime
		castTable[types.T_datetime] = make([]bool, maxTypes)
		castTable[types.T_datetime][types.T_datetime] = true
		castTable[types.T_datetime][types.T_date] = true
		castTable[types.T_datetime][types.T_timestamp] = true
		for _, typ := range strings {
			castTable[types.T_datetime][typ] = true
		}
	}
	{ //  float
		for _, t := range floats {
			castTable[t] = make([]bool, maxTypes)
			for _, typ := range floats {
				castTable[t][typ] = true
			}
			for _, typ := range numbers {
				castTable[t][typ] = true
			}
			for _, typ := range strings {
				castTable[t][typ] = true
			}
		}
	}
	{ //  number
		for _, t := range numbers {
			castTable[t] = make([]bool, maxTypes)
			castTable[t][t] = true
			for _, typ := range floats {
				castTable[t][typ] = true
			}
			for _, typ := range numbers {
				castTable[t][typ] = true
			}
			castTable[t][types.T_timestamp] = true
			castTable[t][types.T_decimal64] = true
			castTable[t][types.T_decimal128] = true
			for _, typ := range strings {
				castTable[t][typ] = true
			}
		}
		castTable[types.T_decimal64] = make([]bool, maxTypes)
		castTable[types.T_decimal64][types.T_decimal64] = true
		castTable[types.T_decimal64][types.T_timestamp] = true
		for _, typ := range strings {
			castTable[types.T_decimal64][typ] = true
		}
		castTable[types.T_decimal128] = make([]bool, maxTypes)
		castTable[types.T_decimal128][types.T_decimal128] = true
		castTable[types.T_decimal128][types.T_timestamp] = true
		for _, typ := range strings {
			castTable[types.T_decimal128][typ] = true
		}
	}
	{ // timestamp
		castTable[types.T_timestamp] = make([]bool, maxTypes)
		castTable[types.T_timestamp][types.T_timestamp] = true
		castTable[types.T_timestamp][types.T_date] = true
		castTable[types.T_timestamp][types.T_datetime] = true
		for _, typ := range strings {
			castTable[types.T_timestamp][typ] = true
		}
	}
	{ // string
		for _, t := range strings {
			castTable[t] = make([]bool, maxTypes)
			for _, typ := range all {
				castTable[t][typ] = true
			}
		}
	}
}

var (
	// GeneralBinaryOperatorTypeCheckFn1 will check if params of the binary operators need type convert work
	GeneralBinaryOperatorTypeCheckFn1 = func(overloads []Function, inputs []types.T) (overloadIndex int32, ts []types.T) {
		return generalBinaryOperatorTypeCheckFn(overloads, inputs, generalBinaryParamsConvert)
	}

	// GeneralBinaryOperatorTypeCheckFn2 will check if params of the DIV and INTEGER_DIV need type convert work
	GeneralBinaryOperatorTypeCheckFn2 = func(overloads []Function, inputs []types.T) (overloadIndex int32, ts []types.T) {
		return generalBinaryOperatorTypeCheckFn(overloads, inputs, generalDivParamsConvert)
	}
)

func generalBinaryOperatorTypeCheckFn(overloads []Function, inputs []types.T, convertRule func(types.T, types.T) (types.T, types.T, bool)) (overloadIndex int32, ts []types.T) {
	if len(inputs) == 2 {
		matched := make([]int32, 0, 4)
		t1, t2, convert := convertRule(inputs[0], inputs[1])
		targets := []types.T{t1, t2}
		for _, o := range overloads {
			if tryToMatch(targets, o.Args) == matchedDirectly {
				matched = append(matched, o.Index)
			}
		}
		if len(matched) == 1 {
			if convert {
				return matched[0], targets
			}
			return matched[0], nil
		} else if len(matched) > 1 {
			for j := range inputs {
				if inputs[j] == ScalarNull {
					return matched[0], nil
				}
			}
			return tooManyFunctionsMatched, nil
		}
	}
	return wrongFunctionParameters, nil
}

func generalBinaryParamsConvert(l, r types.T) (types.T, types.T, bool) {
	ts := binaryTable[l][r] // targets
	if ts.convert {
		return ts.left, ts.right, true
	}
	return l, r, false
}

func generalDivParamsConvert(l, r types.T) (types.T, types.T, bool) {
	ts := binaryTable2[l][r]
	if ts.convert {
		return ts.left, ts.right, true
	}
	return l, r, false
}

func tryToMatch(inputs, requires []types.T) int {
	if len(inputs) == len(requires) {
		matchNumber, convNumber := 0, 0
		for i := 0; i < len(inputs); i++ {
			t1, t2 := inputs[i], requires[i]
			if t1 == t2 || t1 == ScalarNull {
				matchNumber++
			} else if castTable[t1][t2] {
				convNumber++
			} else {
				return matchFailed
			}
		}
		if matchNumber == len(inputs) {
			return matchedDirectly
		}
		return matchedByConvert
	}
	return matchedFailed
}
