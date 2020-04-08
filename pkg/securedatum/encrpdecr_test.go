package secureddatum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingEncryptDecrypt(t *testing.T) {
	datum, err := New(EnvCfgProvider.All()...)
	assert.Nil(t, err)
	assert.NotNil(t, datum)
}
