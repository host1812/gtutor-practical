package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type element struct {
	sprite    [][]string
	separator bool
	blinked   bool
}

func initialize() []element {
	zero := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{"▒", " ", "▒"},
			{"▒", " ", "▒"},
			{"▒", " ", "▒"},
			{"▒", "▒", "▒"},
		},
	}
	one := element{
		sprite: [][]string{
			{" ", " ", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
		},
	}
	two := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{"▒", "▒", "▒"},
			{"▒", " ", " "},
			{"▒", "▒", "▒"},
		},
	}
	three := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{"▒", "▒", "▒"},
		},
	}
	four := element{
		sprite: [][]string{
			{"▒", " ", "▒"},
			{"▒", " ", "▒"},
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
		},
	}
	five := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{"▒", " ", " "},
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{"▒", "▒", "▒"},
		},
	}
	six := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{"▒", " ", " "},
			{"▒", "▒", "▒"},
			{"▒", " ", "▒"},
			{"▒", "▒", "▒"},
		},
	}
	seven := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
		},
	}
	eight := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{"▒", " ", "▒"},
			{"▒", "▒", "▒"},
			{"▒", " ", "▒"},
			{"▒", "▒", "▒"},
		},
	}
	nine := element{
		sprite: [][]string{
			{"▒", "▒", "▒"},
			{"▒", " ", "▒"},
			{"▒", "▒", "▒"},
			{" ", " ", "▒"},
			{" ", " ", "▒"},
		},
	}
	sep := element{
		sprite: [][]string{
			{" ", " ", " "},
			{" ", "░", " "},
			{" ", " ", " "},
			{" ", "░", " "},
			{" ", " ", " "},
		},
		separator: true,
		blinked:   false,
	}
	sepB := element{
		sprite: [][]string{
			{" ", " ", " "},
			{" ", "▓", " "},
			{" ", " ", " "},
			{" ", "▓", " "},
			{" ", " ", " "},
		},
		separator: true,
		blinked:   true,
	}
	sprites := []element{zero, one, two, three, four, five, six,
		seven, eight, nine, sep, sepB}
	return sprites
}

func main() {
	sprites := initialize()
	blinked := false
	for {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		now := time.Now()
		timeString := fmt.Sprintf("%02d:%02d:%02d",
			now.Hour(), now.Minute(), now.Second())
		fmt.Println(timeString)

		// Render each sprite line by line

		for i := 0; i < 5; i++ {
			for _, d := range timeString {
				// fmt.Print(sprites[d-48].sprite[i])
				s := sprites[d-48]
				if s.separator {
					if !blinked {
						s = sprites[d-48+1]
					}
				}
				for _, c := range s.sprite[i] {
					fmt.Printf("%s", c)
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
		time.Sleep(1000 * time.Millisecond)
		blinked = !blinked
	}
}
