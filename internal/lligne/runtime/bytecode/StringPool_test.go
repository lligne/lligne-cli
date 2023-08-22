//
// # Tests of Disassembler.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestStringPool(t *testing.T) {

	t.Run("pooled strings", func(t *testing.T) {
		pool := NewStringConstantPool()

		i0 := pool.Put("Zero")
		i1 := pool.Put("One")
		i2 := pool.Put("Two")
		i3 := pool.Put("Three")
		i4 := pool.Put("Four")

		assert.Equal(t, uint64(0), i0)
		assert.Equal(t, uint64(1), i1)
		assert.Equal(t, uint64(2), i2)
		assert.Equal(t, uint64(3), i3)
		assert.Equal(t, uint64(4), i4)
		assert.Equal(t, "Zero", pool.Get(0))
		assert.Equal(t, "One", pool.Get(1))
		assert.Equal(t, "Two", pool.Get(2))
		assert.Equal(t, "Three", pool.Get(3))
		assert.Equal(t, "Four", pool.Get(4))

		i0 = pool.Put("Zero")
		i1 = pool.Put("One")
		i2 = pool.Put("Two")
		i3 = pool.Put("Three")
		i4 = pool.Put("Four")

		assert.Equal(t, uint64(0), i0)
		assert.Equal(t, uint64(1), i1)
		assert.Equal(t, uint64(2), i2)
		assert.Equal(t, uint64(3), i3)
		assert.Equal(t, uint64(4), i4)
		assert.Equal(t, "Zero", pool.Get(0))
		assert.Equal(t, "One", pool.Get(1))
		assert.Equal(t, "Two", pool.Get(2))
		assert.Equal(t, "Three", pool.Get(3))
		assert.Equal(t, "Four", pool.Get(4))
	})

}

//---------------------------------------------------------------------------------------------------------------------
