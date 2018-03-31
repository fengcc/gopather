package main

import (
	"fmt"
	"os"
	"strings"
)

func check() ([]string, bool) {
	paths := os.Getenv("GOPATHER")
	if len(paths) == 0 {
		return nil, false
	}

	if len(os.Args) != 2 {
		return nil, false
	}
	return strings.Split(paths, ":"), true
}

func main() {
	paths, ok := check()
	if !ok {
		// 有错，不做任何设置
		return
	}
	pwd := os.Args[1]

	for _, p := range paths {
		if strings.HasPrefix(pwd, p) {
			// 输出对应的gopath
			fmt.Print(p)
			return
		}
	}
	// 输出对应的gopath
	fmt.Print(paths[0])
}
