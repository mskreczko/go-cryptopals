package set1

import (
    "encoding/base64"
    "encoding/hex"
    "log"
)

func Solve(input string) string {
    input_bytes, err := hex.DecodeString(input)
    if err != nil {
        log.Fatal(err)
    }

    return base64.StdEncoding.EncodeToString(input_bytes)
}
