package main

import "fmt"

// expression 1
type MinusExpression struct {
	left  PostfixExpression
	right PostfixExpression
}

func (v MinusExpression) interpreter(context map[string]int) int {
	re := v.left.interpreter(context) - v.right.interpreter(context)
	fmt.Println(fmt.Sprintf("%d ,MinusExpression", re))
	return re
	//return v.left.interpreter(context) - v.right.interpreter(context)
}

// expression 2
type PlusExpression struct {
	left  PostfixExpression
	right PostfixExpression
}

func (v PlusExpression) interpreter(context map[string]int) int {
	//return v.left.interpreter(context) + v.right.interpreter(context)
	re := v.left.interpreter(context) + v.right.interpreter(context)
	fmt.Println(fmt.Sprintf("%d ,PlusExpression", re))
	return re
}

// expression 3
type VariableExpression struct {
	variable string
}

func (v VariableExpression) interpreter(context map[string]int) int {
	return context[v.variable]
}

// xpression 4
type PostfixExpression interface {
	interpreter(context map[string]int) int
}

// parser
type PostfixParser struct {
}

func getExpression(c string, stack *Stack) PostfixExpression {
	switch c {
	case "+":
		return PlusExpression{stack.pop(), stack.pop()}
	case "-":
		right := stack.pop()
		left := stack.pop()
		return MinusExpression{left, right}
	default:
		return VariableExpression{c}
	}
}
func (p PostfixParser) parse(expression string) PostfixExpression {
	stack := Stack{}
	for _, ch := range expression {
		stack.push(getExpression(fmt.Sprintf("%c", ch), &stack))
	}
	return stack.pop()
}

func main() {
	postfixParser := PostfixParser{}
	//expression := postfixParser.parse("xyz+-")
	expression := postfixParser.parse("xyz+-a+")
	context := make(map[string]int, 0)
	context["x"] = 1
	context["y"] = 2
	context["z"] = 3
	context["a"] = 4
	result := expression.interpreter(context)
	fmt.Println(result)
}
