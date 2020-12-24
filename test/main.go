package main

import (
	"fmt"
)

func main() {
	fmt.Println("start!")
	fmt.Println(Test2())
	d := Data{}
	d.Test2 = 10
	d.Output()
	d.OutputCopy()
	fmt.Println("x:", d.Test2)
}
func Test2() string {
	return "Value!"
}

type Data struct {
	x      int
	Test2  int
	output int
}

func (d *Data) Output() {
	d.Test2 += 10
}
func (d Data) OutputCopy() {
	d.Test2 += 10
}
