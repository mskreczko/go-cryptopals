package set1

import (
    "encoding/hex"
    "log"
)

func FixedXor(s1 string, s2 string) string {
    if len(s1) != len(s2) {
        log.Fatal("Input strings are not equal")
    }

    s1_bytes, s1_err := hex.DecodeString(s1)
    if s1_err != nil {
        log.Fatal(s1_err)
    }

    s2_bytes, s2_err := hex.DecodeString(s2)
    if s2_err != nil {
        log.Fatal(s2_err)
    }

    output := make([]byte, len(s1_bytes))
    for i := 0; i < len(s1_bytes); i++ {
        output[i] = s1_bytes[i] ^ s2_bytes[i]
    }

    return hex.EncodeToString(output)
}
