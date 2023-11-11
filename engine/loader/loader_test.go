package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultLoader_Load(t *testing.T) {
	loader := DefaultLoader{}

	state := loader.Load()

	assert.Len(t, state.Blueprint, 8)

	for _, row := range state.Blueprint {
		assert.Len(t, row, 8)
	}
}
