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
}

type Term struct {
	Index       int
	EndPosition int
}

func NewMatcher() *Matcher {
	return &Matcher{
		root: newTrieNode(),
		size: 0,
	}
}

func BuildNewMatcher(dictionary []string) *Matcher {
	m := &Matcher{
		root: newTrieNode(),
		size: 0,
	}
	m.Build(dictionary)
	return m
}

// initialize the ahocorasick
func (m *Matcher) Build(dictionary []string) {
	for i := range dictionary {
		m.insert(dictionary[i])
	}
	m.build()
}

// string match search
// return all strings matched as indexes into the original dictionary and their positions on matched string
func (m *Matcher) Match(s string) []*Term {
	curNode := m.root
	var p *trieNode = nil

	mark := make([]bool, m.size)
	ret := make([]*Term, 0)

	for index, rune := range s {
		for curNode.child[rune] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[rune]
		if curNode == nil {
			curNode = m.root
		}

		p = curNode
		for p != m.root && p.count > 0 && !mark[p.index] {
			mark[p.index] = true
			for i := 0; i < p.count; i++ {
				ret = append(ret, &Term{Index: p.index, EndPosition: index})
			}
			p = p.fail
		}
	}

	return ret
}

// just return the number of len(Match(s))
func (m *Matcher) GetMatchResultSize(s string) int {
	curNode := m.root
	var p *trieNode = nil

	mark := make([]bool, m.size)
	num := 0

	for _, v := range s {
		for curNode.child[v] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = m.root
		}

		p = curNode
		for p != m.root && p.count > 0 && !mark[p.index] {
			mark[p.index] = true
			num += p.count
			p = p.fail
		}
	}

	return num
}

func (m *Matcher) build() {
	ll := list.New()
	ll.PushBack(m.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)
		var p *trieNode = nil

		for i, v := range temp.child {
			if temp == m.root {
				v.fail = m.root
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
					v.fail = m.root
				}
			}
			ll.PushBack(v)
		}
	}
}

func (m *Matcher) insert(s string) {
	curNode := m.root
	for _, v := range s {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode()
		}
		curNode = curNode.child[v]
	}
	curNode.count++
	curNode.index = m.size
	m.size++
}
