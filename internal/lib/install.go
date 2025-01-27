package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

func Install() {
	fmt.Println("Installing hooks")

	// check if .git exists
	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		fmt.Println("git not initialized")
		return
	}

	// check if .husky exists
	_, err = os.Stat(".husky")

	if os.IsNotExist(err) {
		fmt.Println(".husky not initialized.")
		return
	}

	// check if .husky/hooks exists
	_, err = os.Stat(".husky/hooks")

	if os.IsNotExist(err) {
		fmt.Println("no hooks found")
		return
	}

	root := ".husky/hooks"
	gitDir := ".git/hooks"

	// delete all files in .git/hooks
	err = os.RemoveAll(gitDir)
	if err != nil {
		panic(err)
	}

	// create .git/hooks
	err = os.Mkdir(gitDir, 0755)
	if err != nil {
		panic(err)
	}

	// copy all files in .husky/hooks to .git/hooks
	var hooks []string
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			hooks = append(hooks, path)
			return nil
		})
	if err != nil {
		panic(err)
	}
	for _, hook := range hooks {

		// skip .husky/hooks
		if hook == root {
			continue
		}

		fmt.Println(hook)

		// copy file to .git/hooks
		err = os.Link(hook, filepath.Join(gitDir, filepath.Base(hook)))
		if err != nil {
			panic(err)
		}

		// make file executable
		err = os.Chmod(filepath.Join(gitDir, filepath.Base(hook)), 0755)
		if err != nil {
			panic(err)
		}

	}
	fmt.Println("Hooks installed")
}
