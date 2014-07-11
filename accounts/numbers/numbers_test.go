package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsoValidator(t *testing.T) {
	assert.Equal(t, true, validateIsoCountry("US"))
	assert.Equal(t, false, validateIsoCountry("USA"))
}

func TestLatLong(t *testing.T) {
	// TODO
}
