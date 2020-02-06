package main

import "fmt"

func sumfunc(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go sumfunc(s, c)
	//チャネルで合計した数を返り値でうけとる
	x := <-c
	fmt.Println(x)
}
