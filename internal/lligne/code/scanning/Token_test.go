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

	t.Run("token size", func(t *testing.T) {

		token := Token{
			SourceOffset: 25,
			SourceLength: 7,
			TokenType:    TokenTypeDoubleQuotedString,
		}
		expected := uintptr(8)
		actual := unsafe.Sizeof(token)

		assert.Equal(t, expected, actual, "to string")
	})

}

//---------------------------------------------------------------------------------------------------------------------
