package pcp

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

func (c *Instance) GetCurrentResult() []int {
	var retInt []int
	for i := 0; i < len(c.SavedResult); i++ {
		retInt = append(retInt, c.SavedResult[i].PotentialResult)
	}
	return retInt
}
