package main

import (
	ch "essential/src/challenges"
	"fmt"
	"os"
)

func main() {
	ch.Fizbuzz()
	fmt.Println("Even ended count:")
	// fmt.Printf("Count is %v\n", evenendcount())
	inp_arr := []int{6, 7, 77, 3, -99}
	fmt.Println("Max value of slice is ", ch.FindMax((inp_arr)))

	ctype, err := ch.ContentType("https://www.linkedin.com")
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println(ctype)
	}

	some_text := "A quick brown dog jumps over a lazy dog"
	fmt.Printf("Word count is %v\n", ch.CountWords(some_text))

	tempsq, err := ch.NewSquare(-55, 6, 42)
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("The area of temp square is %v\n", tempsq.Area())
		tempsq.Move(92, 13)
		fmt.Printf("The new positons are %v and %v\n%#v\n", tempsq.X, tempsq.Y, tempsq)
	}

	c := &ch.Capper{Wrt: os.Stdout}
	fmt.Fprintln(c, "Hello There")

	fmt.Println(ch.Min([]float64{22, 333, 1}))
	fmt.Println(ch.Min([]string{"a", "z", "C"}))

}
