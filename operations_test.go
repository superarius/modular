package modular

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestModInverse(t *testing.T) {
	require := require.New(t)

	for i := 0; i < 100; i++ {
		point, err := RandInt(nil)
		require.NoError(err)
		require.Equal(point.Value.Cmp(defaultP), -1, "out of bounds")
		inverse := ModInverse(point)
		require.Equal(inverse.Value.Cmp(defaultP), -1, "Inverse out of bounds")
		point.Mul(point, inverse)
		require.Equal(0, point.Cmp(NewInt(1, nil)), "Inverse incorrect")
	}
}
func TestRandomInt(t *testing.T) {
	require := require.New(t)
	for i := 0; i < 100; i++ {
		x, err := RandInt(nil)
		require.NoError(err)
		require.Equal(-1, x.Value.Cmp(defaultP), "Inverse out of bounds")
	}
}

func TestOperations(t *testing.T) {
	require := require.New(t)

	// basic multiplication
	check := big.NewInt(1234)
	check.Mul(check, big.NewInt(2))
	res := new(Int).Mul(NewInt(1234, nil), NewInt(2, nil))
	require.Equal(0, res.Value.Cmp(check.Mod(check, res.Base)), "multiplication failure")

	// modular overflow multiplication
	check = new(big.Int).Exp(big.NewInt(2), big.NewInt(500), nil)
	res = new(Int).Exp(NewInt(2, nil), NewInt(500, nil))
	check.Mul(check, check)
	res.Mul(res, res)
	require.Equal(-1, res.Value.Cmp(check), "did not automatically reduce")
	require.Equal(0, res.Value.Cmp(check.Mod(check, res.Base)), "did not reduce properly")

	// basic addition
	check = big.NewInt(1234)
	check.Mul(check, big.NewInt(2))
	res = new(Int).Add(NewInt(1234, nil), NewInt(1234, nil))
	require.Equal(0, res.Value.Cmp(check.Mod(check, res.Base)), "addition failure")

	// modular overflow addition
	check = new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	check.Mul(check, big.NewInt(2))
	res = new(Int).Exp(NewInt(2, nil), NewInt(1000, nil))
	res.Add(res, res)
	require.Equal(-1, res.Value.Cmp(check), "did not automatically reduce")
	require.Equal(0, res.Value.Cmp(check.Mod(check, res.Base)), "did not reduce properly")

	// modular subtraction
	check = new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	check.Sub(check, big.NewInt(1000000))
	res = new(Int).Exp(NewInt(2, nil), NewInt(1000, nil))
	res.Sub(res, NewInt(1000000, nil))
	require.Equal(-1, res.Value.Cmp(check), "did not automatically reduce")
	require.Equal(0, res.Value.Cmp(check.Mod(check, res.Base)), "did not reduce properly")

	// test linear combination
	vec1 := []*Int{NewInt(1, nil), NewInt(2, nil), NewInt(3, nil), NewInt(4, nil)}
	vec2 := []*Int{NewInt(18000, nil), NewInt(9000, nil), NewInt(6000, nil), NewInt(4500, nil)}
	check = new(big.Int).Mul(big.NewInt(18000), big.NewInt(4))
	res = new(Int).LinearCombination(vec1, vec2)
	require.Equal(0, res.Value.Cmp(check.Mod(check, res.Base)), "linear combination failure")

	// test String()
	intstra := "191919199191919191919191919191919"
	intstrb, err := IntFromString(intstra, 10, nil)
	require.NoError(err)
	require.Equal(intstra, intstrb.String(), "output string is not equal to initial string")

}
