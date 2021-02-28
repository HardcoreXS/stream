package lambda

import (
	"runtime"
)

type Lambda func()

func Parallel(lambdas ...Lambda) {
	doneChan := make(chan bool, len(lambdas))
	for i := range lambdas {
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					_, _, _, _ = runtime.Caller(3)
					//caller := fmt.Sprintf("%s:%d", path.Base(file), line)
				}
				doneChan <- true
			}()
			lambdas[i]()
		}(i)
	}
	for range lambdas {
		<-doneChan
	}
}
