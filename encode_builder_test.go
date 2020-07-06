package gojay

import (
	"bytes"
	"testing"
)

// https://www.ascii-code.com/
func Test_writeStringEscape(t *testing.T) {
	type args struct {
		stringAsByte int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should not write a null byte to a string",
			args: args{
				stringAsByte: 0,
			},
			want: "",
		},
		{
			name: "should not write a bell byte to a string",
			args: args{
				stringAsByte: 7,
			},
			want: "",
		},
		{
			name: "should not write a vertical tab byte to a string",
			args: args{
				stringAsByte: 11,
			},
			want: "",
		},
		{
			name: "should not write a shift in byte to a string",
			args: args{
				stringAsByte: 15,
			},
			want: "",
		},
		{
			name: "should not write an escape byte to a string",
			args: args{
				stringAsByte: 27,
			},
			want: "",
		},
		{
			name: "should write a horizontal tab byte to a string",
			args: args{
				stringAsByte: 9,
			},
			want: "\\t",
    },
		{
			name: "should write an at byte to a string",
			args: args{
				stringAsByte: 64,
			},
			want: "@",
		},
	}
  
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			e := NewEncoder(b)
			e.writeStringEscape(string(rune(tt.args.stringAsByte)))
			got := string(e.buf)
			if got != tt.want {
				t.Fatalf("writeStringEscape() got %s want %s", got, tt.want)
			}
		})
	}
}
