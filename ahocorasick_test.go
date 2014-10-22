package ahocorasick

import (
	"math/rand"
	"testing"
)

func Test1(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"she", "he", "say", "shr", "her"}
	ac.Build(dictionary)

	ret := ac.Match("yasherhs")
	if len(ret) != 3 || ret[0] != 0 || ret[1] != 1 || ret[2] != 4 {
		t.Fatal()
	}

	ret = ac.Match("yasherhs")
	if len(ret) != 3 || ret[0] != 0 || ret[1] != 1 || ret[2] != 4 {
		t.Fatal()
	}

	if ac.GetMatchResultSize("yasherhs") != 3 {
		t.Fatal()
	}
}

func Test2(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"hello", "世界", "hello世界", "hello"}
	ac.Build(dictionary)

	ret := ac.Match("hello世界")
	if len(ret) != 4 {
		t.Fatal()
	}

	ret = ac.Match("世界")
	if len(ret) != 1 {
		t.Fatal()
	}

	ret = ac.Match("hello")
	if len(ret) != 2 {
		t.Fatal()
	}
}

func Test3(t *testing.T) {
	ac := NewMatcher()

	dictionary := []string{"abc", "bc", "ac", "bc", "de", "efg", "fgh", "hi", "abcd", "ac"}
	ac.Build(dictionary)

	ret := ac.Match("abcdefghij")
	if len(ret) != ac.GetMatchResultSize("abcdefghij") || len(ret) != 8 {
		t.Fatal()
	}

	ret = ac.Match("abcdef")
	if len(ret) != 5 {
		t.Fatal()
	}

	ret = ac.Match("acdejefg")
	if len(ret) != 4 {
		t.Fatal()
	}

	if len(ac.Match("abcd")) != 4 {
		t.Fatal()
	}

	if len(ac.Match("adefacde")) != 3 {
		t.Fatal()
	}

	ret = ac.Match("agbdfgiadafgha")
	if len(ret) != 1 || dictionary[ret[0]] != "fgh" {
		t.Fatal()
	}
}

func Benchmark1(b *testing.B) {
	ac := NewMatcher()

	dictionary := make([]string, 0)
	for i := 0; i < 200000; i++ {
		dictionary = append(dictionary, randWord(2, 6))
	}
	ac.Build(dictionary)

	for i := 0; i < b.N; i++ {
		ret := ac.Match(randWord(5000, 10000))
		if len(ret) > 0 {
		}
	}
}

func Benchmark2(b *testing.B) {
	ac := NewMatcher()

	dictionary := make([]string, 0)
	for i := 0; i < 200000; i++ {
		dictionary = append(dictionary, randWord(2, 6))
	}
	ac.Build(dictionary)

	for i := 0; i < b.N; i++ {
		if ac.GetMatchResultSize(randWord(5000, 10000)) > 0 {
		}
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
