package meson_node_crypto

func Encrypt(input []byte) []byte {
	for i := range input {
		input[i] = ^input[i]
	}
	return input
}
