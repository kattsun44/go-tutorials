// 入力に2回以上現れた行の数とその行のテキストを表示する
// 読み込みは標準入力もしくはコマンドライン引数で名前が指定されたファイルの一覧から行う
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n >= 2 {
			fmt.Printf("%d\t\"%s\"\n", n, line)
		}
	}
}

func countLines(s *os.File, counts map[string]int) {
	input := bufio.NewScanner(s)
	for input.Scan() {
		counts[input.Text()]++
	}
}
