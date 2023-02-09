package main

import "fmt"

type Stack []PostfixExpression

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) push(data PostfixExpression) {
	*s = append(*s, data)
	fmt.Printf("%d pushed to stack\n", data)
}
func (s *Stack) pop() PostfixExpression {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return nil
	} else {
		top := len(*s) - 1
		data := (*s)[top]
		*s = (*s)[:top]
		return data
	}
}
