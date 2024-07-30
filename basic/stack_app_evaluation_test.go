package basic

import (
	"fmt"
	"slices"
	"testing"
)

func Map[A, B any, S ~[]A](s S, mapFunc func(a A) B) []B {
	l := len(s)
	if l == 0 {
		return []B{}
	}
	result := make([]B, l, l)
	for i, a := range s {
		b := mapFunc(a)
		result[i] = b
	}
	return result
}

// 这是一个对slice容器以左折叠方式进行缩减（reduce）的求值函数
// s是被折叠的slice，foldFunc是对两个值进行折叠得到一个值的折叠操作函数
func LeftFold[A any, S ~[]A](s S, foldFunc func(a1, a2 A) A) A {
	var result A //!!! 此时，result值是类型A的零值
	if len(s) == 0 {
		return result
	}
	for _, a := range s {
		result = foldFunc(result, a)
	}
	return result
}

type SymbolType int

const (
	INVALID SymbolType = iota
	IDENT
	OPERATOR
	LPAREN
	RPAREN
)

type Symbol struct {
	sblType SymbolType
	literal string
}

type Precedence int

const (
	_      int = 0
	LOWEST     = iota
	ADD
	MUL
	PAREN
)

var Operators = []string{"+", "-", "*", "/", "(", ")"}
var OprtPrecd = map[string]int{
	"+": ADD,
	"-": ADD,
	"*": MUL,
	"/": MUL,
	"(": PAREN,
	")": PAREN,
}

func isSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}
func isOPerator(literal string) bool {
	return slices.Contains(Operators, literal)
}
func isPrior(smb1, smb2 Symbol) bool {
	p1 := OprtPrecd[smb1.literal]
	p2 := OprtPrecd[smb2.literal]
	return p1 >= p2
}
func charToSymble(ch byte) Symbol {
	if isSpace(ch) {
		return Symbol{}
	}
	var literal = string(ch)
	if literal == "(" {
		return Symbol{sblType: LPAREN, literal: literal}
	} else if literal == ")" {
		return Symbol{sblType: RPAREN, literal: literal}
	} else if isOPerator(literal) {
		return Symbol{sblType: OPERATOR, literal: literal}
	} else if ch >= 'a' && ch < 'z' || ch >= 'A' && ch < 'Z' {
		return Symbol{sblType: IDENT, literal: literal}
	} else {
		panic("表达式中出现了无法解析的字符")
	}

}

func InfixExpToPostfixExp(infixExp string) (postFixExp string) {
	postFixExp = ""
	postFixSymbols := []Symbol{}
	var stack Stack[Symbol] = &SliceStackAny[Symbol]{}
	var readIdent Symbol
	var readOpt Symbol
	for i := 0; i < len(infixExp); i++ {
		ch := infixExp[i]
		if isSpace(ch) {
			continue
		}
		curSmb := charToSymble(ch)
		//TODO 判断最后一个操作符的合法
		switch curSmb.sblType {
		case IDENT:
			readIdent = curSmb

			if i == len(infixExp)-1 {
				postFixSymbols = append(postFixSymbols, readIdent)
				readOpt = stack.Pop() //弹出上一个操作符
				if stack.IsEmpty() {
					postFixSymbols = append(postFixSymbols, readOpt)

				} else {
					priorOPt := stack.Pop()
					if isPrior(priorOPt, readOpt) {
						postFixSymbols = append(postFixSymbols, priorOPt)
						postFixSymbols = append(postFixSymbols, readOpt)
					} else {
						postFixSymbols = append(postFixSymbols, readOpt)
						postFixSymbols = append(postFixSymbols, priorOPt)
					}
				}
			}
		case LPAREN:
			stack.Push(curSmb) //
		case RPAREN:
			postFixSymbols = append(postFixSymbols, readIdent)
			readIdent = Symbol{}
			for !stack.IsEmpty() && stack.Top().sblType != LPAREN {
				priorOpt := stack.Pop() //弹出前面的操作符，因为它的右操作数已经找到
				postFixSymbols = append(postFixSymbols, priorOpt)
			}
			if stack.IsEmpty() {
				panic("错误，左右括号不匹配")
			}
			stack.Push(curSmb)
		case OPERATOR:
			readOpt = curSmb
			//对于读取到的操作符，其左操作数要么是readIdent，要么是栈顶操作符与readIdent组成的表达式。
			//如果堆栈为空，且readIdent的类型也不是标识符（IDENT），那当前的中缀操作符就没有左操作数了。
			if readIdent.sblType == INVALID {
				if stack.IsEmpty() {
					panic("错误，操作符缺失左操作数")
				} else if stack.Top().sblType != RPAREN {
					panic("错误，出现了相邻的两个中缀操作符")
				} else {
					stack.Pop()
					stack.Pop()
				}
			}

			//当前栈中没有上一个操作符，表明已读出的操作数是已读出的操作符的左操作数
			if stack.IsEmpty() && readIdent.sblType != INVALID {
				left := readIdent
				postFixSymbols = append(postFixSymbols, left)
				stack.Push(readOpt) //已读取的操作数作为左操作数压栈
				readIdent = Symbol{}
				readIdent = Symbol{}
				continue //当前标识符已经处理完毕
			}
			priorOpt := stack.Top()
			//上一个操作符的优先级低于当前读取的操作符，意味着当前读取的操作数是下一个操作符的左操作数，
			//当前操作符及其连接的表达式是上一个操作符的右操作数。
			if priorOpt.sblType == LPAREN || !isPrior(priorOpt, readOpt) { //当前操作符优先级低于前一个操作
				left := readIdent
				postFixSymbols = append(postFixSymbols, left)
				stack.Push(readOpt)  //把当前的操作符作为压栈，等待右操作数的形成
				readIdent = Symbol{} //设置为零值
				readOpt = Symbol{}   //设置为零值
			} else { // 上一个操作符的优先级高于或等于所读出来的操作符，则表明前一个操作符的右操作符是rendIdent
				right := readIdent
				postFixSymbols = append(postFixSymbols, right)
				postFixSymbols = append(postFixSymbols, priorOpt)
				stack.Pop() //弹出上一个操作符
				for !stack.IsEmpty() {
					priorOpt := stack.Top()
					if isPrior(priorOpt, readOpt) {
						postFixSymbols = append(postFixSymbols, priorOpt)
						stack.Pop()
					} else {
						break
					}
				}
				stack.Push(readOpt)
				readIdent = Symbol{}
				readOpt = Symbol{}
			}
		}
	}

	postFixExp = LeftFold(Map(postFixSymbols,
		func(smbl Symbol) string { return smbl.literal }),
		func(s1, s2 string) string { return s1 + s2 })
	return postFixExp
}

func TestInfixExpToPostfixExp(t *testing.T) {
	exp := "a+(b-c)*d/(e+g)*h-d"
	fmt.Println(exp)
	result := InfixExpToPostfixExp(exp)
	fmt.Println(result)
}
