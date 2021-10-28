package main

import "testing"

func Test_isOdd(t *testing.T) {
	type args struct {
		number int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Teste número par", args: args{number: 4}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.number); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
