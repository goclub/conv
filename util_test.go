package xconv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Case struct {
	Input  interface{}
	Want   interface{}
	ErrStr string
}

// 用法
// func TestStringBool(t *testing.T) {
// 	Cases(t, StringBool, []Case{
// 		{"1", true, ``},
// 		{"0", false, ``},
// 		{"a", false, `strconv.ParseBool: parsing "a": invalid syntax`},
// 	})
// }
func Cases(t *testing.T, fn interface{}, cases []Case) {
	for _, c := range cases {
		fnReflact := reflect.ValueOf(fn)
		code := fmt.Sprintf("%s(%v)", fnReflact.Type().Name(), c.Input)
		result := fnReflact.Call([]reflect.Value{reflect.ValueOf(c.Input)})
		switch len(result) {
		case 1:
		case 2:
		default:
			t.Errorf(code+" return to many: %d", len(result))
		}
		output := result[0].Interface()
		assert.Equal(t, output, c.Want, code)
		if len(result) == 2 {
			errAny := result[1].Interface()
			if errAny != nil {
				assert.EqualError(t, errAny.(error), c.ErrStr, code)
			}
		}
	}
}
