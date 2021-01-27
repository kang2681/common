package funcs

import "testing"

func TestMd5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			args: args{str: ""},
			want: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name: "123456",
			args: args{str: "123456"},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.str); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrCharlen(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			args: args{},
			want: 0,
		},
		{
			name: "数字",
			args: args{str: "1"},
			want: 1,
		},
		{
			name: "中文",
			args: args{str: "好"},
			want: 1,
		},

		{
			name: "英文",
			args: args{str: "a"},
			want: 1,
		},
		{
			name: "特殊字符",
			args: args{str: "@#$%^&*"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrCharlen(tt.args.str); got != tt.want {
				t.Errorf("StrCharlen() = %v, want %v", got, tt.want)
			}
		})
	}
}
