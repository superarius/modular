package modular

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func PrintArray(arr []*Int) {
	for _, a := range arr {
		b := a.Value
		fmt.Printf("%d, ", b)
	}
	fmt.Println("")
}

func TestBasicMatrix(t *testing.T) {
	require := require.New(t)

	data := make([]*Int, 10)
	for i := range data {
		v, err := RandInt(nil)
		require.NoError(err)
		data[i] = v
	}

	// Test Create Matrix
	m := NewMatrix(2, 5, data)

	// Test Get Row/Col
	row := m.GetRow(2)
	require.Equal(row[4].Cmp(data[9]), 0, "get row failed")
	col := m.GetCol(5)
	require.Equal(row[4].Cmp(col[1]), 0, "get column failed")

	// Test Scalar Multiplication
	m.ScalarMul(new(Int).Exp(NewInt(2, nil), NewInt(256, nil)))
	require.Equal(0, m.values[9].Cmp(data[9].Mul(data[9], new(Int).Exp(NewInt(2, nil), NewInt(256, nil)))), "scalar mult failed")

	// Test Set Row/Col
	m.SetRow(1, []*Int{NewInt(1, nil), NewInt(1, nil), NewInt(1, nil), NewInt(1, nil), NewInt(1, nil)})
	require.Equal(0, m.values[0].Cmp(m.values[4]), "set row failed")
	m.SetCol(5, []*Int{NewInt(1, nil), NewInt(1, nil)})
	require.Equal(0, m.values[9].Cmp(m.values[4]), "set column failed")

}

func TestMultiplication(t *testing.T) {
	require := require.New(t)

	m1 := NewMatrix(2, 3, []*Int{NewInt(1, nil), NewInt(2, nil), NewInt(3, nil), NewInt(4, nil), NewInt(5, nil), NewInt(6, nil)})
	m2 := NewMatrix(3, 1, []*Int{NewInt(3, nil), NewInt(2, nil), NewInt(1, nil)})
	res, err := new(Matrix).Mul(m1, m2)
	require.NoError(err)
	require.Equal(len(res.values), 2, "wrong structure")
	require.Equal(0, res.values[0].Cmp(NewInt(10, nil)), "multiplication failed")
	require.Equal(0, res.values[1].Cmp(NewInt(28, nil)), "multiplication failed")
}

func TestInverse(t *testing.T) {
	require := require.New(t)

	// Test Gauss Jordan
	linearSystem := NewMatrix(2, 2, []*Int{NewInt(-1, nil), NewInt(2, nil), NewInt(1, nil), NewInt(0, nil)})
	ls := linearSystem.Represent2D()
	linearSystemResult := []*Int{NewInt(13, nil), NewInt(1, nil)}

	// Warning: the system is modified during the Gaussian reduction process
	// So it's better to pass a copy
	result, err := GaussJordan(linearSystem.Represent2D(), linearSystemResult)

	require.NoError(err, "gauss jordan failed")

	lhs := NewInt(0, nil)
	for i := 0; i < len(result); i++ {
		factor := new(Int).Mul(ls[0][i], result[i])
		lhs = new(Int).Add(lhs, factor)
	}

	require.Equal(0, linearSystemResult[0].Cmp(lhs), "System 1 failed")

	// require.Equal(0, result[0].Cmp(NewInt(1)), "gauss jordan failed")
	// require.Equal(0, result[1].Cmp(NewInt(1).Mod()), "gauss jordan failed")

	// Second matrix
	linearSystem2 := NewMatrix(3, 3, []*Int{NewInt(1, nil), NewInt(2, nil), NewInt(3, nil), NewInt(4, nil), NewInt(5, nil), NewInt(6, nil), NewInt(7, nil), NewInt(8, nil), NewInt(9, nil)})
	ls2 := linearSystem2.Represent2D()

	linearSystemResult2 := []*Int{NewInt(0, nil), NewInt(1, nil), NewInt(2, nil)}
	_ = linearSystemResult2

	result2, err := GaussJordan(ls2, linearSystemResult2)

	lhs = NewInt(0, nil)
	for i := 0; i < len(result2); i++ {
		factor := new(Int).Mul(ls2[0][i], result2[i])
		lhs = new(Int).Add(lhs, factor)
	}
	require.Equal(0, linearSystemResult2[0].Cmp(lhs), "System 2 failed")

	// Test Inverses
	m := NewMatrix(2, 2, []*Int{NewInt(7, nil), NewInt(-3, nil), NewInt(-2, nil), NewInt(1, nil)})
	inv, err := m.Inverse()
	require.NoError(err, "inverse failed")
	_ = inv
	require.Equal(0, inv.values[0].Cmp(NewInt(1, nil)), "inverse failed")

}
