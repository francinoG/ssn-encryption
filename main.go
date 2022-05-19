package main

import (
	"fmt"
	"ssn/encryption"
)

func main() {
	ktp := "3507250310920001"
	ktpEncrypted := encryption.EncryptSSN(ktp)
	ktpDecrypted := encryption.DecryptSSN(ktpEncrypted)
	fmt.Println("ktp encrypted", ktpEncrypted)
	fmt.Println("ktp decrypted", ktpDecrypted)
}
