package main

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}

	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
