package pcp

import (
	"errors"
	"fmt"
	"strings"
)

// Domino combine two data together
// Ex: {"ab"/"a"}  DataA= "ab" , DataB="a"
type Domino struct {
	DataA string
	DataB string
}

// Diff: Difference of two compare string
// It will store if A B is equal or the difference string for cyclic checking.
type Diff struct {
	DiffCompare int // -1(B is longer), 0(eqal) or 1 (A is longer)
	DiffDomino  string
}

//Use result and current Dif to store for cyclic checking
type Result struct {
	PotentialResult int
	CurrentDiff     Diff
}

// PCP (Post’s Correspondence Problems) contains all input dominos
type PCP struct {
	//Public:
	Dominos []Domino
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

// To validate the input domino, if it is valid for apply.
func (p *PCP) IsDominoValid(curState Instance, inputDomino Domino) bool {
	strA := curState.GetString(0)
	strB := curState.GetString(1)

	tempA := strA + inputDomino.DataA
	tempB := strB + inputDomino.DataB
	fmt.Println("TempA=", tempA)
	fmt.Println("TempB=", tempB)

	prefix, exist := getSubsetPrefix(tempA, tempB)
	fmt.Println("Pre:= ", prefix)
	if !exist {
		return false
	}

	//fmt.Println("Domino:", inputDomino, " tempA=", tempA, " tempB=", tempB, " prefix=", prefix)
	return tempA == prefix || tempB == prefix
}

// Apply a domino in current status and check the difference.
func (p *PCP) CheckDiff(curState Instance, dom Domino) (Diff, error) {
	strA := curState.GetString(0) + dom.DataA
	strB := curState.GetString(1) + dom.DataB
	//fmt.Println("CheckDiff strA=", strA, " strB=", strB)
	//Get which string is longer
	retDiff := Diff{}
	retDiff.DiffCompare = strings.Compare(strA, strB)

	if retDiff.DiffCompare == 0 {
		return retDiff, nil
	}

	//fmt.Println("Diff=", retDiff.diffCompare, strings.TrimPrefix(strA, strB))

	if retDiff.DiffCompare == 1 { //A>B
		retDiff.DiffDomino = strings.TrimPrefix(strA, strB)
	} else { //A<B
		retDiff.DiffDomino = strings.TrimPrefix(strB, strA)
	}

	//fmt.Println("Check Diff:", retDiff)
	return retDiff, nil
}

func (p *PCP) ApplyDomino(curState Instance, dominoIndex int) (Instance, error) {
	newDom := p.Dominos[dominoIndex]

	if p.IsDominoValid(curState, newDom) {
		//fmt.Println("Dom is valid")
		fmt.Printf("Cur (%v) Add (%d)->\n", curState.GetCurrentResult(), dominoIndex)
		newRet := Result{}
		newRet.PotentialResult = dominoIndex
		if newDiff, err := p.CheckDiff(curState, newDom); err == nil {
			newRet.CurrentDiff = newDiff
			curState.SavedResult = append(curState.SavedResult, newRet)
			return curState, nil
		}

		return Instance{}, errors.New("Diff error on apply Domino")
	}
	return Instance{}, errors.New("Domino not valid in apply Domino")
}

// Add Domino into PCP, will return the domino index.
// ex: p.AddDomino("abc","c") = 1, it means index is 1.
// The index use for solution.
func (p *PCP) AddDomino(strA, strB string) int {
	newDom := Domino{DataA: strA, DataB: strB}
	p.Dominos = append(p.Dominos, newDom)
	return len(p.Dominos) - 1
}

func (p *PCP) recursiveSolve(cur Instance) (Instance, error) {
	if cur.isCyclicResult() {
		fmt.Println("Cyclic result")
		return Instance{}, errors.New(" Cyclic Result .....")
	}

	if cur.isResultReach() {
		fmt.Println("Result find!")
		return cur, nil
	}

	for index, dom := range p.Dominos {
		fmt.Println("dom=", dom, " index=", index)
		if p.IsDominoValid(cur, dom) {
			fmt.Println("Find dom valid=", dom)
			fmt.Printf("[%d]", index)
			cur, _ = p.ApplyDomino(cur, index)
			var err error
			fmt.Println("StrA=", cur.GetString(0))
			fmt.Println("StrB=", cur.GetString(1))
			// if len(cur.GetString(0)) > 10 {
			// 	return cur, errors.New("Too long")
			// }
			cur, err = p.recursiveSolve(cur)
			if err == nil {
				return cur, nil
			}

		}
	}
	return cur, errors.New("Don't have result")
}

// Try to find solution and return index slice if exist result
// If don't have result error will not be nil
func (p *PCP) FindSolution() ([]int, error) {
	process := Instance{}
	process.SavedDominos = p.Dominos
	//fmt.Println("ST:StrA=", process.GetString(0))
	//fmt.Println("ST:StrB=", process.GetString(1))
	//fmt.Println("current dominos=", process.SavedDominos)
	if retInst, err := p.recursiveSolve(process); err != nil {
		fmt.Println("err:", err)
		return nil, err
	} else {
		pcpRet := retInst.GetCurrentResult()
		return pcpRet, nil
	}
}
