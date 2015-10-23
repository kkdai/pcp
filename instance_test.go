package pcp_test

import (
	"testing"

	. "github.com/kkdai/pcp"
)

func TestInstance(t *testing.T) {
	p := PCP{}
	p.AddDomino("ab", "a")
	inst := Instance{}
	inst.SavedDominos = append(inst.SavedDominos, Domino{DataA: "ab", DataB: "a"})
	inst, _ = p.ApplyDomino(inst, 0)
	if inst.GetString(0) != "ab" {
		t.Errorf("Get string errors")
	}

	if retList := inst.GetCurrentResult(); len(retList) != 1 {
		t.Errorf("Cannot get result list")
	}
}
