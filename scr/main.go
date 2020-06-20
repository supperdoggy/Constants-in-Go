package main

import "fmt"

func main() {
	// to set the constant use word const
	// we dont name constants uppercase, we just
	// name them like a regular
	// variable
	const myConst int = 42                   // 42
	fmt.Printf("%v, %T\n", myConst, myConst) // 42, int
	/*
		You cant assign a function return to a constant
		const myConst float64 = math.Sin(1) will cause an error
	*/
	// Also you can use const to int, string, float, bool
	// but you can`t make an array constant
	// you can use constant like a variable but cant reassign it

	// also you can assign constants like this
	const a = 42     // 42 of integer
	var b int16 = 27 // 27 of int16
	// and you can mix different types of variables
	// with constants
	fmt.Printf("%v, %T\n", a+b, a+b) // 69, int 16

	// iota is a special counter
	// which changes i+=1 with every iota created

	// also you can create a block of constants
	const (
		c = iota // 0
		l = iota // 1
	)

	// aslo you can create an iota constants like this
	// and if you create a new set of iotas then it will restart the counter
	const (
		n = iota // 0
		v        // 1
		p        // 2
	)

	// we can very efficiency strode data of rules in a single byte of information
	const (
		isAdmin = 1 << iota
		isHQ
		canFire
	)
	var roles byte = isAdmin
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Can fire? %v\n", canFire&roles == canFire)
	// very good way to check permissions!!
}
