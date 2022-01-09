package main

import "fmt"

func main() {
	strArray := [5]string{"I", "am", "stupid", "and", "weak"}

	fmt.Printf("修改前:%s\n", strArray)

	for index, str := range strArray {
		switch str {
		case "stupid":
			strArray[index] = "smart"
		case "weak":
			strArray[index] = "strong"
		default:
		}
	}

	fmt.Printf("修改后:%s\n", strArray)

}
