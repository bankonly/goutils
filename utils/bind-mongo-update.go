package utils

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BindUpdate(data interface{}) primitive.D {
	v := reflect.ValueOf(data)
	typeOfS := v.Type()

	result := bson.D{{Key: "updated_at", Value: primitive.NewDateTimeFromTime(time.Now())}}

	for i := 0; i < v.NumField(); i++ {
		f := typeOfS.Field(i)
		if f.Name != "DefaultField" {
			field := typeOfS.Field(i).Tag.Get("bson")
			val := v.Field(i).Interface()
			result = append(result, bson.E{Key: field, Value: val})
		}
	}
	return result
}
