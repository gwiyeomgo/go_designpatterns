package main

import "fmt"

var (
	shapeFactorySingleInstance = &ShapeFactory{
		shapeMap: make(map[string]Shape),
	}
)

type ShapeFactory struct {
	shapeMap map[string]Shape
}

func (s *ShapeFactory) getCircle(color string) Shape {
	circle := s.shapeMap[color]
	if circle == nil {
		s.shapeMap[color] = newCircle(color)
		fmt.Println("---- 새로운 객체 생성 " + color + "색 원 ----")
		return s.shapeMap[color]
	}
	return circle
}
func getShapeFactorySingleInstance() *ShapeFactory {
	return shapeFactorySingleInstance
}
