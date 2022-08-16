package main

import (
	"fmt"
)

// Create a Struct of type Student
type Student struct {
	id      int
	name    string
	subject string
}

func (obj Student) getinfo() {
	fmt.Println(obj.id, obj.name, obj.subject)

}

func (obj *Student) setinfo() { //obj==> Adress of Student
	obj.id = 1
	obj.name = "Manzoor Faisal"
	obj.subject = "CS"

}
func main() {
	stu := Student{}
	stu.setinfo()
	stu.getinfo()

}
