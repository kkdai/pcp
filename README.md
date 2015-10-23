PCP: Post’s Correspondence Problems implement in Golang
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/cyk/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/pcp?status.svg)](https://godoc.org/github.com/kkdai/pcp)  [![Build Status](https://travis-ci.org/kkdai/cyk.svg?branch=master)](https://travis-ci.org/kkdai/pcp)


What is this PCP Problem
=============

The Post correspondence problem is an undecidable decision problem that was introduced by Emil Post in 1946.[1] Because it is simpler than the halting problem and the Entscheidungsproblem it is often used in proofs of undecidability.

(cited from [wiki](https://en.wikipedia.org/wiki/Post_correspondence_problem))
 
 
**Please note it is not total solution until now.**

Installation and Usage
=============


Install
---------------

    go get github.com/kkdai/pcp


Usage
---------------



```go

package main

import (
    "github.com/kkdai/pcp"
)

func main() {
	p := PCP{}
	p.AddDomino("ab", "b")
	p.AddDomino("b", "a")
	p.AddDomino("a", "ab")

	ret, _ := p.FindSolution()
	fmt.Println("Ret=", ret)
}

```

Inspired By
=============

- [Coursera: Automata](https://class.coursera.org/automata-004/)
- [Wiki](https://en.wikipedia.org/wiki/Post_correspondence_problem)
- [Theory of computation / Post’s Correspondence Problems (PCP)](http://www.slideshare.net/ThamerAlamery/theory-of-computation-presentation-final)
- [Github:PCPSolver](https://github.com/dcatteeu/PCPSolver)
- [PCP News](https://webdocs.cs.ualberta.ca/~games/PCP/)
- [Paper: Tackling Post’s Correspondence Problem](https://webdocs.cs.ualberta.ca/~games/PCP/paper/CG2002.pdf)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.
