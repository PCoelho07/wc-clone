package wc

import "fmt"

func Print(lines, words, bytes int, cfg Config) {
    if cfg.C {
        fmt.Printf("%8d", bytes)
    }

    if cfg.W {
        fmt.Printf("%8d", words)
    }

    if cfg.L {
        fmt.Printf("%8d", lines)
    }

    fmt.Printf(" %s\n", cfg.File)
}
