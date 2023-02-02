package polymorphy

import "fmt"

type Light interface {
	shining()
}

type blue string
type green string
type red string

func (b blue) shining() {
	fmt.Println("blue is shining")
}
func (b green) shining() {
	fmt.Println("green is shining")
}
func (b red) shining() {
	fmt.Println("red is shining")
}

var lights = make(map[string]Light)

func init() {
	lights = map[string]Light{
		"blue":  blue(""),
		"green": green(""),
		"red":   red(""),
	}
}

func GetLighter(id string) Light {
	if l, ok := lights[id]; ok {
		return l
	}
	return nil
}
