package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/swayedev/fcrypt"
)

func main() {
	// Set the chunk size for file encryption and decryption
	chunkSize := 2 * 1024 * 1024 // 2MB
	// Example usages
	passphrase := "this is a passphrase"

	filePath := "./files/example.txt"
	largeFilePath := "./files/freeman-zhou-oV9hp8wXkPE-unsplash.jpg"

	stringToHash := "Example data"
	stringToEncrypt := "Sensitive information"

	// Hashing
	hash := HashString(stringToHash)
	fmt.Printf("Hash: %s\n", hash)

	// Comparing Hashes
	isEqual := CompareHash(stringToHash, hash)
	fmt.Printf("Hash comparison: %v\n", isEqual)

	// compare keys
	// if the master key and salt is the same, the key will be the same
	// if the master key is the same but the salt is different, the key will be different
	key1, err := fcrypt.GenerateKey(passphrase, nil, fcrypt.DefaultKeyLength)
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	fmt.Printf("Key1: %x\n", key1)

	key2, err := fcrypt.GenerateKey(passphrase, nil, fcrypt.DefaultKeyLength)
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	fmt.Printf("Key2: %x\n", key2)
	areEqual := CompareKeys(key1, key2)
	fmt.Printf("Keys 1 and 2 are equal: %v\n", areEqual)

	// for a totally random key and salt pair, use fcrypt.GenerateSaltAndKey
	salt, key3, err := fcrypt.GenerateSaltAndKey(passphrase, 16, fcrypt.DefaultKeyLength)
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	// store the salt if you wish to create the key again
	fmt.Printf("Salt: %x\n", salt)
	// store the key for encryption and decryption
	fmt.Printf("Key3: %x\n", key3)

	// Encrypting strings
	encryptedString, err := EncryptString(stringToEncrypt, key1)
	if err != nil {
		log.Fatalf("Failed to encrypt string: %v", err)
	}
	fmt.Println("Encrypted String:", string(encryptedString))

	// Decrypting strings
	decryptedString := DecryptString(encryptedString, key1)
	fmt.Println("Decrypted String:", decryptedString)

	// Encrypting files
	EncryptFile(filePath, filePath+".enc", key1, chunkSize)

	// Decrypting files
	DecryptFile(filePath+".enc", filePath+".dec", key1, chunkSize)

	// Rotating keys
	RotateKey(filePath+".enc", filePath+".reenc", key1, key3, chunkSize)

	// Encrypting large files
	EncryptLargeFile(largeFilePath, largeFilePath+".enc", key3, chunkSize)

	// Decrypting large files
	DecryptLargeFile(largeFilePath+".enc", largeFilePath+".dec", key3, chunkSize)
}

// HashString hashes a string using SHA3-256
func HashString(data string) string {
	return fmt.Sprintf("%x", fcrypt.HashBytesSHA3([]byte(data)))
}

// CompareHash compares a string to its hashed value
func CompareHash(data, hash string) bool {
	return HashString(data) == hash
}

// CompareKeys compares two keys
func CompareKeys(key1, key2 []byte) bool {
	return bytes.Equal(key1, key2)
}

// EncryptString encrypts a string using the given passphrase
func EncryptString(data string, key []byte) ([]byte, error) {
	encrypted, err := fcrypt.Encrypt([]byte(data), key)
	if err != nil {
		return nil, err
	}
	return encrypted, nil // Salt is prepended to encrypted data
}

// DecryptString decrypts an encrypted string using the given passphrase
func DecryptString(encryptedData, key []byte) string {
	decrypted, err := fcrypt.Decrypt(encryptedData, key)
	if err != nil {
		log.Fatalf("Failed to decrypt string: %v", err)
	}
	return string(decrypted)
}

// EncryptFile encrypts a file and writes to a destination file
func EncryptFile(src, dest string, key []byte, chunkSize int) {
	// Open the source file
	file, err := os.Open(src)
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}
	defer file.Close()

	err = fcrypt.EncryptFileToFile(file, key, chunkSize, dest)
	if err != nil {
		log.Fatalf("Failed to encrypt file: %v", err)
	}

	fmt.Printf("File encrypted successfully: %s\n", dest)
}

// DecryptFile decrypts an encrypted file and writes to a destination file
func DecryptFile(src, dest string, key []byte, chunkSize int) {
	// Open the encrypted file to extract the salt
	file, err := os.Open(src)
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}
	defer file.Close()

	// Perform the decryption
	err = fcrypt.DecryptFileToFile(src, dest, key, chunkSize)
	if err != nil {
		log.Fatalf("Failed to decrypt file: %v", err)
	}

	fmt.Printf("File decrypted successfully: %s\n", dest)
}

// RotateKey re-encrypts a file with a new passphrase
func RotateKey(src, dest string, oldKey, newKey []byte, chunkSize int) error {
	err := fcrypt.ReEncryptFileToFile(src, dest, oldKey, newKey, chunkSize)
	if err != nil {
		return err
	}
	fmt.Printf("File re-encrypted with a new key successfully: %s\n", dest)
	return nil
}

func EncryptLargeFile(src, dest string, key []byte, chunkSize int) error {
	// Open the source file
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer file.Close()

	// Create the destination file
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	// Create an encryption writer
	encryptedFile, err := fcrypt.StreamEncrypt(file, key)
	if err != nil {
		return fmt.Errorf("failed to create encryption stream: %v", err)
	}

	// Write data in chunks
	buffer := make([]byte, chunkSize)
	for {
		n, readErr := encryptedFile.Read(buffer)
		if n > 0 {
			_, writeErr := destFile.Write(buffer[:n])
			if writeErr != nil {
				return fmt.Errorf("failed to write to destination file: %v", writeErr)
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return fmt.Errorf("failed to read from source file: %v", readErr)
		}
	}

	fmt.Printf("File encrypted successfully: %s\n", dest)
	return nil
}

func DecryptLargeFile(src, dest string, key []byte, chunkSize int) error {
	// Open the encrypted source file
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open encrypted file: %v", err)
	}
	defer file.Close()

	// Create the destination file for the decrypted output
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer destFile.Close()

	// Create a decryption reader
	decryptedFile, err := fcrypt.StreamDecrypt(file, key)
	if err != nil {
		return fmt.Errorf("failed to create decryption stream: %v", err)
	}

	// Write data in chunks
	buffer := make([]byte, chunkSize)
	for {
		n, readErr := decryptedFile.Read(buffer)
		if n > 0 {
			_, writeErr := destFile.Write(buffer[:n])
			if writeErr != nil {
				return fmt.Errorf("failed to write to destination file: %v", writeErr)
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return fmt.Errorf("failed to read from encrypted file: %v", readErr)
		}
	}

	fmt.Printf("File decrypted successfully: %s\n", dest)
	return nil
}
