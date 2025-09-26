package positions_test

import (
	"testing"

	"github.com/edipermadi/lameduck/pkg/core/positions"
	"github.com/stretchr/testify/require"
)

func TestPosition_File(t *testing.T) {
	require.Equal(t, 'a', positions.A1.File())
	require.Equal(t, 'b', positions.B2.File())
	require.Equal(t, 'c', positions.C3.File())
	require.Equal(t, 'd', positions.D4.File())
	require.Equal(t, 'e', positions.E5.File())
	require.Equal(t, 'f', positions.F6.File())
	require.Equal(t, 'g', positions.G7.File())
	require.Equal(t, 'h', positions.H8.File())

	require.Equal(t, '1', positions.A1.Rank())
	require.Equal(t, '2', positions.B2.Rank())
	require.Equal(t, '3', positions.C3.Rank())
	require.Equal(t, '4', positions.D4.Rank())
	require.Equal(t, '5', positions.E5.Rank())
	require.Equal(t, '6', positions.F6.Rank())
	require.Equal(t, '7', positions.G7.Rank())
	require.Equal(t, '8', positions.H8.Rank())
}

func TestNew(t *testing.T) {
	require.Equal(t, positions.A1, positions.New('a', '1'))
	require.Equal(t, positions.B2, positions.New('b', '2'))
	require.Equal(t, positions.C3, positions.New('c', '3'))
	require.Equal(t, positions.D4, positions.New('d', '4'))
	require.Equal(t, positions.E5, positions.New('e', '5'))
	require.Equal(t, positions.F6, positions.New('f', '6'))
	require.Equal(t, positions.G7, positions.New('g', '7'))
	require.Equal(t, positions.H8, positions.New('h', '8'))
}
