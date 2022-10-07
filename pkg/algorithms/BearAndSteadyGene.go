package algorithms

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const A = "A"
const C = "C"
const T = "T"
const G = "G"
const DIVIDER int = 4

func SteadyGene(gene string) int {
	fmt.Println("This Algorithms resolve the problem \"Bear And Steady Gene\"")
	lenOfGen := len(gene)
	lenOfGenSteady := lenOfGen / DIVIDER
	initialStatusGen := getGenCount(gene)
	minSubstring, gensToInclude := getMinSubstring(initialStatusGen, lenOfGenSteady)
	return getMinimunWithRoteSolution(int(minSubstring), gene, gensToInclude)
}

func getGenCount(gen string) map[string]int {
	result := make(map[string]int)
	for _, val := range gen {
		result[string(val)]++
	}
	return result
}

func getMinSubstring(value map[string]int, steadyValue int) (int, map[string]int) {
	var result int
	var gens = make(map[string]int)

	if (steadyValue - value[A]) < 0 {
		gens[A] = value[A] - steadyValue
		result += gens[A]
	}

	if (steadyValue - value[C]) < 0 {
		gens[C] = value[C] - steadyValue
		result += gens[C]
	}

	if (steadyValue - value[G]) < 0 {
		gens[G] = value[G] - steadyValue
		result += gens[G]
	}

	if (steadyValue - value[T]) < 0 {
		gens[T] = value[T] - steadyValue
		result += gens[T]
	}
	return result, gens
}

func isValidSubString(substring string, substringToCheckOut map[string]int) bool {
	for key, val := range substringToCheckOut {
		if strings.Count(substring, key) != int(val) {
			return false
		}
	}
	return true
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func getMinimunWithRoteSolution(minSubstring int, gene string, gensToInclude map[string]int) int {
	if minSubstring > len(gene) {
		return 0
	}
	finalize := false
	internalMinSubstring := minSubstring
	internalGene := gene

	for !finalize {
		var subGene = internalGene[0:internalMinSubstring]
		if isValidSubString(subGene, gensToInclude) {
			return internalMinSubstring
		} else {
			internalGene = trimFirstRune(internalGene)
		}

		if len(internalGene) < internalMinSubstring {
			internalMinSubstring++
			if internalMinSubstring > len(gene) {
				finalize = true
			} else {
				internalGene = gene
			}
		}
	}

	return 0
}
