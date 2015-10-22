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

type Diff struct {
	DiffCompare int // -1(B is longer), 0(eqal) or 1 (A is longer)
	DiffDomino  string
}

//Use result and current Dif to store for cyclic checking
type Result struct {
	PotentialResult int
	CurrentDiff     Diff
}

type Instance struct {
	SavedResult  []Result
	SavedDominos []Domino
}

//index:0  String A ,  index:1 String B.
func (c *Instance) GetString(index int) string {
	retString := ""

	for _, result := range c.SavedResult {
		if index == 0 { //DataA
			retString = retString + c.SavedDominos[result.PotentialResult].DataA
		} else { //DataB
			retString = retString + c.SavedDominos[result.PotentialResult].DataB
		}
	}
	return retString
}

func (c *Instance) isResultReach() bool {
	if len(c.GetString(0)) == 0 && len(c.GetString(1)) == 0 {
		return false
	}
	return c.GetString(0) == c.GetString(1)
}

func (c *Instance) isCyclicResult() bool {
	if len(c.SavedResult) == 0 {
		return false
	}

	checkingRet := c.SavedResult[len(c.SavedResult)-1]
	for i := 0; i < len(c.SavedResult)-1; i++ {
		ret := c.SavedResult[i]
		//Find save result list has the same
		if ret.PotentialResult == checkingRet.PotentialResult && ret.CurrentDiff == checkingRet.CurrentDiff {
			return true
		}
	}
	return false
}

// PCP (Post’s Correspondence Problems) contains all input dominos
type PCP struct {
	//Public:
	Dominos []Domino
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

func (p *PCP) IsDominoValid(curState Instance, inputDomino Domino) bool {
	strA := curState.GetString(0)
	strB := curState.GetString(1)

	tempA := strA + inputDomino.DataA
	tempB := strB + inputDomino.DataB
	//fmt.Print("TempA=", tempA, " TempB:", tempB)

	prefix, exist := getSubsetPrefix(tempA, tempB)
	if !exist {
		return false
	}

	//fmt.Println("Domino:", inputDomino, " tempA=", tempA, " tempB=", tempB, " prefix=", prefix)
	return tempA == prefix || tempB == prefix
}

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
	fmt.Println("Got dom:", newDom)

	if p.IsDominoValid(curState, newDom) {
		fmt.Println("Dom is valid")
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
		return Instance{}, errors.New(" Cyclic Result .....")
	}

	if cur.isResultReach() {
		fmt.Println("Result find!")
		return cur, nil
	}

	for index, dom := range p.Dominos {
		fmt.Println("dom=", dom)
		if p.IsDominoValid(cur, dom) {
			fmt.Println("Find dom valid=", dom)

			cur, _ = p.ApplyDomino(cur, index)
			var err error
			fmt.Println("StrA=", cur.GetString(0))
			fmt.Println("StrB=", cur.GetString(1))
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
	//Use first one as initialize (MPCP)

	process := Instance{}
	process.SavedDominos = p.Dominos
	fmt.Println("ST:StrA=", process.GetString(0))
	fmt.Println("ST:StrB=", process.GetString(1))
	fmt.Println("current dominos=", process.SavedDominos)
	if retInst, err := p.recursiveSolve(process); err != nil {
		fmt.Println("err:", err)
		return nil, err
	} else {
		var retInt []int
		for i := 0; i < len(retInst.SavedResult); i++ {
			//aretInt = append(retInt, retInt.SavedResult[i].PotentialResult)
		}
		return retInt, nil
	}
}
