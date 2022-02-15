package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func encode(buf *bytes.Buffer, v reflect.Value, indent string) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)
	case reflect.Array, reflect.Slice:
		indent += " "
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(indent)
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}

			if i != v.Len()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		indent += " "
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprint(buf, "\n ")
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+strings.Repeat(" ", len(v.Type().Field(i).Name)+1)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map:
		indent += " "
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(indent)
			}
			buf.WriteByte('(')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte(')')

			if i != len(v.MapKeys())-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// MarshalはGoの値をS式形式でエンコードします
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	var indent = " "
	if err := encode(&buf, reflect.ValueOf(v), indent); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
