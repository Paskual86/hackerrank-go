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
const DIVIDER int32 = 4

func SteadyGene(gene string) int {
	fmt.Println("This Algorithms resolve the problem \"Bear And Steady Gene\"")
	lenOfGen := int32(len(gene))
	lenOfGenSteady := lenOfGen / DIVIDER
	initialStatusGen := getGenCount(gene)

	// This is the min substring
	minSubstring, gensToInclude := getMinSubstring(initialStatusGen, lenOfGenSteady)
	return getMinimunWithRoteSolution(int(minSubstring), gene, gensToInclude)
}

func getGenCount(gen string) map[string]int32 {
	result := make(map[string]int32)
	result[A] = int32(strings.Count(gen, A))
	result[C] = int32(strings.Count(gen, C))
	result[G] = int32(strings.Count(gen, G))
	result[T] = int32(strings.Count(gen, T))

	return result
}

func getMinSubstring(value map[string]int32, steadyValue int32) (int32, map[string]int32) {
	var result int32
	var gens = make(map[string]int32)

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

func isValidSubString(substring string, substringToCheckOut map[string]int32) bool {
	var gensInclude = getGenCount(substring)

	for key, val := range substringToCheckOut {
		if gensInclude[key] != val {
			return false
		}
	}
	return true
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func getMinimunWithRoteSolution(minSubstring int, gene string, gensToInclude map[string]int32) int {
	if minSubstring > len(gene) {
		return 0
	}
	found := false
	var counter int
	internalMinSubstring := minSubstring
	internalGene := gene
	for !found {
		var subGene = internalGene[0:internalMinSubstring]
		if isValidSubString(subGene, gensToInclude) {
			found = true
			return internalMinSubstring
		} else {
			fistGen := internalGene[0]
			internalGene = trimFirstRune(internalGene) + string(fistGen)
			counter++
		}

		if counter == len(internalGene) {
			counter = 0
			internalMinSubstring++
		}

		if internalMinSubstring == len(internalGene) {
			found = true
		}
	}
	return 0
}
