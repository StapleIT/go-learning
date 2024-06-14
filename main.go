package main

import (
	"fmt"
	"os/exec"

	ft "github.com/StapleIT/go-learning/function_types"
	h "github.com/StapleIT/go-learning/httpbymatryer"
	p "github.com/StapleIT/go-learning/pointers"
	u "github.com/StapleIT/go-learning/utils"
)

func main() {
	u.UtilFunc()

	fmt.Println("Hello")
	a := 0
	// note 'p.' below just means the function (or type in the case of type 'S') is from the 'pointers' package which was nemed 'p' in the import above
	p.Incr(&a) // pointer to the variable as attribute of function
	fmt.Println(a)
	var s = p.S{Name: "Odie", Age: 42}
	p.ChangeAge(&s, 23) // must explicitly include the pointer to 's' (&s)
	fmt.Println(s)      // {"Odie" 23}

	cmd := exec.Command("ls")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// this is exercising the principle of using a pointer as reveiver in a method definition
	// thus ensuring that the variable of the method's target type is not only
	// modified within the method but outside of it too
	// I'm not sure why I don't need to say 'p.IncrMthd()' for a method coming from the pointers (p) package ??
	// I think it's because the variable 'b' has already been created as type p.T so the methods of p.T are
	// automatically available without referencing the package ??

	var b p.T = 0
	for i := 0; i < 5; i++ {
		b.IncrMthd()
		//fmt.Println(i)
		//fmt.Println(b)
	}
	fmt.Println(b)

	//running code with funtion type definition from function_types

	// ft.MyFunc is a type which merely states that a variable declared of this type must
	// be a function with no arguments and no return values: func()
	var d ft.MyFunc = func() {
		fmt.Println("function is printing to stdout")
	}
	//interestingly, one can now call the variable as if it is a function itself
	d()
	d.RequiredMethod()
	var e func() = func() {
		fmt.Println("function of type func() with no method")
	}
	//type conversion to include the method from d witht the function declared as e:

	f := ft.MyFunc(e)
	f()
	f.RequiredMethod()

	//run the http server test
	h.Run()
}
