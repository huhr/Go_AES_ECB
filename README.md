AES_ECB
=======
golang的AES ECB模式加解密，以及PKCS7Pad

```
import (
	"huhr/security"
)

func main() {
	key := "thisiskeyandlen>16"
	//加密
	ciphertext := encrypt(PKCS7Pad([]byte("asdasdasdasdasd")), key)
	//解密
	plaintext := string(PKCS7UPad([]byte(decrypt(ciphertext, key))))
}
```
