package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
NOTE: No entiendo la razón de la existencia de este paquete porque
al final mutex hace lo mismo y tienes que invocar mucho código para
hacer un cambio, los canales y los mutex pueden combinarse para
resolver el problema.
*/

type atomCounter struct {
	val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	X := 100
	Y := 4
	var waitGroup sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func(no int) {
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}

	waitGroup.Wait()
	fmt.Println(counter.Value())
}
