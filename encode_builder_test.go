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
			name: "should write a horizontal tab byte to a string",
			args: args{
				stringAsByte: 9,
			},
			want: "\\t",
		},
		{
			name: "should write an escape byte to a string",
			args: args{
				stringAsByte: 27,
			},
			want: "\\u001b",
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

var _StopCompilerOptimizationBool bool

func Benchmark_isABlackListedControlAcceptable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := isABlackListedControl(64)
		_StopCompilerOptimizationBool = x
	}
}

func Benchmark_isABlackListedControlBell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := isABlackListedControl(7)
		_StopCompilerOptimizationBool = x
	}
}

func Benchmark_isABlackListedControlVertical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := isABlackListedControl(11)
		_StopCompilerOptimizationBool = x
	}
}

func Benchmark_isABlackListedControlEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := isABlackListedControl(27)
		_StopCompilerOptimizationBool = x
	}
}

func Benchmark_isABlackListedControlShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := isABlackListedControl(15)
		_StopCompilerOptimizationBool = x
	}
}
