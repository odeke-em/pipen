package main

import (
	"fmt"
	"io"
	"os"

	"github.com/odeke-em/pipen"
)

func wcByShellUtil(f io.Reader) (io.Reader, error) {
	return pipen.StreamCommand(f, "wc")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "expecting paths...\n")
		return
	}

	rest := os.Args[1:]
	for _, arg := range rest {
		f, _ := os.Open(arg)
		if f == nil {
			continue
		}

		pr, err := wcByShellUtil(f)
		if err != nil {
			_ = f.Close()
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		io.Copy(os.Stdout, pr)
		_ = f.Close()
	}
}
