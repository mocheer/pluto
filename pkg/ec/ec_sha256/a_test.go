package ec_sha256_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ec/ec_sha256"
)

func TestXxx(t *testing.T) {
	t.Log([]byte("a"))
	t.Log(ec_sha256.EncryptToString([]byte("a")))
}
