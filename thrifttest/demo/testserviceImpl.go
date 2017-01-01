package demo

import "fmt"

type TestImpl struct {
}

func NewTest() *TestImpl {
	return &TestImpl{}
}
func (t *TestImpl) HelloWorld() (r string, err error) {
	fmt.Println("return:hello,world")
	return "hello,world", nil
}

// Parameters:
//  - Name
func (t *TestImpl) HelloWorldForString(name string) (r string, err error) {
	fmt.Println("return:hello,world-" + name)
	//time.Sleep(6 * time.Second)
	return "hello,world-" + name, nil
}

// Parameters:
//  - Name
func (t *TestImpl) HelloWorldForMap(name map[string]int32) (r string, err error) {
	fmt.Printf("return:hello,world-%d\n", name["test"])
	return fmt.Sprintf("hello,world-%d\n", name["test"]), nil
}

// Parameters:
//  - Name
func (t *TestImpl) HelloWorldForStruct(name *Name) (r string, err error) {
	fmt.Println("return:hello,world-" + name.GetName())
	return "hello,world-" + name.GetName(), nil
}
