package ahocorasick

import (
	"math/rand"
	"testing"
)

func Test1(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"she", "he", "say", "shr", "her"}
	ac.Build(dictionary)

	expected := []*Term{
		{Index: 0, EndPosition: 5},
		{Index: 1, EndPosition: 5},
		{Index: 4, EndPosition: 6},
	}

	s := "yasherhs"
	ret := ac.Match(s)
	if len(expected) != len(ret) {
		t.Fatal()
	}
	for i, _ := range ret {
		if ret[i].Index != expected[i].Index || ret[i].EndPosition != expected[i].EndPosition {
			t.Fatal()
		}

		original := dictionary[ret[i].Index]
		matched := s[ret[i].EndPosition-len(original) : ret[i].EndPosition]
		if original != matched {
			t.Fatal()
		}
	}
}

func Test2(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"中国人民", "国人", "中国人", "hello世界", "hello"}
	ac.Build(dictionary)

	if len(ac.Match("中国人")) != 2 {
		t.Fatal()
	}
	if len(ac.Match("世界")) != 0 {
		t.Fatal()
	}

	s := "hello世界"
	ret := ac.Match(s)
	if len(ret) != 2 {
		t.Fatal()
	}

	for i, _ := range ret {
		original := dictionary[ret[i].Index]
		matched := s[ret[i].EndPosition-len(original) : ret[i].EndPosition]
		if original != matched {
			t.Fatal()
		}
	}
}

func Benchmark(b *testing.B) {
	ac := NewMatcher()

	dictionary := make([]string, 0)
	for i := 0; i < 200000; i++ {
		dictionary = append(dictionary, randWord(2, 6))
	}
	ac.Build(dictionary)

	for i := 0; i < b.N; i++ {
		ac.Match(randWord(5000, 10000))
	}
}

func randWord(m, n int) string {
	num := rand.Intn(n-m) + m
	var s string
	var a rune = 'a'

	for i := 0; i < num; i++ {
		c := a + rune(rand.Intn(26))
		s = s + string(c)
	}

	return s
}
