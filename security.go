//security.go
//@auther huhaoran<huhaoran@domob.cn>
package security

import (
	"crypto/aes"
)

func encrypt(plaintext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	if len(plaintext) % aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(plaintext) > 0 {
		cipher.Encrypt(text, plaintext)
		plaintext = plaintext[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext
}

func decrypt(ciphertext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key[:aes.BlockSize]))
	if err != nil {
		panic(err.Error())
	}

	if len(ciphertext) % aes.BlockSize != 0 {
		panic("Need a multiple of the blocksize 16")
	}

	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	for len(ciphertext) > 0 {
		cipher.Decrypt(text, ciphertext)
		ciphertext = ciphertext[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}
	return plaintext
}

func PKCS7Pad(data []byte) []byte {
	padLength := aes.BlockSize - len(data)%aes.BlockSize
	for i:=0;i<padLength;i++ {
		data = append(data, byte(padLength))
	}
	return data
}

func PKCS7UPad(data []byte) []byte {
	padLength := int(data[len(data)-1])
	return data[:len(data)-padLength]
}

