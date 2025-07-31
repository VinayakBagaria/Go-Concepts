package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

func DoWork() {
	path := "."
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	err := tree(path, "")
	if err != nil {
		fmt.Println(err)
	}
}

func tree(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		fmt.Printf("couldn't stat: %s - %v", root, err)
		return err
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("couldn't read dir: %s - %v", root, err)
		return err
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		var add string
		if i == len(names)-1 {
			fmt.Printf(indent + "└── ")
			add = "   "
		} else {
			fmt.Printf(indent + "├── ")
			add = "│   "
		}
		err := tree(filepath.Join(root, name), indent+add)
		if err != nil {
			return err
		}
	}

	return nil
}
