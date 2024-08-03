## ahocorasick

Aho-Corasick string matching algorithm for golang

~~~ go
package main

import (
	"fmt"
	"log"

	"github.com/gansidui/ahocorasick"
)

func main() {
	ac := ahocorasick.NewMatcher()

	dictionary := []string{"hello", "world", "世界", "google", "golang", "c++", "love"}
	ac.Build(dictionary)

	s := "hello世界, hello google, i love golang!!!"
	ret := ac.Match(s)

	for i, _ := range ret {
		original := dictionary[ret[i].Index]
		matched := s[ret[i].EndPosition-len(original) : ret[i].EndPosition]
		if original != matched {
			log.Fatal()
		}
		fmt.Println(original, matched)
	}
}
~~~

## LICENSE

MIT
