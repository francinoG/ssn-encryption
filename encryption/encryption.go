package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var (
	encryptKey = []byte("axiappsupersecretkeyaxiappsecret")
	nonce      = []byte("axiappSecret")
	aesgcm     cipher.AEAD
)

func init() {
	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	aesgcm, err = cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
}

func EncryptSSN(sensitiveVal string) string {
	newByte := []byte(sensitiveVal)
	encrypted := fmt.Sprintf("%x", aesgcm.Seal(nil, nonce, newByte, nil))
	return encrypted
}

func DecryptSSN(encrypted string) string {
	str, _ := hex.DecodeString(encrypted)
	decrypted, _ := aesgcm.Open(nil, nonce, str, nil)
	return string(decrypted)
}
