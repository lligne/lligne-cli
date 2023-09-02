//
// # Tests of Disassembler.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestStringPool(t *testing.T) {

	t.Run("pooled built-in types", func(t *testing.T) {
		pool := NewTypePool()

		assert.Equal(t, TypeTypeInstance, pool.Get(0))
		assert.Equal(t, BoolTypeInstance, pool.Get(1))
		assert.Equal(t, Float64TypeInstance, pool.Get(2))
		assert.Equal(t, Int64TypeInstance, pool.Get(3))
		assert.Equal(t, StringTypeInstance, pool.Get(4))
	})

}

//---------------------------------------------------------------------------------------------------------------------
