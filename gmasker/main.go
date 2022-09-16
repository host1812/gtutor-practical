package main

import (
	"fmt"
	"unicode/utf8"
)

func sanitize(s []rune, b *[]byte, i int) (pos int) {
	pos = i
	for _, r := range s {
		if r != ' ' {
			fmt.Printf("debug: need to replace %c => '*' !\n", r)
			*b = append(*b, '*')
			pos++
		} else {
			*b = append(*b, ' ')
			break
		}
	}
	// for ; i < len(b); i++ {
	// 	if b[i] != ' ' {
	// 		//replace
	// 		if utf8.RuneStart(b[i]) {
	// 			fmt.Printf("debug: rune detected!\n")
	// 			utf8.
	// 		// } else {
	// 		// 	b[i] = '*'
	// 		// }
	// 	} else {
	// 		break
	// 	}
	// }
	return pos
}

func main() {
	text := `
Here: http://😀.c😀 Click!
Other: http://testing.test.com need to be sanitized!"
Emojii attack: http://😀😀😀😀.😀😀.😀😀😀 need to be sanitized!"`
	buffer := []byte{}
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%d = %c\n", i, runes[i])
		if (i+7) < len(runes) && string(runes[i:i+7]) == "http://" {
			// buffer = append(buffer, text[:i+7]...)
			fmt.Printf("debug: found 'http://'\n")
			buffer = append(buffer, "http://"...)
			i = sanitize(runes[i+7:], &buffer, i+7)
		} else {
			buffer = utf8.AppendRune(buffer, runes[i])
		}
		// if r == 'h' {
		// 	fmt.Printf("debug: found 'h' at %d\n", i)
		// 	if text[i:i+7] == "http://" {
		// 		fmt.Printf("debug: found 'http://'\n")
		// 		buffer = append(buffer, text[:i+7]...)
		// 		i = sanitize(text[i+7:], &buffer, i+7)
		// 		// i = i + 7
		// 	}
		// } else {
		// 	buffer = append(buffer, text[i])
		// }
	}
	fmt.Println(string(buffer))
}
