package main

import (
	"fmt"
	"log"
	"os"
)

const meow int32 = 1

func main() {

	log.Println("hello there..", os.Args[1], []string{"hello hello2","hello3"})
	fmt.Println("hello there..", os.Args[1], []string{"hello hello2","hello3"})

	meow3 := "1"
	
	log.Println(meow3)
	meow3 = "33"
	log.Println(meow3)


	meowSlice := []string{"1","2", "3"}
	meowSlice[0] = "4"
	meowSlice = append(meowSlice, "append")

	newMeow := []string{"44"}
	newMeow = append( meowSlice, newMeow...)

	log.Println(newMeow, meowSlice)


	map1 := make(map[string]int)

	map1["index1"] = 1
	map1["index2"] = 1
	map1["index3"] = 1

	nestedMap := make(map[string]map[string]string)
	nestedMap["Iceland"] = make(map[string]string)
	nestedMap["Iceland"]["airport"] = "road1"


	log.Println(nestedMap)
	log.Println(nestedMap["Iceland"]["airport"])

	human1 := new(Account)
	human1.Age = 20
	human1.Height = 1.80
	human1.Name = "Programmer"

	human2 := new(Guest)
	human2.Name = "Programmer"
	
	gender := new(Config)
	gender.title = "binary"

	human1.Gender = gender
	log.Println(human1, human1.Gender)


	SecurityCheck(human1)
	SecurityCheck(human2)
}

// SecurityCheck ...
func SecurityCheck(thing Actor){
	thing.DoSomething()
}

// Actor ...
type Actor interface {
	DoSomething()
}
// Account ...
type Account struct {
	Age int
	Height float64
	Name string
	Gender *Config
}

// DoSomething ..
func (a *Account)DoSomething(){
	log.Println("Accounts can view restricted materials..")
}

// Guest ...
type Guest struct {
	Name string
}

// DoSomething ..
func (g *Guest)DoSomething(){
	log.Println("NO BUENO")
}
// Config ..
type Config struct {
	title string
}
