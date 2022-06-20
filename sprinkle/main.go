package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets " + otherWord,
}

func main() {
	// 乱数の元となる値を現在時刻から生成
	rand.Seed(time.Now().UTC().UnixNano())
	// 標準入力のストリームからデータを読み込むbufio.Scannerオブジェクト
	s := bufio.NewScanner(os.Stdin)
	// 次のブロックのデータが入力元から読み込まれ、データがあればtrueを返し、forループの継続条件に使われる
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		// bufio.ScannerのTextメソッドは読み込まれたバイト列(スライス)を文字列に変換する
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
