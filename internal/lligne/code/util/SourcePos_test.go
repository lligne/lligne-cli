//
// # Data types related to token origins.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneOrigin(t *testing.T) {

	t.Run("SourcePos size should be 8 bytes", func(t *testing.T) {

		token := SourcePos{
			startOffset: 45,
			endOffset:   50,
		}
		expected := uintptr(8)
		actual := unsafe.Sizeof(token)

		assert.Equal(t, expected, actual, "to string")
	})

}

//---------------------------------------------------------------------------------------------------------------------
