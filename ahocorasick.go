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
	// indicates the index of the matching string in the original dictionary
	Index int

	// indicates the ending position index of the matched keyword in the input string s
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
	var ret []*Term

	for index, rune := range s {
		for curNode.child[rune] == nil && curNode != m.root {
			curNode = curNode.fail
		}

		if curNode.child[rune] != nil {
			curNode = curNode.child[rune]
		}

		for p := curNode; p != m.root; p = p.fail {
			if p.count > 0 {
				for i := 0; i < p.count; i++ {
					ret = append(ret, &Term{Index: p.index, EndPosition: index})
				}
			}
		}
	}

	return ret
}

func (m *Matcher) build() {
	ll := list.New()
	ll.PushBack(m.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)

		for i, v := range temp.child {
			if temp == m.root {
				v.fail = m.root
			} else {
				p := temp.fail
				for p != nil {
					if childNode, ok := p.child[i]; ok {
						v.fail = childNode
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
