package json_encoder

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func encode(buf *bytes.Buffer, v reflect.Value, indent string) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)
	case reflect.Array, reflect.Slice:
		indent += "  "
		buf.WriteString("[\n")
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(indent)
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}

			if i != v.Len()-1 {
				buf.WriteString(",\n")
			} else {
				buf.WriteByte('\n')
			}
		}
		indent = strings.Repeat(" ", len(indent)-2)
		buf.WriteString(indent + "]")
	case reflect.Map:
		indent += "  "
		buf.WriteString("{\n")
		for i, key := range v.MapKeys() {
			buf.WriteString(indent)
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			if i != len(v.MapKeys())-1 {
				buf.WriteString(",\n")
			} else {
				buf.WriteByte('\n')
			}
		}
		indent = strings.Repeat(" ", len(indent)-2)
		buf.WriteString(indent + "}")
	case reflect.Struct:
		indent += "  "
		buf.WriteString("{\n")
		for i := 0; i < v.NumField(); i++ {
			buf.WriteString(indent)
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				buf.WriteString(",\n")
			} else {
				buf.WriteByte('\n')
			}
		}
		indent = strings.Repeat(" ", len(indent)-2)
		buf.WriteString(indent + "}")
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	var indent = ""
	if err := encode(&buf, reflect.ValueOf(v), indent); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
