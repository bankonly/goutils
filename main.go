package main

import "fmt"

type UserData struct {
	Name string
}

func main() {
	d1 := UserData{Name: "Souksaanh"}
	d2 := UserData{Name: "Souksaan 1"}

	fmt.Println(d1, d2)
}
