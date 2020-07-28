package wordlists

import (
	"fmt"
	"strings"
)

type List int

const (
	EFFLarge List = iota
	EFFShort
	EFFPrefix
	Punct
	Upper
	Digit
)

func (l List) String() string {
	return [...]string{"EFF Large", "EFF Short", "EFF Prefix", "Punctuation", "Upper Case", "Number"}[l]
}

// Lists contains the available wordlists and their display name.
var Lists = map[List][]string{
	EFFLarge:  Eff_large_list,
	EFFShort:  Eff_short_list,
	EFFPrefix: Eff_short_prefix_list,
	Punct:     Punct_list,
	Upper:     Upper_list,
	Digit:     Digit_list,
}

// Typeable name to word list
var Aliases = map[string]List{
	"eff large":   EFFLarge,
	"large":       EFFLarge,
	"eff short":   EFFShort,
	"short":       EFFShort,
	"eff prefix":  EFFPrefix,
	"prefix":      EFFPrefix,
	"punctuation": Punct,
	"punct":       Punct,
	"upper case":  Upper,
	"upper":       Upper,
	"number":      Digit,
	"numbers":     Digit,
	"digits":      Digit,
}

func Wordlist(name string) ([]string, error) {
	list, ok := Aliases[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("'%s' is not a valid word list name", name)
	}
	return Lists[list], nil
}
