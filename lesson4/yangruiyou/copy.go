package main

import (
	"flag"
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create((dst))
	defer dstFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return io.Copy(dstFile, srcFile)

}

func copyFileAction(src, dst string, showProgress, force bool) {
	if !force {
		if fileExists(dst) {

			fmt.Printf("%s exists override?y/n\n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()

			if strings.TrimSpace(string(data)) != "y" {
				return
			}

		}
	}
	copyFile(src, dst)

	if showProgress{

	}
}

func main() {

	var showProgress, force bool

	flag.BoolVar(&showProgress, "f", false, "force copy when existing")
	flag.BoolVar()
}
