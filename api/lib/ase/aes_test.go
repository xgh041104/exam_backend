package ase

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {

	//key := []byte{34, 219, 206, 233, 5, 202, 86, 231, 192, 155, 116, 62, 111, 207, 16, 94}
	keyStr := "1951EC8DA5"
	// key, _ := StringToByte(keyStr)
	key := GenerateKey(keyStr)
	iv := []byte("1234567890123456") // IV 也应与 C# 中相同
	ciphertext, _ := base64.StdEncoding.DecodeString("LBwBMTT71NFNXxdmSC3tENxoeBrYcFr2T4tMbfXD/Tk=")

	decrypted, err := decryptAESCFB(ciphertext, key, iv)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Decrypted:", string(decrypted))
	decryptedData, err := Decrypt(key, "F//ZDMvF0YfBxynSdRWKWjF7u6G5/sKVih09bbkod4g8DMMc7NT58efaBfw0aAGw")
	// encryptedData := "SWNSqytmzVxjU8V0j1AlUiLZx+0Km360IqTwesOlBATSjqcxaQgC9vos+s4SysqC5acYm8AzCuSXY2sDL/Lrfg=="
	data := "CookVR:UserAccount123456"
	fmt.Println("key", key)
	// 0B4TIDZLJzry1K2Lqs0EmCyzjFJM4sHDhYOckLJYAD0=
	//加密
	fmt.Println("data", []byte(data))
	encryptedData, _ := Encrypt(key, []byte(data))

	// encryptedData := "XpSOOlmD8vMmjniVuQcn6gMhcBFKOFNSX7dSzN0GQtVNxtTtaNJwCSiV1MD7fbC7"
	fmt.Println(encryptedData)
	//解密
	// decryptedData, err := Decrypt(key, encryptedData)
	fmt.Println("err:", err)
	fmt.Println("decryptedData:", string(decryptedData))
	fmt.Println("data:", data)
	fmt.Println(string(decryptedData) == data)
}

func decryptAESCFB(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
