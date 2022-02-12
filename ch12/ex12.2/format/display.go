package format

import (
	"fmt"
	"reflect"
	"strings"
)

const MAXCOUNT = 100

func display(path string, v reflect.Value, count int) {
	if count > MAXCOUNT {
		fmt.Println("reached recursion limit")
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), count)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), count+1)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			var index string
			switch key.Kind() {
			case reflect.Array:
				index += "["
				elements := []string{}
				for i := 0; i < key.Len(); i++ {
					elements = append(elements, formatAtom(key.Index(i)))
				}
				index += strings.Join(elements, ", ") + "]"
			case reflect.Struct:
				index += key.Type().Name() + " {"
				fields := []string{}
				for i := 0; i < key.NumField(); i++ {
					fields = append(fields, fmt.Sprintf("%s: %s", key.Type().Field(i).Name, key.Field(i)))
				}
				index += strings.Join(fields, ", ") + "}"
			default:
				index = formatAtom(key)
			}
			display(fmt.Sprintf("%s[%s]", path, index), v.MapIndex(key), count)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), count)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), count)
		}
	default: // basic type, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}
