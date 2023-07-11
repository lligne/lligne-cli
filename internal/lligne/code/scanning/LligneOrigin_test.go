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
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneOrigin(t *testing.T) {

	t.Run("to string", func(t *testing.T) {
		sample := LligneOrigin{
			FileName: "sample.lligne",
			Line:     23,
			Column:   34,
		}

		assert.Equal(t, "sample.lligne(23,34)", sample.String(), "to string")
	})

}

//---------------------------------------------------------------------------------------------------------------------
