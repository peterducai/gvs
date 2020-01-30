package gvs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func decrypt(cipherstring string, keystring string) string {
	// Byte array of the string
	ciphertext := []byte(cipherstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func encrypt(plainstring, keystring string) string {
	// Byte array of the string
	plaintext := []byte(plainstring)

	// Key
	key := []byte(keystring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func waitForKey() {
	fmt.Print("\nPlease Enter to continue...")
	var buf [1]byte
	os.Stdin.Read(buf[:])
}

func isFile(file string) bool {
	if s, err := os.Stat(file); s.IsDir() || err != nil {
		return false
	}

	return true
}

// func main() {

// 	if len(os.Args) != 2 {
// 		fmt.Println("You must drag and drop a file!")
// 		waitForKey()
// 		os.Exit(1)
// 	} else if !isFile(os.Args[1]) {
// 		fmt.Println("File does not exist!")
// 		waitForKey()
// 		os.Exit(1)
// 	}

// 	file := os.Args[1]
// 	key := "testtesttesttest"

// 	if strings.ToLower(filepath.Ext(file)) != ".enc" {
// 		content, err := readFromFile(file)
// 		if err != nil {
// 			fmt.Println(err)
// 			waitForKey()
// 			os.Exit(1)
// 		}
// 		encrypted := encrypt(string(content), key)
// 		writeToFile(encrypted, file+".enc")
// 	} else {
// 		content, err := readFromFile(file)
// 		if err != nil {
// 			fmt.Println(err)
// 			waitForKey()
// 			os.Exit(1)
// 		}
// 		decrypted := decrypt(string(content), key)
// 		writeToFile(decrypted, file[:len(file)-4])
// 	}
// }
