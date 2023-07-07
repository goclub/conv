package xconv

import (
	xerr "github.com/goclub/error"
	"reflect"
	"testing"
)

func TestStringInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "negative input",
			args: args{"-123"},
			want: -123,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
		{
			name:    "integer overflow",
			args:    args{"99999999999999999999999999999999999999999999999999"},
			wantErr: true,
		},
		{
			name:    "negative integer overflow",
			args:    args{"-99999999999999999999999999999999999999999999999999"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringInt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInt8(t *testing.T) {
	testCases := []struct {
		input       string
		expected    int8
		expectedErr error
	}{
		{"123", 123, nil},
		{"0", 0, nil},
		{"-123", -123, nil},
		{"127", 127, nil},
		{"-128", -128, nil},
		{"128", 0, xerr.New("strconv.ParseInt: parsing \"128\": value out of range")},
		{"-129", 0, xerr.New("strconv.ParseInt: parsing \"-129\": value out of range")},
		{"hello", 0, xerr.New("strconv.ParseInt: parsing \"hello\": invalid syntax")},
	}

	for _, tc := range testCases {
		actual, actualErr := StringInt8(tc.input)

		if actual != tc.expected {
			t.Errorf("StringInt8(%s) = %d; expected %d", tc.input, actual, tc.expected)
		}

		if (tc.expectedErr == nil && actualErr != nil) || (tc.expectedErr != nil && actualErr == nil) || (tc.expectedErr != nil && actualErr != nil && tc.expectedErr.Error() != actualErr.Error()) {
			t.Errorf("StringInt8(%s) error = %v; expected %v", tc.input, actualErr, tc.expectedErr)
		}
	}
}

func TestStringInt16(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "negative input",
			args: args{"-123"},
			want: -123,
		},
		{
			name: "max input",
			args: args{"32767"},
			want: 32767,
		},
		{
			name: "min input",
			args: args{"-32768"},
			want: -32768,
		},
		{
			name:    "overflow input",
			args:    args{"99999"},
			wantErr: true,
		},
		{
			name:    "underflow input",
			args:    args{"-99999"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringInt16(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInt32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "negative input",
			args: args{"-123"},
			want: -123,
		},
		{
			name: "max input",
			args: args{"2147483647"},
			want: 2147483647,
		},
		{
			name: "min input",
			args: args{"-2147483648"},
			want: -2147483648,
		},
		{
			name:    "overflow input",
			args:    args{"9999999999"},
			wantErr: true,
		},
		{
			name:    "underflow input",
			args:    args{"-9999999999"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringInt32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringUint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"18446744073709551615"},
			want: 18446744073709551615,
		},
		{
			name:    "overflow input",
			args:    args{"18446744073709551616"},
			wantErr: true,
		},
		{
			name:    "overflow input -1",
			args:    args{"-1"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUint(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringUint8(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"255"},
			want: 255,
		},
		{
			name:    "overflow input",
			args:    args{"256"},
			wantErr: true,
		},
		{
			name:    "overflow input -1",
			args:    args{"-1"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUint8(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringUint16(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"65535"},
			want: 65535,
		},
		{
			name:    "overflow input",
			args:    args{"65536"},
			wantErr: true,
		},
		{
			name:    "overflow input -1",
			args:    args{"-1"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUint16(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringUint32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"4294967295"},
			want: 4294967295,
		},
		{
			name:    "overflow input",
			args:    args{"4294967296"},
			wantErr: true,
		},
		{
			name:    "overflow input -1",
			args:    args{"-1"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUint32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringUint64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123"},
			want: 123,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"18446744073709551615"},
			want: 18446744073709551615,
		},
		{
			name:    "overflow input",
			args:    args{"18446744073709551616"},
			wantErr: true,
		},
		{
			name:    "overflow input -1",
			args:    args{"-1"},
			wantErr: true,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUint64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringFloat32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123.45"},
			want: 123.45,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"3.4028235e+38"},
			want: 3.4028235e+38,
		},
		{
			name: "underflow input",
			args: args{"1.17549435e-38"},
			want: 1.17549435e-38,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringFloat32(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringFloat64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{"123.45"},
			want: 123.45,
		},
		{
			name: "zero input",
			args: args{"0"},
			want: 0,
		},
		{
			name: "max input",
			args: args{"1.7976931348623157e+308"},
			want: 1.7976931348623157e+308,
		},
		{
			name: "underflow input",
			args: args{"4.9406564584124654e-324"},
			want: 4.9406564584124654e-324,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringFloat64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBool(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "valid input true",
			args: args{"true"},
			want: true,
		},
		{
			name: "valid input false",
			args: args{"false"},
			want: false,
		},
		{
			name: "valid input true",
			args: args{"1"},
			want: true,
		},
		{
			name: "valid input false",
			args: args{"0"},
			want: false,
		},
		{
			name: "valid input true",
			args: args{"t"},
			want: true,
		},
		{
			name: "valid input false",
			args: args{"f"},
			want: false,
		},
		{
			name:    "invalid input",
			args:    args{"hello"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringBool(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("StringBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "valid input",
			args: args{"hello world"},
			want: []byte("hello world"),
		},
		{
			name: "empty input",
			args: args{""},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
