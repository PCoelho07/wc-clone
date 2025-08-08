package wc

import (
	"flag"
	"os"
)

type Config struct {
    C bool
    W bool
    L bool
	File string
}

const emptyTag = "-"

func ParseInput() Config {
    c := flag.Bool("c", false, "The number of bytes.")
    l := flag.Bool("l", false, "The number of lines.")
    w := flag.Bool("w", false, "The number of words.")

    flag.Parse()
    filename := flag.Arg(0)

    if len(filename) == 0 {
        filename = emptyTag
    }

    if !*c && !*w && !*l {
        return Config{
            C: true,
            W: true,
            L: true,
            File: filename,
        }
    }

	return Config{
        C: *c,
        W: *w,
        L: *l,
		File: filename,
	}
}

func OpenFileOrStdin(filename string) (*os.File, error) {
    if filename == emptyTag {
        return os.Stdin, nil
    }

    return os.Open(filename)
}
