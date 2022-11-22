package util

import "testing"

func TestEncrypt(t *testing.T) {
	data := "123456@3213dsfsdf"
	res := Encrypt(data)
	t.Log(res)

	r := Decrypt(res)
	t.Log(r)
}
