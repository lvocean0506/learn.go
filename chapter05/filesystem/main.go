package main

import "fmt"

func main() {
	var data string
	{
		var equipment IOInterface = &soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Paper{}
		data = equipment.Read()
		fmt.Println(data)
	}
	/*equipment := "纸带"
	if equipment == "纸带" {
		data = readFromPaper()
	} else if equipment == "磁带" {
		data = readFromMag()
	} else if equipment == "1.4软盘" {
		data = readFrom14Soft()
	} //todo else if ...else if ...*/
}

type IOInterface interface {
	Read() string
}

type soft struct{}

func (soft) Read() string {
	return "从软盘读******00000"
}

type Mag struct{}

func (Mag) Read() string {
	return "从磁带读******00000"
}

type Paper struct{}

func (Paper) Read() string {
	return "从纸带读******00000"
}
