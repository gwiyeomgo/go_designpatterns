package main

import (
	"fmt"
	"strconv"
	"strings"
)

// client
func main() {
	fontFactoryInstance := getSingleInstance()
	characterA := Character{
		value: "h",
		color: "white",
		font:  fontFactoryInstance.getFont("nanum:12"),
	}
	characterB := Character{
		value: "e",
		color: "white",
		font:  fontFactoryInstance.getFont("nanum:12"),
	}
	characterC := Character{
		value: "1",
		color: "red",
		font:  fontFactoryInstance.getFont("nanum:11"),
	}
	fmt.Println(characterA)
	fmt.Println(characterB)
	fmt.Println(characterC)
}

type Character struct {
	value string
	color string
	font  Font
}

// flytweight
type Font struct {
	family string
	size   int
}

// flytweight facotry
var (
	singleInstance = &FontFactory{
		fontMap: make(map[string]Font),
	}
)

type FontFactory struct {
	fontMap map[string]Font
}

// flyweight 을 가져오는 역할
func (f *FontFactory) getFont(font string) Font {
	data := f.fontMap[font]
	if data.family == "" {
		slice := strings.Split(font, ":")
		size, _ := strconv.Atoi(slice[1])
		f.fontMap[font] = Font{family: slice[0], size: size}
		fmt.Println("---- 새로운 객체 생성 " + font + "새 폰트 ----")
		return f.fontMap[font]
	}
	return data
}

func getSingleInstance() *FontFactory {
	return singleInstance
}
