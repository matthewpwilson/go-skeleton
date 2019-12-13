package skeleton_test

import (
	. "skeleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThingReturnsZero(t *testing.T) {
	assert.Equal(t, 0, Thing())
}
