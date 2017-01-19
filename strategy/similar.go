package strategy

import (
	"fmt"
	"strings"

	"zenithar.org/go/typogenerator/helpers"
	"zenithar.org/go/typogenerator/mapping"
)

type similarStrategy struct {
	_mapping mapping.Mapping
}

// Similar returns a similar generation strategy
func Similar(m mapping.Mapping) Strategy {
	return &similarStrategy{
		_mapping: m,
	}
}

// -----------------------------------------------------------------------------

func (s *similarStrategy) Generate(domain string) ([]string, error) {
	res := []string{}

	// Split domain and gTLD
	parts := strings.SplitN(domain, ".", 2)

	dom := []rune(parts[0])
	tld := parts[1]

	for ws := range dom {
		for i := 0; i < ((len(dom) - ws) + 1); i++ {
			win := dom[i : i+ws]

			j := 0
			for j < ws {
				c := rune(win[j])

				repList := s._mapping.GetSimilar(c)
				if len(repList) > 0 {
					for _, g := range repList {
						win = []rune(fmt.Sprintf("%s%c%s", string(win[:j]), g, string(win[j+1:])))
						res = append(res, fmt.Sprintf("%s%s%s.%s", string(dom[:i]), string(win), string(dom[i+ws:]), tld))
					}
				}

				j++
			}
		}
	}

	return helpers.Dedup(res), nil
}