package usage_test

import (
	meson_node_crypto "github.com/daqnext/meson-node-crypto"
	"log"
	"testing"
)

func Test_use(t *testing.T) {
	input := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	encrypt := meson_node_crypto.Encrypt(input)
	log.Println(encrypt)

	decrypt := meson_node_crypto.Decrypt(encrypt)
	log.Println(decrypt)
}
