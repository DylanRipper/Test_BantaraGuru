package main

import "fmt"

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(char string) {
	*s = append(*s, char)
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var parenthesis Stack
	var x = "[()]{}{[()()]()}"
mainLoop:
	for _, value := range x {
		var char = string(value)
		switch char {
		case "[", "(", "{":
			parenthesis.Push(char)
			break
		case "]", "}", ")":
			if char == "]" && parenthesis.Peek() == "[" || char == "}" && parenthesis.Peek() == "{" || char == ")" && parenthesis.Peek() == "(" {
				parenthesis.Pop()
			} else {
				fmt.Println("False")
				break mainLoop
			}
		default:
			{
				fmt.Println("invalid char of parenthesis")
			}
		}
		if parenthesis.isEmpty() {
			fmt.Println("True")
		}
	}
}
