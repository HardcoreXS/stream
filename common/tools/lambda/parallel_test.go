package lambda

import (
	"fmt"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	type args struct {
		lambdas []Lambda
	}
	tt := struct {
		name string
		args args
	}{"test", args{[]Lambda{
		func() {
			fmt.Println("ok, print", "hello")
		},
		func() {
			fmt.Println("ok, cacl 1+2+3=", 1+2+3)
		},
		func() {
			fmt.Println("panic")
			x := 0
			_ = 1 / x
		},
	}}}
	t.Run(tt.name, func(t *testing.T) {
		Parallel(tt.args.lambdas...)
		t.Logf("Parallel run finish")
	})
	time.Sleep(time.Second)
}
