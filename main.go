package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"unicode/utf8"
)

func failed(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	emoji := "👋"
	r, _ := utf8.DecodeRuneInString(emoji)

	// 小文字の16進数文字列に変換する
	codepoint := fmt.Sprintf("%04x", r)

	// svg 画像をダウンロードする
	res, err := http.Get(fmt.Sprintf("https://jdecked.github.io/twemoji/v/latest/svg/%s.svg", codepoint))
	if err != nil {
		failed(err)
	}
	if res.StatusCode != http.StatusOK {
		failed(fmt.Errorf("failed to download emoji: %s", res.Status))
	}

	// ファイルに保存する
	file, err := os.Create(fmt.Sprintf("%s.svg", emoji))
	if err != nil {
		failed(err)
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		failed(err)
	}
}
