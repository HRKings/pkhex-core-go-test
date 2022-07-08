package utils

func GetOffset(bytes []byte, offset int, quantity int) []byte {
	return bytes[offset : offset+quantity]
}
