package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const (
	// 加密密钥
	key = "qwertyuiopasdfgh"
	// 初始向量
	iv = "qwertyuiopasdfgh"
)

// pKCS7Padding：假设数据长度需要填充 n(n>0) 个字节才对齐，那么填充 n 个字节，每个字节都是 n 。如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小

func pKCS7Padding(text []byte, blockSize int) []byte {
	// 计算待填充的长度
	padding := blockSize - len(text)%blockSize

	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(text, paddingText...)
}

func unPKCS7Padding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}

func Encrypt(text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	// 填充
	paddText := pKCS7Padding([]byte(text), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	// 加密
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)

	return base64.StdEncoding.EncodeToString(result)
}

func Decrypt(secret string) string {
	b, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println(err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	result := make([]byte, len(b))
	blockMode.CryptBlocks(result, b)
	// 去除填充
	result = unPKCS7Padding(result)

	return string(result)
}
