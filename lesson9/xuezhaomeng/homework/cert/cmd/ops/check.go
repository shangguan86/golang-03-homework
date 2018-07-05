package ops

import (
	"cert/cmd/pkg/utils"
	"fmt"
)

func Monitor() {
	// for k, v := range utils.CertInfoMap {
	// 	longago := utils.Get30goDate()
	// 	if v.ExpDate <= longago {
	// 		fmt.Printf("域名 %s 证书即将到期，到期时间 %s \n", k, v.ExpDate)
	// 	}

	// }

	taskChan := make(chan *utils.CertInfo, 100)
	resChan := make(chan *utils.CertInfo, 100)

	go func() {
		for _, v := range utils.CertInfoMap {
			//fmt.Println(fmt.Sprintf("%T", v))
			taskChan <- v
		}
		close(taskChan)
	}()

	for i := 0; i < 4; i++ {
		go Exec(taskChan, resChan)
	}
	for {
		select {
		case v := <-resChan:
			//fmt.Println(v)
			message := fmt.Sprintf("域名 %s 证书即将到期，到期时间 %s \n", v.DomainName, v.RegDate)
			fmt.Println(message)
			subject := "您的证书即将过期"
			email := utils.NewEmail("xiaoxuexue@lianluo.com", subject, message)
			err := utils.SendEmail(email)
			if err != nil {
				fmt.Println("邮件发送失败！")
			}
		default:
			//fmt.Println("sleep")
			//time.Sleep(time.Second * 2)
		}

	}
}

func Exec(taskChan chan *utils.CertInfo, resChan chan *utils.CertInfo) {
	longago := utils.Get30goDate()

	for v := range taskChan {
		//fmt.Println(fmt.Sprintf("%T", v))
		if v.ExpDate <= longago {
			resChan <- v
		}

	}
}
