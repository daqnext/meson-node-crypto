package meson_node_crypto

func Decrypt(input []byte) []byte {
	for i := range input {
		input[i] = ^input[i]
	}
	return input
}
