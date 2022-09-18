package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserData struct {
	Name string
}

func main() {
	d1 := UserData{Name: "Souksaanh"}
	d2 := UserData{Name: "Souksaan 1"}

	objID := primitive.NewObjectID().Hex()
	fmt.Println(d1, d2, objID)
}
