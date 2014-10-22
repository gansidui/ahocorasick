##ahocorasick

Aho-Corasick string matching algorithm for golang


~~~ go
package main

import (
	"fmt"
	"github.com/gansidui/ahocorasick"
)

func main() {
	ac := ahocorasick.NewMatcher()

	dictionary := []string{"hello", "world", "世界", "google", "golang", "c++", "love"}

	ac.Build(dictionary)

	ret := ac.Match("hello世界, hello google, i love golang!!!")

	for _, i := range ret {
		fmt.Println(dictionary[i])
	}
}


~~~

##LICENSE

MIT