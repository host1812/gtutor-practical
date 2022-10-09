package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const offStr = "─ "
const filler = "│   "

func readFile(p string, baseName string, offset int, isLast bool) {
	pref := "├─"
	if offset < 2 {
		fmt.Printf("%s%s\n", strings.Repeat(pref+offStr, offset), baseName)
	} else {
		fmt.Printf("%s%s%s\n", strings.Repeat(filler, offset), strings.Repeat(pref+offStr, offset), baseName)
	}

	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		fmt.Printf("%s - doesn't exists\n", p)
		os.Exit(1)
	}
	dirEntries, _ := os.ReadDir(p)

	for i, e := range dirEntries {
		if i == len(dirEntries)-1 {
			pref = "└─"
			isLast = true
		}
		if e.IsDir() {
			readFile(path.Join(p, e.Name()), e.Name(), offset+1, isLast)

		} else {
			fmt.Printf("%s%s%s\n", strings.Repeat(filler, offset), strings.Repeat(pref+offStr, offset), e.Name())
		}
	}
	// pref := "├─"
	// 	// if i == 0 {
	// 	// 	pref = "└"
	// 	// } else if i == len(dirEntries) {
	// 	// 	pref = "└"
	// 	// }
	// 	if e.IsDir() {
	// 		fmt.Printf("%s%s %s(d)\n", pref, strings.Repeat(offStr, offset), e.Name())
	// 		readFile(path.Join(dir, e.Name()), offset+1)
	// 	} else {
	// 		fmt.Printf("%s%s %s(f)\n", pref, strings.Repeat(offStr, offset), e.Name())
	// 	}
	// }
	// if fs.IsDir() {
	// 	// fmt.Printf("%s - is not a directory\n", dir)
	// 	fmt.Printf("%s- %s (d)", strings.Repeat(offStr, offset), fs.Name())
	// 	readFile(fs.Name(), offset+1)
	// }
	// fmt.Printf("%s- %s (f)", strings.Repeat(offStr, offset), fs.Name())
	// os.ReadDir(dir)
}

type Entry interface {
	Print() string
	IsDir() bool
}

type File struct {
	Name string
	Size int
}

func (f *File) Print() string {
	return fmt.Sprintf("%s(f)", f.Name)
}

func (f *File) IsDir() bool {
	return false
}

type Directory struct {
	Name    string
	Entries []Entry
}

func (d *Directory) Print() string {
	return fmt.Sprintf("%s(d)", d.Name)
}

func (d *Directory) IsDir() bool {
	return true
}

func main() {
	// readFile("./generated", "./generated", 0, false)
	dir0 := Directory{
		Name: "dir00",
		Entries: []Entry{
			&Directory{
				Name: "dir10",
				Entries: []Entry{
					&File{Name: "file11"},
					&File{Name: "file12"},
				},
			},
			&Directory{
				Name: "dir20",
				Entries: []Entry{
					&File{Name: "file21"},
					&File{Name: "file22"},
				},
			},
			// &File{Name: "file01"},
			// &File{Name: "file02"},
		},
	}
	fmt.Println(dir0.Print())
	for _, e := range dir0.Entries {
		fmt.Println("   ", e.Print())
		if e.IsDir() {
			for _, e1 := range e.(*Directory).Entries {
				fmt.Println("       ", e1.Print())
			}
		}
	}

}
