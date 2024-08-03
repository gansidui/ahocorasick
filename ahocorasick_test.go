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
		{Index: 0, EndPosition: 4},
		{Index: 1, EndPosition: 4},
		{Index: 4, EndPosition: 5},
	}

	ret := ac.Match("yasherhs")
	if len(expected) != len(ret) {
		t.Fatal()
	}
	for i, _ := range ret {
		if ret[i].Index != expected[i].Index || ret[i].EndPosition != expected[i].EndPosition {
			t.Fatal()
		}
	}
}

func Test2(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"hello", "世界", "hello世界"}
	ac.Build(dictionary)

	if len(ac.Match("hello世界")) != 3 {
		t.Fatal()
	}
	if len(ac.Match("世界")) != 1 {
		t.Fatal()
	}
	if len(ac.Match("hello")) != 1 {
		t.Fatal()
	}
}

func Test3(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"中国人民", "国人", "中国人", "hello世界", "hello"}
	ac.Build(dictionary)

	if len(ac.Match("中国人")) != 2 {
		t.Fatal()
	}
	if len(ac.Match("世界")) != 0 {
		t.Fatal()
	}
	if len(ac.Match("hello世界")) != 2 {
		t.Fatal()
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
