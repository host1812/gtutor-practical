package main

import (
	"fmt"
)

func main() {
	// start, stop := 'A', 'z'
	// for i := start; i <= stop; i++ {
	// 	fmt.Printf("%d => %c\n", i, rune(i))
	// }
	text := `
	Veckan före deras avresa till Arrakis, när allt det sista tjafset hade nått en nästan outhärdlig frenesi, kom en gammal kära för att besöka pojkens mamma, Paul.
Det var en varm natt på slottet Caladan, och den uråldriga stenhögen som hade tjänat Atreides-familjen som hem i tjugosex generationer bar den kyliga svettkänslan den fick innan vädret förändrades.
Den gamla kvinnan släpptes in genom sidodörren nedför den välvda gången vid Pauls rum och hon fick en stund titta in på honom där han låg i sin säng.
`
	// fmt.Println(len(text))
	// fmt.Println(len([]byte(text)))
	// fmt.Println(len([]rune(text)))

	for _, r := range text {
		// r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)
		// i += size
	}
}
