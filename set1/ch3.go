package set1

import (
    "encoding/hex"
    "log"
)

func Score(input []byte) float64 {
    freq_table := map[string]float64 {
        "e": 11.1607,
        "a": 8.4966,
        "r": 7.5809,
        "i": 7.5448,
        "o": 7.1635,
        "t": 6.9509,
        "n": 6.6544,
        "s": 5.7351,
        "l": 5.4893,
        "c": 4.5388,
        "u": 3.3608,
        "d": 3.3844,
        "p": 3.1671,
        "m": 3.0129,
        "h": 3.0034,
        "g": 2.4705,
        "b": 2.0720,
        "f": 1.8121,
        "y": 1.7779,
        "w": 1.2899,
        "k": 1.1016,
        "v": 1.0074,
        "x": 0.2902,
        "z": 0.2722,
        "j": 0.1965,
        "q": 0.1962,
    }

    score := 0.0
    for _, el := range input {
        score += freq_table[string(el)]
    }
    return score
}

func xor(input []byte, key byte) []byte {
    output := make([]byte, len(input))
    for i := 0; i < len(input); i++ {
        output[i] = input[i] ^ key
    }
    return output
}

func SingleByteXor(input string) string {
    byte_input, err := hex.DecodeString(input)
    if err != nil {
        log.Fatal(err)
    }

    best_score := 0.0
    var best_key byte
    for i := 0; i < 256; i++ {
        result := Score(xor(byte_input, byte(i)))
        if result > best_score {
            best_score = result
            best_key = byte(i)
        }
    }

    return string(xor(byte_input, best_key))
}
