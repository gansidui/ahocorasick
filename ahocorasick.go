package ahocorasick

import (
	"container/list"
)

type trieNode struct {
	count int
	fail  *trieNode
	child map[rune]*trieNode
	index int
}

func newTrieNode() *trieNode {
	return &trieNode{
		count: 0,
		fail:  nil,
		child: make(map[rune]*trieNode),
		index: -1,
	}
}

type Matcher struct {
	root *trieNode
	size int
	mark []bool
}

func NewMatcher() *Matcher {
	return &Matcher{
		root: newTrieNode(),
		size: 0,
		mark: make([]bool, 0),
	}
}

// initialize the ahocorasick
func (this *Matcher) Build(dictionary []string) {
	for i, _ := range dictionary {
		this.insert(dictionary[i])
	}
	this.build()
	this.mark = make([]bool, this.size)
}

// string match search
// return all strings matched as indexes into the original dictionary
func (this *Matcher) Match(s string) []int {
	curNode := this.root
	this.resetMark()
	var p *trieNode = nil

	ret := make([]int, 0)

	for _, v := range s {
		for curNode.child[v] == nil && curNode != this.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = this.root
		}

		p = curNode
		for p != this.root && p.count > 0 && !this.mark[p.index] {
			this.mark[p.index] = true
			for i := 0; i < p.count; i++ {
				ret = append(ret, p.index)
			}
			p = p.fail
		}
	}

	return ret
}

// just return the number of len(Match(s))
func (this *Matcher) GetMatchResultSize(s string) int {

	curNode := this.root
	this.resetMark()
	var p *trieNode = nil

	num := 0

	for _, v := range s {
		for curNode.child[v] == nil && curNode != this.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = this.root
		}

		p = curNode
		for p != this.root && p.count > 0 && !this.mark[p.index] {
			this.mark[p.index] = true
			num += p.count
			p = p.fail
		}
	}

	return num
}

func (this *Matcher) build() {
	ll := list.New()
	ll.PushBack(this.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)
		var p *trieNode = nil

		for i, v := range temp.child {
			if temp == this.root {
				v.fail = this.root
			} else {
				p = temp.fail
				for p != nil {
					if p.child[i] != nil {
						v.fail = p.child[i]
						break
					}
					p = p.fail
				}
				if p == nil {
					v.fail = this.root
				}
			}
			ll.PushBack(v)
		}
	}
}

func (this *Matcher) insert(s string) {
	curNode := this.root
	for _, v := range s {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode()
		}
		curNode = curNode.child[v]
	}
	curNode.count++
	curNode.index = this.size
	this.size++
}

func (this *Matcher) resetMark() {
	for i := 0; i < this.size; i++ {
		this.mark[i] = false
	}
}
