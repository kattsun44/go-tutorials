// 標準入力から2回以上現れる行を出現回数と一緒に表示する
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 標準入力を終わらせる方法が分からないので空文字列が来ると break させるようにする
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
