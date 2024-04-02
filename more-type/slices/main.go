package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]
	t := primes[:3]
	u := primes[0:3]
	v := primes[5:]
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(v)
}
