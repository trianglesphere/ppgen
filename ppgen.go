package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"text/tabwriter"

	"internal/wordlists"
)

func printEntropy() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "List\tBits of entropy per word")
	for name, list := range wordlists.Lists {
		fmt.Fprintf(w, "%s\t%0.1f\n", name, math.Log2(float64(len(list))))
	}
	w.Flush()
}

func printWordList(name string) {
	list, ok := wordlists.Lists[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Do not have list %s.\n", name)
		os.Exit(-1)
	}
	for _, word := range list {
		fmt.Println(word)
	}
}

func main() {
	var (
		entropy                                  bool
		num                                      int
		upper, underscore, digit, punct, special bool
		list, printList                          string
	)
	// Usage
	// ppgen --entropy
	// ppgen --print LIST
	// ppgen [-n 10] [--list LIST] [-s] [-p] [-d] [-u]
	flag.BoolVar(&entropy, "entropy", false, "print per word entropy for each list")
	flag.IntVar(&num, "n", 6, "number of words in passphrase")
	flag.IntVar(&num, "number", 6, "number of words in passphrase")
	flag.BoolVar(&upper, "upper", false, "include upper case letter")
	flag.BoolVar(&digit, "d", false, "include digit (0-9)")
	flag.BoolVar(&digit, "digit", false, "include digit (0-9)")
	flag.BoolVar(&punct, "p", false, "include punctuation characters")
	flag.BoolVar(&punct, "punctuation", false, "include punctuation characters")
	flag.BoolVar(&special, "s", false, "include upper case, digit, and punctuation character to comply with password requirements")
	flag.BoolVar(&special, "special", false, "include upper case, digit, and punctuation character to comply with password requirements")
	flag.BoolVar(&underscore, "u", false, "replace spaces in word with underscores")
	flag.BoolVar(&underscore, "underscore", false, "replace spaces in word with underscores")
	flag.StringVar(&list, "list", "Number", "word list to use (case insensitive, use --version to see installed word lists & aliases)")
	flag.StringVar(&printList, "print", "", "word list to print")
	flag.Parse()

	switch {
	case entropy:
		printEntropy()
		os.Exit(0)
	case printList != "":
		printWordList(printList)
		os.Exit(0)
	// case version
	default:
		// verify list
	}
	wlist := wordlists.Lists[list]
	for i := 0; i < num; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
		fmt.Printf("%s ", wlist[n.Int64()])
	}
}
