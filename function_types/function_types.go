package function_types

import "fmt"

type MyFunc func()

func (f MyFunc) RequiredMethod() {
	fmt.Println("executing required method")
}
