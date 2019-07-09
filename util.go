package gocy

import (
	"fmt"
	"reflect"
)

const gocy = "gocy"

func parseStruct(i interface{}, format string) (string, string) {
	iValue := reflect.ValueOf(i)
	iType := reflect.TypeOf(i)
	structName := iType.Name()

	switch iType.Kind() {
	case reflect.Struct:

		fields := ""

		for i := 0; i < iType.NumField(); i++ {
			f := iType.Field(i)
			value := iValue.Field(i).Interface()

			if key, ok := f.Tag.Lookup(gocy); ok && value != "" {
				fields += fmt.Sprintf(`%s: "%s"`, key, value)
			}
		}

		if fields != "" {
			fields = fmt.Sprintf("{%s}", fields)
		}

		return fmt.Sprintf(format,
			fmt.Sprintf(
				"%[1]s:%[1]s%[2]s",
				structName,
				fields,
			)), structName

	default:
		//TODO
		return "", structName
	}
}
