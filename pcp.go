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

type diff struct {
	diffCompare int // -1(B is longer), 0(eqal) or 1 (A is longer)
	diffDomino  string
}

//Use result and current Dif to store for cyclic checking
type result struct {
	potentialResult int
	currentDiff     diff
}

type processInstance struct {
	savedResult  []result
	savedDominos []Domino
}

//index:0  String A ,  index:1 String B.
func (c *processInstance) getString(index int) string {
	retString := ""

	for _, result := range c.savedResult {
		if index == 0 { //DataA
			retString = retString + c.savedDominos[result.potentialResult].DataA
		} else { //DataB
			retString = retString + c.savedDominos[result.potentialResult].DataB
		}
	}
	return retString
}

func (c *processInstance) isResultReach() bool {
	return c.getString(0) == c.getString(1)
}

func (c *processInstance) isCyclicResult() bool {
	if len(c.savedResult) == 0 {
		return false
	}

	checkingRet := c.savedResult[len(c.savedResult)-1]
	for i := 0; i < len(c.savedResult)-1; i++ {
		ret := c.savedResult[i]
		//Find save result list has the same
		if ret.potentialResult == checkingRet.potentialResult && ret.currentDiff == checkingRet.currentDiff {
			return true
		}
	}
	return false
}

// PCP (Post’s Correspondence Problems) contains all input dominos
type PCP struct {
	//Public:
	Dominos []Domino

	//private:
	resultList     []result
	combineDominoA string
	combineDominoB string
	diffDomino     diff
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

func (p *PCP) isDominoValid(curState processInstance, inputDomino Domino) bool {
	strA := curState.getString(0)
	strB := curState.getString(1)

	tempA := strA + inputDomino.DataA
	tempB := strB + inputDomino.DataB

	prefix, exist := getSubsetPrefix(tempA, tempB)
	if !exist {
		return false
	}

	fmt.Println("Domino:", inputDomino, " tempA=", tempA, " tempB=", tempB, " prefix=", prefix)
	return tempA == prefix || tempB == prefix
}

func (p *PCP) checkDiff(curState processInstance) (diff, error) {
	strA := curState.getString(0)
	strB := curState.getString(1)
	//Get which string is longer
	retDiff := diff{}
	retDiff.diffCompare = strings.Compare(strA, strB)

	if retDiff.diffCompare == 0 {
		return diff{}, errors.New("No diff")
	}

	if retDiff.diffCompare == 1 { //A>B
		retDiff.diffDomino = strings.TrimPrefix(strA, strB)
	} else { //A<B
		retDiff.diffDomino = strings.TrimPrefix(strB, strA)
	}

	return retDiff, nil
}

func (p *PCP) applyDomino(curState processInstance, dominoIndex int) (processInstance, error) {
	newDom := p.Dominos[dominoIndex]
	if p.isDominoValid(curState, newDom) {
		newRet := result{}
		newRet.potentialResult = dominoIndex
		if newDiff, err := p.checkDiff(curState); err == nil {
			newRet.currentDiff = newDiff
			curState.savedResult = append(curState.savedResult, newRet)
			return curState, nil
		}

		return processInstance{}, errors.New("Diff error on apply Domino")
	}
	return processInstance{}, errors.New("Domino not valid in apply Domino")
}

func (p *PCP) isSolutionCyclic() bool {
	return false
}

// Add Domino into PCP, will return the domino index.
// ex: p.AddDomino("abc","c") = 1, it means index is 1.
// The index use for solution.
func (p *PCP) AddDomino(strA, strB string) int {
	newDom := Domino{DataA: strA, DataB: strB}
	p.Dominos = append(p.Dominos, newDom)
	return len(p.Dominos) - 1
}

func (p *PCP) recursiveSolve(cur processInstance) (processInstance, error) {
	if cur.isCyclicResult() {
		return processInstance, errors.New(" Cyclic Result .....")
	}

	if cur.isResultReach() {
		return cur, nil
	}

	for index, dom := range p.Dominos {
		if p.isDominoValid(cur, dom) {
			cur = p.applyDomino(cur, index)
			p.recursiveSolve(cur)
		}
	}
	return cur, errors.New("Don't have result")
}

// Try to find solution and return index slice if exist result
// If don't have result error will not be nil
func (p *PCP) FindSolution() ([]int, error) {
	//Use first one as initialize (MPCP)

	process := processInstance{}
	process.savedDominos = p.Dominos

	if retInst, err := p.recursiveSolve(process); err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	fmt.Println("RET : ", retInst.savedResult)
	return nil, nil
}
