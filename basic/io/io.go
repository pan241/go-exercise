package io_

import (
	"bufio"
	"os"
)

// 命令行输入用Reader
func ReadAndPrint() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	text, err := reader.ReadString('\n')
	if err != nil {
		writer.WriteString(err.Error())
	}
	writer.WriteString(text)
}

// 小文件输入用Scanner
func ScanAndPrint() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

}
