package ops

import (
	"cert/cmd/pkg/utils"
	"fmt"
	"io/ioutil"
	"log"
)

func SaveFile(filename string) {

	data := utils.MarshalCert()
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("文件写入失败！\n")
		log.Fatal(err)
	}
}
