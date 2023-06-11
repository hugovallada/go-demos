package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name     string `required:"true"`
	Age      uint8  `required:"true"`
	Document string `required:"true"`
}

func validateFields(stc any) error {
	stcType := reflect.TypeOf(stc)
	stcValue := reflect.ValueOf(stc)

	for i := 0; i < stcType.NumField(); i++ {
		typeField := stcType.Field(i)
		required := typeField.Tag.Get("required")
		if required == "" || required == "false" {
			continue
		}
		valueField := stcValue.Field(i)
		switch valueField.Kind() {
		case reflect.String:
			if valueField.String() == "" {
				return fmt.Errorf("o campo %s é obrigatório", typeField.Name)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if valueField.Uint() == 0 {
				return fmt.Errorf("o campo %s é obrigatório", typeField.Name)
			}
		}
	}
	return nil
}

func main() {
	p := Person{
		Name:     "Hugo",
		Age:      27,
		Document: "2929828281",
	}

	err := validateFields(p)
	if err != nil {
		panic(err)
	}
}
