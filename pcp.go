package pcp

type Domino struct {
	DataA string
	DataB string
}

type PCP struct {
	//Public:
	Dominos []Domino

	//private:
	resultList     []int
	combineDominoA string
	combineDominoB string
}

func contrainPrefix(str1, str2 string) bool {
	if sub, find := getSubsetPrefix(str1, str2); find {
		return sub == str2
	}
	return false
}

func getSubsetPrefix(str1, str2 string) (string, bool) {
	findSubset := false
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			retStr := str1[:i]
			return retStr, findSubset
		}
		findSubset = true
	}

	if len(str1) > len(str2) {
		return str2, findSubset
	} else if len(str1) == len(str2) {
		//fix "" not a subset of ""
		return str1, str1 == str2
	}

	return str1, findSubset
}

// Add Domino into PCP, will return the domino index.
// ex: p.AddDomino("abc","c") = 1, it means index is 1.
// The index use for solution.
func (p *PCP) AddDomino(strA, strB string) int {
}

// Try to find solution and return index slice if exist result
// If don't have result error will not be nil
func (p *PCP) FindSolution() ([]int, error) {
}
