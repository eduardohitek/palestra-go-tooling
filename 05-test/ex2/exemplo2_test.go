package main

import "testing"

func Test_notSoRandom(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "Teste 1", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notSoRandom(); got != tt.want {
				t.Errorf("notSoRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
