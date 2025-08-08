package main

import (
	"fmt"
	"time"

	"github.com/PCoelho07/ccwc/internal/wc"
)

func main() {
	start := time.Now()
	cfg := wc.ParseInput()
    lines, words, bytes := wc.Process(cfg)

    if len(cfg.File) > 0 {
        wc.Print(lines, words, bytes, cfg)
    }

	elapsed := time.Since(start)
	fmt.Printf("\nexecution time: %s\n", elapsed)
}
