package pcp_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/pcp"
)

func TestInstance(t *testing.T) {
	p := PCP{}
	p.AddDomino("ab", "a")
	inst := Instance{}
	inst.SavedDominos = append(inst.SavedDominos, Domino{DataA: "ab", DataB: "a"})
	inst, _ = p.ApplyDomino(inst, 0)
	fmt.Println(inst.GetString(0))
	fmt.Println(inst.GetString(1))
}

func TestSimplePCP(t *testing.T) {
	p := PCP{}
	p.AddDomino("ab", "b")
	p.AddDomino("b", "a")
	p.AddDomino("a", "ab")

	ret, err := p.FindSolution()
	fmt.Println("err=", err, " ret=", ret)
}
