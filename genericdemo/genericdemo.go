package genericdemo

import (
	"fmt"
	"reflect"
	"sync"
)

func Add[T int | float64](x, y T) T {
	return x + y
}

func ReflectDemo() {
	str := "hello world"
	reflectType := reflect.TypeOf(str)
	println(reflectType)
}

func ChanDemo() {
	var wg sync.WaitGroup
	ch := make(chan int)
	go func() {
		for v := range ch {
			println(v)
		}
	}()

	for x := range 100 {
		wg.Add(1)
		go func() {
			ch <- x
			defer wg.Done()
		}()
	}

	wg.Wait()
	close(ch)
}
func ChanDemo2() {
	ch := make(chan int, 10)
	for x := range 10 {
		ch <- x
	}
	close(ch)
	for i := 0; i < 10; i++ {
		n, ok := <-ch
		fmt.Println(n, ok)
	}
}
