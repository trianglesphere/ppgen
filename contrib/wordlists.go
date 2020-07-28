package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		var words []string
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		name := strings.TrimSuffix(path.Base(filename), path.Ext(filename))
		out, err := os.Create("../internal/wordlists/" + name + ".go")
		defer out.Close()
		o := bufio.NewWriter(out)
		defer o.Flush()
		fmt.Fprintln(o, "package wordlists\n")
		fmt.Fprintf(o, "var %s_list = []string{", strings.Title(name))
		for i := 0; len(words) >= 1 && i < len(words)-1; i++ {
			fmt.Fprintf(o, "\"%s\",", words[i])
		}
		fmt.Fprintf(o, "\"%s\"}\n", words[len(words)-1])
	}
}
