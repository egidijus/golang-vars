package main

import (
	"fmt"
	"os"
	// "reflect" // used for typing
	"runtime"
	"strings"
)

// func SearchForStringIndex(str string, substr string) int {
//     m := search.New(language.English, search.IgnoreCase)
//     start, _ := m.IndexString(str, substr)
//     if start == -1 {
//         return start, false
//     }
//     return start, true
// }

// index, found := SearchForStringIndex(os.LookupEnv("PROMPT"), 'E');

// if found {
//     fmt.Println("match starts at", index);
// }

func main() {

	if runtime.GOOS == "windows" {
		// fmt.Println(os.Getenv("PROMPT"))
		// fmt.Println(strings.Contains(os.Getenv("PROMPT"), "E"))
		// if strings.Contains(os.Getenv("PROMPT"), "$E") == true {
		// 	// 	// fmt.Println("Hello from Windows")
		// 	// 	// os.Unsetenv("PROMPT")
		// 	fmt.Println(os.Getenv("PROMPT"))
		// 	os.Unsetenv("PROMPT")
		// 	fmt.Println(os.Getenv("PROMPT"))
		// }
	}
	for _, pair := range os.Environ() {
		os.Unsetenv("CommonProgramFiles(x86)")
		// fmt.Println(reflect.TypeOf(pair))
		if strings.Contains(strings.Split(pair, "=")[1], "(x86") == true {
			os.Unsetenv(strings.Split(pair, "=")[0])
			fmt.Println("THESE ARE SMASHED ::" + strings.Split(pair, "=")[0])
			fmt.Println(pair)
		} else {
			// fmt.Println(pair)
		}
	}
}
