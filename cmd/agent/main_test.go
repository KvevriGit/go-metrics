package main

import (
	"reflect"
	"testing"
)

func TestSendTribute(t *testing.T) {
	type args struct {
		t string
		n string
		v string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SendTribute(tt.args.t, tt.args.n, tt.args.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendTribute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SendTribute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
