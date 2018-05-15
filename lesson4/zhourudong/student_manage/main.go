package main

type Student struct {
	Id   int    `json:"student_id"`
	Name string `json:"student_name"`
}

var students = make(map[string]Student)

func main() {
	menu()
}
