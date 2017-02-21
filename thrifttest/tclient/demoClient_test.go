package tclient

import (
	"fmt"
	"sync"
	"testing"
)

func TestDemoClient(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			result, err := DefaultServiceFactory().DemoService().Call("HelloWorld")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result.(string))
			}
			result, err = DefaultServiceFactory().DemoService().Call("HelloWorldForString", "wuyun")
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result.(string))
			}
			wg.Done()

		}()
	}
	wg.Wait()

}
