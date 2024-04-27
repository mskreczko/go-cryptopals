package set1

import (
    "bufio"
    "log"
    "os"
)

func applySingleByteXor(input string) string {
    return SingleByteXor(input)
}

func DetectSingleByteXor(filename string) string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    lines := make([]string, 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    xored := make([]string, 0)
    for _, ln := range lines {
        xored = append(xored, applySingleByteXor(ln))
    }

    best_score := 0.0
    var best_key string
    for _, ln := range xored {
        if Score([]byte(ln)) > best_score {
            best_score = Score([]byte(ln))
            best_key = ln
        }
    }
    return best_key
}
