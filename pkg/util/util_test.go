package util

import (
	"reflect"
	"strings"
	"testing"
)

func TestDump(t *testing.T) {

}

func TestJSON(t *testing.T) {
	type Test struct {
		Name  string  `json:"name"`
		Age   int     `json:"age"`
		Money float64 `json:"money"`
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "测试正常输出",
			args: args{
				v: Test{
					Name:  "zhan",
					Age:   10,
					Money: 10.1,
				},
			},
			want:    []byte("{\"name\":\"zhan\",\"age\":10,\"money\":10.1}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSON(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONParseString(t *testing.T) {
	type Test struct {
		Name  string  `json:"name"`
		Age   int     `json:"age"`
		Money float64 `json:"money"`
	}
	type args struct {
		d string
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Test
	}{
		{
			name: "测试 json 解析",
			args: args{
				d: "{\"name\":\"zhan\",\"age\":10,\"money\":10.1}",
				v: &Test{},
			},
			wantErr: false,
			want: Test{
				Name:  "zhan",
				Age:   10,
				Money: 10.1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONParseString(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JSONParseString() error = %v, wantErr %v", err, tt.wantErr)
			} else if !reflect.DeepEqual(*tt.args.v.(*Test), tt.want) {
				t.Errorf("JSONParseString want: %v, got %v", tt.want, *tt.args.v.(*Test))
			}
		})
	}
}

func TestJSONString(t *testing.T) {
	type Test struct {
		Name  string  `json:"name"`
		Age   int     `json:"age"`
		Money float64 `json:"money"`
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "结构体转 json",
			args: args{
				v: Test{
					Name:  "zhan",
					Age:   10,
					Money: 10.1,
				},
			},
			want: "{\"name\":\"zhan\",\"age\":10,\"money\":10.1}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JSONString(tt.args.v); got != tt.want {
				t.Errorf("JSONString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandStringN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "获取随机字符串 len=10",
			args: args{
				n: 10,
			},
			want: 10,
		},
		{
			name: "获取随机字符串 len=0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "获取随机字符串 len=-1",
			args: args{
				n: -1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandStringN(tt.args.n); len(got) != tt.want {
				t.Errorf("RandStringN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "测试生成UUID",
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UUID(); len(got) != tt.want {
				t.Errorf("UUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlToBase64QrCode(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试生成二维码",
			args: args{
				url: "http://www.baidu.com",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UrlToBase64QrCode(tt.args.url)
			if !strings.Contains(got, "data:image/png;base64,") {
				t.Errorf("返回值中不包含 data:image/png;base64, 字符")
			}
		})
	}
}
