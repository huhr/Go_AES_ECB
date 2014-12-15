AES_ECB
=======
AES的ECB模式被认为是不安全的，所以golang没有提供官方函数，这里使用golang库中的块加解密算法，实现了
AES的ECB模式加密和解密算法，以及PKCS7Pad补全算法


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
