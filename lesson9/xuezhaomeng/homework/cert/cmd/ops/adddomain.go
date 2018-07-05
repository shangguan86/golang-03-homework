package ops

import (
	"cert/cmd/pkg/utils"
)

func CreateCert(domainname, regdate, expdate string) {
	utils.CertInfoMap[domainname] = &utils.CertInfo{
		DomainName: domainname,
		RegDate:    regdate,
		ExpDate:    expdate,
	}

}
