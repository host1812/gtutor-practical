package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const fieldSizeX = 25
const fieldSizeY = 150

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error running clear command, err: %s", err)
		os.Exit(1)
	}
}

func render(f *[fieldSizeX][fieldSizeY]bool) {
	for x := 0; x < fieldSizeX; x++ {
		for y := 0; y < fieldSizeY; y++ {
			if f[x][y] {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
}

func update(f *[fieldSizeX][fieldSizeY]bool, x, y, xdir, ydir int) (int, int, int, int) {
	f[x][y] = false
	if x > 0 && x < fieldSizeX-1 {
		x += xdir
	} else {
		xdir = -xdir
		x += xdir
	}

	if y > 0 && y < fieldSizeY-1 {
		y += ydir
	} else {
		ydir = -ydir
		y += ydir
	}

	f[x][y] = true
	return x, y, xdir, ydir
}

func main() {
	x, y := rand.Intn(fieldSizeX), rand.Intn(fieldSizeY)
	xdir, ydir := 1, 1
	field := [fieldSizeX][fieldSizeY]bool{}
	field[x][y] = true
	for {
		clear()

		render(&field)
		x, y, xdir, ydir = update(&field, x, y, xdir, ydir)

		time.Sleep(20 * time.Millisecond)
	}
}
