package main

type Talk interface {
	Hello(username string) string
	Talk(heard string) (saying string, end bool, err error)
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

type myTalk string

func (talk *myTalk) Hello(username string) string {
	//
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	//

}

type simpleCN struct {
	name string
	talk Talk
}

func main() {
	var talk Talk = new(myTalk)

}
