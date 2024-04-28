package set1

import (
    "encoding/hex"
)

func RepeatingKeyXor(input string) string {
    byte_input := []byte(input)

    key := []byte("ICE")
    result := make([]byte, len(byte_input))
    for i, b := range byte_input {
        result[i] = b ^ key[i % 3]
    }
    return hex.EncodeToString(result)
}
