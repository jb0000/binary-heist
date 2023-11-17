package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultLoader_Load(t *testing.T) {
	config := Config{
		layout: `
X 
 X`,
	}

	loader := NewConfigLoader(config)

	state := loader.Load()

	assert.Len(t, state.Blueprint, 3)

	for _, row := range state.Blueprint {
		assert.Len(t, row, 2)
	}

	assert.False(t, state.Blueprint[0][0].IsPassable)
	assert.False(t, state.Blueprint[0][1].IsPassable)
	assert.False(t, state.Blueprint[1][0].IsPassable)
	assert.True(t, state.Blueprint[1][1].IsPassable)
	assert.True(t, state.Blueprint[2][0].IsPassable)
	assert.False(t, state.Blueprint[2][1].IsPassable)
}
