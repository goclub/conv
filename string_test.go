package xconv

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStringInt(t *testing.T) {
	i , err := StringInt("123")
	assert.Equal(t, int(123), i)
	assert.Equal(t, nil, err)
}
func TestStringInt64(t *testing.T) {
	{
		i64 , err := StringInt64("123")
		assert.Equal(t, int64(123), i64)
		assert.Equal(t, nil, err)
	}
	{
		{
			i64 , err := StringInt64("-123")
			assert.Equal(t, int64(-123), i64)
			assert.Equal(t, nil, err)
		}
	}
}


func TestStringFloat64(t *testing.T) {
	i , err := StringFloat64("123.1")
	assert.Equal(t, float64(123.1), i)
	assert.Equal(t, nil, err)
}
func TestStringFloat32(t *testing.T) {
	i , err := StringFloat32("123.1")
	assert.Equal(t, float32(123.1), i)
	assert.Equal(t, nil, err)
}
func TestStringBool(t *testing.T) {
	{
		sList := []string{"true", "True","1", "t", "T"}
		for i:=0 ; i< len(sList) ; i++ {
			b, err := StringBool(sList[0])
			assert.Equal(t, b, true)
			assert.Equal(t, err, nil)
		}
	}
	{
		sList := []string{"false", "False","f", "F", "0"}
		for i:=0 ; i< len(sList) ; i++ {
			b, err := StringBool(sList[0])
			assert.Equal(t, b, false)
			assert.Equal(t, err, nil)
		}
	}
	{
		sList := []string{"asd", "","3t", "2e3f", "sd"}
		for i:=0 ; i< len(sList) ; i++ {
			b, err := StringBool(sList[0])
			assert.Equal(t, b, false)
			assert.EqualError(t, err, "goclub/conv: " + sList[0]  + " can't conv to bool")
		}
	}
}
func TestStringToReflect(t *testing.T) {
	type Data struct {
		s string
		i int
		i8 int8
		i16 int16
		i32 int32
		i64 int64
		ui uint
		ui8 uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		bool bool
		f32 float32
		f64 float64
		bytes []byte
	}
	data := Data{}
	assert.NoError(t,
		StringToReflect("s", reflect.ValueOf(&data.s)),
	)
	assert.NoError(t,
		StringToReflect("-1", reflect.ValueOf(&data.i)),
	)
	assert.NoError(t,
		StringToReflect("-2", reflect.ValueOf(&data.i8)),
	)
	assert.NoError(t,
		StringToReflect("-3", reflect.ValueOf(&data.i16)),
	)
	assert.NoError(t,
		StringToReflect("-4", reflect.ValueOf(&data.i32)),
	)
	assert.NoError(t,
		StringToReflect("-5", reflect.ValueOf(&data.i64)),
	)
	assert.NoError(t,
		StringToReflect("1", reflect.ValueOf(&data.ui)),
	)
	assert.NoError(t,
		StringToReflect("2", reflect.ValueOf(&data.ui8)),
	)
	assert.NoError(t,
		StringToReflect("3", reflect.ValueOf(&data.ui16)),
	)
	assert.NoError(t,
		StringToReflect("4", reflect.ValueOf(&data.ui32)),
	)
	assert.NoError(t,
		StringToReflect("5", reflect.ValueOf(&data.ui64)),
	)
	assert.NoError(t,
		StringToReflect("true", reflect.ValueOf(&data.bool)),
	)
	assert.NoError(t,
		StringToReflect("0.1", reflect.ValueOf(&data.f32)),
	)
	assert.NoError(t,
		StringToReflect("0.2", reflect.ValueOf(&data.f64)),
	)
	assert.NoError(t,
		StringToReflect("b我", reflect.ValueOf(&data.bytes)),
	)
	assert.Equal(t,data, Data{"s", -1, -2, -3, -4, -5, 1, 2, 3, 4, 5, true, 0.1, 0.2, []byte("b我")})
	{
		{
			data := struct {
				Name string
			}{}
			assert.Errorf(t, StringToReflect("nimoc", reflect.ValueOf(data).Field(0)), "goclub/conv: StringToReflect(s, rValue) rValue must can set, mu be you should use reflect.ValueOf(pointer)")
		}
		{
			data := struct {
				Name string
			}{}
			assert.Errorf(t, StringToReflect("nimoc", reflect.ValueOf(data.Name)), "goclub/conv: StringToReflect(s, rValue) rValue must can set, mu be you should use reflect.ValueOf(pointer)")
		}
		{
			data := struct {
				Name string
			}{}
			assert.NoError(t,
				StringToReflect("nimoc", reflect.ValueOf(&data).Elem().Field(0)),
			)
			assert.Equal(t,data.Name, "nimoc")
		}
		{
			data := struct {
				Name string
			}{}
			assert.NoError(t,
				StringToReflect("nimoc", reflect.ValueOf(&data.Name)),
			)
			assert.Equal(t,data.Name, "nimoc")
		}
	}
}

func TestReflectToString(t *testing.T) {
	type Data struct {
		s string
		i int
		i8 int8
		i16 int16
		i32 int32
		i64 int64
		ui uint
		ui8 uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		bool bool
		f32 float32
		f64 float64
		bytes []byte
	}
	data := Data{"s", -1, -2, -3, -4, -5, 1, 2, 3, 4, 5, true, 0.1, 0.2, []byte("b我")}
	testdata := []string{"s", "-1", "-2", "-3", "-4","-5", "1", "2", "3", "4", "5", "1", "0.1", "0.2", "b我"}
	for i, v := range testdata {
		item := reflect.ValueOf(data).Field(i)
		value, err := ReflectToString(item)
		assert.NoError(t, err)
		assert.Equal(t, value, v)
	}
}