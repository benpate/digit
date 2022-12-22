package digit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapOfString(t *testing.T) {
	m := MapOfString{}

	require.True(t, m.SetString("a", "1"))
	require.True(t, m.SetString("b", "2"))
	require.True(t, m.SetString("c", "3"))

	require.Equal(t, "1", m.GetString("a"))
	require.Equal(t, "2", m.GetString("b"))
	require.Equal(t, "3", m.GetString("c"))

	m.Delete("b")
	require.Equal(t, "", m.GetString("b"))
}
