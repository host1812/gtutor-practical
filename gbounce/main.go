package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const fieldSizeX = 15
const fieldSizeY = 30

type board struct {
	field            [fieldSizeX][fieldSizeY]bool
	x, y, dirx, diry int
}

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

func update2(f *[fieldSizeX][fieldSizeY]bool, x, y, xdir, ydir *int) {
	f[*x][*y] = false
	if *x > 0 && *x < fieldSizeX-1 {
		*x += *xdir
	} else {
		*xdir = -*xdir
		*x += *xdir
	}

	if *y > 0 && *y < fieldSizeY-1 {
		*y += *ydir
	} else {
		*ydir = -*ydir
		*y += *ydir
	}

	f[*x][*y] = true
}

func update3(b *board) {
	b.field[b.x][b.y] = false
	if b.x > 0 && b.x < fieldSizeX-1 {
		b.x += b.dirx
	} else {
		b.dirx = -b.dirx
		b.x += b.dirx
	}

	if b.y > 0 && b.y < fieldSizeY-1 {
		b.y += b.diry
	} else {
		b.diry = -b.diry
		b.y += b.diry
	}

	b.field[b.x][b.y] = true
}

func main() {
	rand.Seed(time.Now().UnixMilli())
	// xdir, ydir := 1, 1
	// field := [fieldSizeX][fieldSizeY]bool{}
	// field[x][y] = true
	b := board{
		field: [fieldSizeX][fieldSizeY]bool{},
		dirx:  1,
		diry:  1,
		x:     rand.Intn(fieldSizeX),
		y:     rand.Intn(fieldSizeY),
	}
	b.field[b.x][b.y] = true
	for {
		clear()

		render(&b.field)
		// x, y, xdir, ydir = update(&field, x, y, xdir, ydir)
		// update2(&field, &x, &y, &xdir, &ydir)
		update3(&b)

		time.Sleep(200 * time.Millisecond)
	}
}
