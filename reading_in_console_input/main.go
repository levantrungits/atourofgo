package main

import (
	"fmt"
	"bufio"
	"os"
	//"strings"
)

// The above code will infinitely ask scan for input and echo back whatever is entered.
func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main()  {
	
/*
	Reading in Full Sentences
*/
fmt.Println("==============================================================Reading in Full Sentences:")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Simple Shell")
	// fmt.Println("---------------------------")

	// for {
	// 	fmt.Print("-> ")
	// 	text, _ := reader.ReadString('\n')
	// 	// convert CRLF to LF
	// 	//	Windows
	// 	text = strings.Replace(text, "\r\n", "", -1)
	// 	//	Unix
	// 	//text = strings.Replace(text, "\n", "", -1)

	// 	if strings.Compare("hi", text) == 0 {
	// 		fmt.Println("hello, Yourself")
	// 	} else {
	// 		fmt.Println("hello, ", text)
	// 	}
	// }

/*
	Reading Single UTF-8 Encoded Unicode Characters
*/
fmt.Println("==============================================================Reading Single UTF-8 Encoded Unicode Characters:")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// Print out the unicode value i.e. A->65, a->97
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}

/*
	Using Bufio's Scanner
		A third way you could potentially read in input from the console in go is by creating a new scanner and passing os.
		Stdin just as we have done above creating new readers and then using scanner.
*/
fmt.Println("==============================================================Using Bufio's Scanner:")
	scanner()

}