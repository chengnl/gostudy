package tclient

import (
	"fmt"
	"gostudy/thrifttest/demo"
	"sync"
	"testing"
)

func TestDemoClient(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			demoClient := DefaultServiceFactory().DemoService()
			client := demoClient.GetClient().(*demo.TestServiceClient)
			result, err := client.HelloWorld()
			if err != nil {
				demoClient.HandlerError(err)
			} else {
				fmt.Println(result)
			}
			wg.Done()

		}()
	}
	wg.Wait()

}
