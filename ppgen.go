package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"internal/wordlists"
)

func printVersion() {
	fmt.Println("Version 0 of ppgen.")
	fmt.Println("Author: Joshua Gutow")
	fmt.Println("Use --print LIST_NAME to print the specific list.")
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "List\tLength\tEntropy/word")
	ids := make([]int, 0, len(wordlists.Lists))
	for id := range wordlists.Lists {
		ids = append(ids, int(id))

	}
	sort.Ints(ids)
	for id := range ids {
		listID := wordlists.ListID(id)
		list := wordlists.Lists[listID]
		bits := math.Log2(float64(len(list)))
		fmt.Fprintf(w, "%s\t%d\t%0.1f\n", listID, len(list), bits)
	}
	w.Flush()
}

func printWordList(name string) {
	list, err := wordlists.List(name)
	if err != nil {
		log.Fatalf("%v. Use --version to see valid wordlists.", err)
	}
	for _, word := range list {
		fmt.Println(word)
	}
}

const usage = `Usage:
	ppgen --version
	ppgen --print LIST_NAME
	ppgen [-n NUM_WORDS]  [-l LIST_NAME] [-u] [-s] [-p] [-d] [--upper]

Options:
	--version
	--print LIST_NAME
	--list LIST_NAME
	-n, --number
	-u, --underscore
	-s, --special
	-p, --punctuation
	-d, --digit
	--upper`

func main() {
	var (
		versionFlag                                  bool
		numFlag                                      int
		upperFlag, digitFlag, punctFlag, specialFlag bool
		underscoreFlag                               bool
		listFlag, printListFlag                      string
	)
	log.SetFlags(0)
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }
	flag.BoolVar(&versionFlag, "version", false, "print version and information about installed word lists")
	flag.IntVar(&numFlag, "n", 6, "number of words in passphrase")
	flag.IntVar(&numFlag, "number", 6, "number of words in passphrase")
	flag.BoolVar(&upperFlag, "upper", false, "include upper case letter")
	flag.BoolVar(&digitFlag, "d", false, "include digit (0-9)")
	flag.BoolVar(&digitFlag, "digit", false, "include digit (0-9)")
	flag.BoolVar(&punctFlag, "p", false, "include punctuation characters")
	flag.BoolVar(&punctFlag, "punctuation", false, "include punctuation characters")
	flag.BoolVar(&specialFlag, "s", false, "include upper case, digit, and punctuation character")
	flag.BoolVar(&specialFlag, "special", false, "include upper case, digit, and punctuation character")
	flag.BoolVar(&underscoreFlag, "u", false, "replace spaces in word with underscores")
	flag.BoolVar(&underscoreFlag, "underscore", false, "replace spaces in word with underscores")
	flag.StringVar(&listFlag, "list", "large", "word list to use (case insensitive, use --version to see installed word lists)")
	flag.StringVar(&printListFlag, "print", "", "word list to print")
	flag.Parse()

	switch {
	case printListFlag != "":
		printWordList(printListFlag)
	case versionFlag:
		printVersion()
	default:
		list, err := wordlists.List(listFlag)
		if err != nil {
			log.Fatalf("%v. Use --version to see valid wordlists.", err)
		}
		phrase := []string{}
		max := big.NewInt(int64(len(list)))
		for i := 0; i < numFlag; i++ {
			n, _ := rand.Int(rand.Reader, max)
			phrase = append(phrase, list[n.Int64()])
		}
		extra, ok := extra(specialFlag, upperFlag, digitFlag, punctFlag)
		if ok {
			phrase = append(phrase, extra)
		}
		var separator string
		if underscoreFlag {
			separator = "_"
		} else {
			separator = " "
		}
		fmt.Println(strings.Join(phrase, separator))
	}
}

func extra(special, upper, digit, punct bool) (string, bool) {
	var extra strings.Builder
	if upper || special {
		list, _ := wordlists.List("upper case")
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
		extra.WriteString(list[n.Int64()])
	}
	if digit || special {
		list, _ := wordlists.List("number")
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
		extra.WriteString(list[n.Int64()])
	}
	if punct || special {
		list, _ := wordlists.List("punctuation")
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
		extra.WriteString(list[n.Int64()])
	}
	return extra.String(), special || upper || digit || punct
}
