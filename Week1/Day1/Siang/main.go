package main

import (
	"fmt"
	"reflect"
)

type Users struct {
	Name     string `required:"true"`
	Username string `required:"true"`
	Role     string `required:"true"`
}

func ValidateStruct(s any) error {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()
			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
	}
	return nil
}

func main() {
	newUser := Users{
		Name:     "Agus",
		Username: "Agus123",
		Role:     "Admin",
	}

	userValue := reflect.ValueOf(newUser)
	fmt.Println(userValue)
	fmt.Println(userValue.Kind())

	userType := reflect.TypeOf(newUser)
	fmt.Println(userType)

	var num float64 = 21.98
	reflectValue := reflect.ValueOf(num)
	fmt.Println(reflectValue.Kind())
	reflectType := reflect.TypeOf(num)
	fmt.Println(reflectType)
}
