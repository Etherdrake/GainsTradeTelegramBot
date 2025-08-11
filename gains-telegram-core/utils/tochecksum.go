package utils

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"strings"
	"unicode"
)

// ToChecksumAddress converts an Ethereum address to its checksummed version according to EIP-55
func ToChecksumAddress(address string) (string, error) {
	// Remove the '0x' prefix if it exists
	if strings.HasPrefix(address, "0x") {
		address = address[2:]
	}

	// Convert the address to lowercase
	address = strings.ToLower(address)

	// Compute the Keccak-256 hash of the lowercase address
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(address))
	hashBytes := hash.Sum(nil)
	hashHex := hex.EncodeToString(hashBytes)

	// Apply the checksum
	var checksumAddress strings.Builder
	checksumAddress.WriteString("0x")
	for i := 0; i < len(address); i++ {
		if hashHex[i] >= '8' {
			checksumAddress.WriteByte(byte(unicode.ToUpper(rune(address[i]))))
		} else {
			checksumAddress.WriteByte(address[i])
		}
	}

	return checksumAddress.String(), nil
}
