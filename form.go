package form

import (
	"github.com/iancoleman/strcase"
	"net/url"
	"reflect"
	"strconv"
)

// MarshalForm returns url.Values created from v struct fields
func MarshalForm(v interface{}) url.Values {
	form := url.Values{}
	t := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		alias := field.Tag.Get("form")
		if alias == "-" {
			continue
		}
		if alias == "" {
			alias = strcase.ToLowerCamel(field.Name)
		}
		switch field.Type.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := strconv.FormatInt(values.Field(i).Int(), 10)
			form.Add(alias, value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value := strconv.FormatUint(values.Field(i).Uint(), 10)
			form.Add(alias, value)
		case reflect.Float32, reflect.Float64:
			value := strconv.FormatFloat(values.Field(i).Float(), 'g', 10, 64)
			form.Add(alias, value)
		case reflect.String:
			form.Add(alias, values.Field(i).String())
		case reflect.Bool:
			value := strconv.FormatBool(values.Field(i).Bool())
			form.Add(alias, value)
		default:
			continue
		}
	}
	return form
}
