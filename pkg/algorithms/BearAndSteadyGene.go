package algorithms

import (
	"fmt"
	"strings"
)

const A = "A"
const C = "C"
const T = "T"
const G = "G"
const DIVIDER int32 = 4

func SteadyGene(gene string) int32 {
	fmt.Println("This Algorithms resolve the problem \"Bear And Steady Gene\"")
	lenOfGen := int32(len(gene))
	lenOfGenSteady := lenOfGen / DIVIDER
	initialStatusGen := getGenCount(gene)

	// This is the min substring
	minSubstring, gensToInclude := getMinSubstring(initialStatusGen, lenOfGenSteady)

	//fmt.Printf("Len of Gen: %d. Len of Gen Steady: %d. Gene Original: %s. Gene Final: %s. MaxIndex:%d, MinIndex:%d MinSubString: %d \n", lenOfGen, lenOfGenSteady, gene, genFinal, maxIndex, minIndex, minLenSubString)
	return getMinSubStringRecursive(minSubstring, gene, gensToInclude)
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

	aValue, aExists := substringToCheckOut[A]
	cValue, cExists := substringToCheckOut[C]
	gValue, gExists := substringToCheckOut[G]
	tValue, tExists := substringToCheckOut[T]

	return (aExists && gensInclude[A] == aValue) ||
		(cExists && gensInclude[C] == cValue) ||
		(gExists && gensInclude[G] == gValue) ||
		(tExists && gensInclude[T] == tValue)
}

func getMinSubStringRecursive(minSubstring int32, gene string, gensToInclude map[string]int32) int32 {

	if minSubstring > int32(len(gene)) {
		return 0
	}

	for index, _ := range gene {
		if (int32(index) + minSubstring) >= int32(len(gene)) {
			break
		} else {
			var subGene = gene[index : minSubstring+int32(index)]
			if isValidSubString(subGene, gensToInclude) {
				return minSubstring
			}
		}
	}

	return getMinSubStringRecursive(minSubstring+1, gene, gensToInclude)
}
