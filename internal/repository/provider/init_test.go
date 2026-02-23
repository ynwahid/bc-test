package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAPIRepository(t *testing.T) {
	got := NewAPIRepository()
	require.NotNil(t, got)
}
