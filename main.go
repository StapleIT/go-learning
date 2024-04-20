package main

import (
	"fmt"
	"os/exec"

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
}
