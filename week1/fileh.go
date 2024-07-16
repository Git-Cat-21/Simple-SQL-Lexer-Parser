/*package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("wow1.txt")
	if err != nil {
		fmt.Println("Error openinng file:: ", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	numBytes, err := file.Read(buffer)
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}
	filecontent := string(buffer[:numBytes])
	fmt.Println("file contents")
	fmt.Println(filecontent)
}*/

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	content := []byte("hello world!\n")
// 	file, err := os.Create("output.txt")
// 	if err != nil {
// 		fmt.Println("Error creating a file", err)
// 		return
// 	}
// 	defer file.Close()
// 	_, err = file.Write(content)
// 	if err != nil {
// 		fmt.Println("Error writing to file", err)
// 		return
// 	}
// 	fmt.Println("data written to file successfully!")

// }
