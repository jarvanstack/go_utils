package stringu

import "encoding/base64"

func Base64ToBytes(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func ByteToBase64(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}
