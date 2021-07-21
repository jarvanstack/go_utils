package stringu

import "math/rand"

//LETTER_BYTES some random data
const LETTER_BYTES string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GetLengthString GetStringPointer random stringu by length.
func GetLengthString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Int63()%int64(len(LETTER_BYTES))]
	}
	return string(b)
}
func GetLengthBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Int63()%int64(len(LETTER_BYTES))]
	}
	return b
}
