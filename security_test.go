//test_security.go
package security

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key := "thisiskeyandlen>16"
	ciphertext := encrypt(PKCS7Pad([]byte("asdasdasdasdasd")), key)
	fmt.Printf("ciphertext: %x \n", ciphertext)
	plaintext := decrypt(ciphertext, key)
	fmt.Printf("plaintext: %s \n", PKCS7UPad([]byte(plaintext)))
}
