//
// # Data types related to token origins.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneOrigin(t *testing.T) {

	t.Run("Token size should be 8 bytes", func(t *testing.T) {

		token := Token{
			SourceOffset: 25,
			SourceLength: 7,
			TokenType:    TokenTypeDoubleQuotedString,
		}
		expected := uintptr(8)
		actual := unsafe.Sizeof(token)

		assert.Equal(t, expected, actual, "wrong token size")

		assert.Equal(t, uintptr(16), unsafe.Sizeof("abc"), "wrong string size")
	})

}

//---------------------------------------------------------------------------------------------------------------------
