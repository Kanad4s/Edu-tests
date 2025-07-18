package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a int
	fmt.Fscan(in, &a)
	var str string
	for range a {
		fmt.Fscan(in, &str)
	}
	fmt.Fprint(out, a)
	fmt.Fprint(out, str)
}

func isExist(str string) bool {
	for i := 0; i < len(str); i++ {
		if
	}
}
