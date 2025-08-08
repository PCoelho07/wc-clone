package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Process(cfg Config) (lines, words, bytes int) {
    file, err := OpenFileOrStdin(cfg.File)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return 0, 0, 0
    } 

    defer file.Close()

    scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		bytes += len(line) + 2
        words += len(strings.Fields(line))
	}

    return lines, words, bytes
}
