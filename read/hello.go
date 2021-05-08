package main

//go mod init read
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("../data.txt")
	check(err)
	fmt.Print(string(dat))
	/*https://gobyexample.com/reading-files********************/
	f, err := os.Open("../data.txt")
	check(err)

	b1 := make([]byte, 16)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("lbl16 b1 %d bytes: %s\n", n1, string(b1[n1-2:n1]))
	/*https://gobyexample.com/reading-files********************/
	o3, err := f.Seek(18, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2) /*no err*/
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	/*https://gobyexample.com/reading-files********************/

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(18)
	check(err)
	fmt.Printf("lbl r4 18 bytes: %s\n", string(b4))
}
