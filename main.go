package main

import "fmt"

type bot interface {
	getGreeting() string //if there is any other type in our program that has a function getGreeting() that is also promoted to be type 'bot
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//print greeting has VERY similar implentation (use of interfaces would help a lot here)
/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//very custom logic for english greeting -> reciever
func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

//very custom logic for spanish greeting -> reciever
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
