package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type CertInfo struct {
	DomainName string
	RegDate    string
	ExpDate    string
}

var CertInfoMap = make(map[string]*CertInfo)

//获取当前时间+30天的时间
func Get30goDate() string {

	now := time.Now()
	longago := now.AddDate(0, 0, +30)
	return fmt.Sprintf(longago.Format("2006-01-02"))

}

func MarshalCert() []byte {

	data, err := json.Marshal(CertInfoMap)
	if err != nil {
		err = fmt.Errorf("json.marshal failed, err:", err)
		log.Fatal(err)
	}
	// fmt.Println(string(data))
	return data

}

func UnMarshalCert() {

	buf, err := ioutil.ReadFile("cert.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(buf, &CertInfoMap)
	if err != nil {
		err = fmt.Errorf("json.marshal failed, err:", err)
		log.Fatal(err)
	}
	// for k, v := range CertInfoMap {
	// 	fmt.Println(k, v)
	// }

}
