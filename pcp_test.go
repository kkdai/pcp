package pcp_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/pcp"
)

func TestFirstPCP(t *testing.T) {
	p := PCP{}
	p.AddDomino("ab", "b")
	p.AddDomino("b", "a")
	p.AddDomino("a", "ab")

	ret, err := p.FindSolution()
	if err != nil {
		t.Errorf("Error occur not found result")
	}
	if !(ret[0] == 2 && ret[1] == 1 && ret[2] == 0) {
		t.Errorf("Result is wrong, get %v\n", ret)
	}
	fmt.Println("1: ret=", ret)
}

func TestSecnodPCP(t *testing.T) {
	p := PCP{}
	p.AddDomino("01110", "011")
	p.AddDomino("101", "0101")
	p.AddDomino("1110", "10111")

	ret, err := p.FindSolution()
	if err != nil {
		t.Errorf("Error occur not found result")
	}

	fmt.Println("2: ret=", ret)
}

// Case: More domino cause the diff longer.
// Still under progress.

// func TestThirdPCP(t *testing.T) {
// 	p := PCP{}
// 	p.AddDomino("a", "ac")
// 	p.AddDomino("c", "ba")
// 	p.AddDomino("ba", "a")
// 	p.AddDomino("acb", "b")

// 	ret, err := p.FindSolution()
// 	if err != nil {
// 		t.Errorf("Error occur not found result")
// 	}

// 	fmt.Println("3: ret=", ret)
// }

// Case: More domino cause the diff longer.
// Still under progress.
// func Test4thPCP(t *testing.T) {
// 	p := PCP{}
// 	p.AddDomino("bb", "b")
// 	p.AddDomino("ab", "ba")
// 	p.AddDomino("c", "bc")

// 	ret, err := p.FindSolution()
// 	if err != nil {
// 		t.Errorf("Error occur not found result")
// 	}
// 	fmt.Println("1: ret=", ret)
// }
