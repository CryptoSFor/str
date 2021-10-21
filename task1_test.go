package task1

import "testing"

func TestNumberOfCharacters(t *testing.T) {
	type str struct {
		s string
	}
	tests := []struct {
		name string
		args str
		want int
	}{
		{
			name: "string",
			args: str{s: "whoiam"},
			want: 6,
		},
		{
			name: "special characters",
			args: str{s: "!@#$%^*"},
			want: 7,
		},
		{
			name: "Chinese characters",
			args: str{s: "拷贝印相"},
			want: 4,
		},
		{
			name: "empty string",
			args: str{s: ""},
			want: 0,
		},
		{
			name: "string with spaces",
			args: str{s: "abc__ d  q "},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfCharacters(tt.args.s); got != tt.want {
				t.Errorf("NumberOfCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
