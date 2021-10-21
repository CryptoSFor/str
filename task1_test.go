package task1

import "testing"

func TestReverseString(t *testing.T) {
	type str struct {
		s string
	}
	tests := []struct {
		name string
		args str
		want string
	}{
		{
			name: "string",
			args: str{s: "whoiam"},
			want: "maiohw",
		},
		{
			name: "special characters",
			args: str{s: "!@#$%^&*()_+"},
			want: "+_)(*&^%$#@!",
		},
		{
			name: "Chinese characters",
			args: str{s: "拷贝印相"},
			want: "相印贝拷",
		},
		{
			name: "empty string",
			args: str{s: ""},
			want: "",
		},
		{
			name: "string with spaces",
			args: str{s: "abc__ d  q "},
			want: " q  d __cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
