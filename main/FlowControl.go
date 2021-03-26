package main

import (
	"fmt"
	"time"
)

func loop() {
	sum := 0
	// Go only have this one version of loop construct, no parenthesis required, but the braces are required always
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func loop2() {
	var sum int = 1
	// optional omit init and post statements
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)
}

func loop3() {
	var sum = 1;
	// "while" structure in Go
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func loop4() {
	// forever
	for{

	}
}

func ifStatement() {
	// Parenthesis of if statement could be omitted
	if true {
		fmt.Println("gg")
	}
	// if statement could begin with a init statement, which will be executed before the condition, in the scope of this "if"
	if temp := 3; temp < 4 {
		fmt.Println("aa")
	}

	if value := 4; value < 3 {
		fmt.Println("The variable declared in the if statement could be accessed by the else block")
	} else {
		fmt.Println(value)
	}
}

func switchStatement() {
	switch a := "gg"; a {
	case "hh": fmt.Println("NEU")
	case "gg": fmt.Println("MIT")
	default:
		fmt.Println("not matched")
	}
}

func switchStatementWithFunctionAsCondition()  {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithNoCondition()  {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

func main() {
	loop()
	loop2()
	loop3()
	//loop4()
	ifStatement()
	switchStatement()
	switchStatementWithFunctionAsCondition()
	switchWithNoCondition()
}
