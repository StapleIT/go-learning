package utilities

import (
	"fmt"
	"log"
	"os"
)

func UtilFunc() {
	fmt.Println("Test of utils ... this function also writes to a log.txt file")
	// O_APPEND = Append data to the file when writing.
	// O_CREATE = Create a new file if none exists.
	// O_WRONLY = Open the file write-only.
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("log.txt", flags, 0666)
	if err != nil {
		fmt.Printf("Couldn't open the log file.  Error: %v", err)
	}
	logger := log.New(file, "Log test prefix: ", log.Ldate|log.Ltime)
	logger.Println(log.Ltime)

}
func letter2Number(ltr string) int {
	//intString, _ := strconv.Atoi(ltr)
	runes := []rune(ltr)
	intString := int(runes[0])
	fmt.Printf("letter: %v, string as ascii: %v of type %T\n", ltr, intString, intString)
	out := intString - 64
	// if ltr == rune("H") {
	// 	return 8
	// }
	return out
}
