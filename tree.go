package main

import (
	"fmt"
	"os"

	"github.com/ahmetb/go-linq"
	"github.com/xlab/treeprint"
)

func main() {
	root := treeprint.New()

	// Добавление файлов и папок к дереву
	addDirToTree(".", root)

	// Вывод дерева в терминал
	fmt.Println(root.String())
}

func addDirToTree(path string, node treeprint.Tree) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	files, _ := file.Readdir(-1)
	linq.From(files).OrderBy(func(i interface{}) interface{} {
		return i.(os.FileInfo).Name()
	}).ForEach(func(file interface{}) {
		if file.(os.FileInfo).IsDir() {
			// Добавление папки к дереву
			subNode := node.AddBranch(file.(os.FileInfo).Name())
			addDirToTree(path+"/"+file.(os.FileInfo).Name(), subNode)
		} else {
			// Добавление файла к дереву
			node.AddNode(file.(os.FileInfo).Name())
		}
	})
}
