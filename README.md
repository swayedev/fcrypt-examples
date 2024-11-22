# Fcrypt Examples

This repository demonstrates the use of the [Fcrypt](https://github.com/swayedev/fcrypt) `v0.2.2` package for cryptographic operations in Go. It includes examples for hashing, encrypting/decrypting strings and files, and rotating encryption keys. Additionally, it handles large files efficiently by processing data in chunks.

---

## Features

- **Hashing**: 
  - Strings
  - Files
- **Encryption and Decryption**:
  - Strings
  - Files
  - Large files (chunked processing)
- **Key Management**:
  - Key generation
  - Salt usage
  - Key comparison
  - Key rotation

---

## Key Management

The master key can create new keys that vary based on the salt used. Fcrypt supports a `nil` salt, which will generate a consistent key based solely on the master passphrase. This allows flexibility in key generation while maintaining security.

---

## How to Use

To explore the various features of the repository, run the `main.go` file as follows:

```bash
go run main.go
```

The `main.go` file demonstrates different aspects of Fcrypt's capabilities. Each function showcases a specific use case, such as hashing, encrypting strings, or processing files.

### Example Functions in `main.go`

- **`HashString`**: Demonstrates how to hash a string using SHA3-256.
- **`CompareHash`**: Shows how to compare a string against a hash.
- **`HashFile`**: Demonstrates file hashing to verify file integrity.
- **`EncryptString` / `DecryptString`**: Encrypt and decrypt strings using a generated key.
- **`EncryptFile` / `DecryptFile`**: Encrypt and decrypt files with a specified chunk size.
- **`EncryptLargeFile` / `DecryptLargeFile`**: Efficiently process large files by handling them in chunks.
- **`RotateKey`**: Re-encrypt a file with a new key.

### Inspecting and Using Functions

To fully understand and integrate Fcrypt into your projects, inspect the functions in `main.go`. Each function is designed to demonstrate Fcrypt's features and provide a base for custom implementations in your own applications.

---

## Prerequisites

1. Install Go (1.23 or later recommended).
2. Clone this repository:
   ```bash
   git clone https://github.com/swayedev/fcrypt-examples.git
   cd fcrypt-examples
   ```
3. Install the `fcrypt` package:
   ```bash
   go get github.com/swayedev/fcrypt
   ```

---

## How to Run

```bash
go run main.go
```

Inspect the output for examples of hashing, encryption, decryption, and key rotation. To adapt the functions for your use case, review their implementations in the `main.go` file.

---

## Credits

- **Image Used**:
  - Filename: `freeman-zhou-oV9hp8wXkPE-unsplash.jpg`
  - Photographer: [Freeman Zhou](https://unsplash.com/@freeman_zhou)
  - Source: [Unsplash](https://unsplash.com/photos/oV9hp8wXkPE)
  - License: [Unsplash License](https://unsplash.com/license)

---

## License

This project is licensed under the [MIT License](LICENSE).
