package basic

import (
	"errors"
	"fmt"
	"slices"
	"testing"
)

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
	ADD_LEVEL
	MUL_LEVEL
	PAREN_LEVEL
)

// Operators定义了操作符列表
var Operators = []string{"+", "-", "*", "/"}

// OprtPrecd 定义了操作符优先级字典
var OprtPrecd = map[string]int{
	"+": ADD_LEVEL,
	"-": ADD_LEVEL,
	"*": MUL_LEVEL,
	"/": MUL_LEVEL,
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

type ExprParser struct {
	originExpr     string
	postFixSymbols []Symbol
	stack          Stack[Symbol]
}

func NewExprParser(expr string) ExprParser {
	return ExprParser{
		originExpr:     expr,
		postFixSymbols: []Symbol{},
		stack:          &SliceStackAny[Symbol]{},
	}
}

func (ep ExprParser) GetPostFixExpr() string {
	postFixExp := LeftFold(Map(ep.postFixSymbols,
		func(smbl Symbol) string { return smbl.literal }),
		func(s1, s2 string) string { return s1 + s2 })
	return postFixExp
}

func (ep *ExprParser) Execute() {
	var lastSymbol Symbol //上一个处理过的符号，用来辅助检查表达式的符号之间的连接是否合理。
	for i := 0; i < len(ep.originExpr); i++ {
		ch := ep.originExpr[i]
		if isSpace(ch) {
			continue
		}
		curSmb := charToSymble(ch)
		switch curSmb.sblType {
		case IDENT: //!!!遇到标识符就压栈，有优先级的操作符才需要处理彼此的先后顺序。
			if lastSymbol.sblType == IDENT {
				panic("表达式错误,连续出现了两个标识符")
			}
			ep.postFixSymbols = append(ep.postFixSymbols, curSmb)
		case LPAREN: //!!!遇到左括号就压栈处理，作为“括号帧的帧底”，当遇到右括号时，经将该括号帧全部弹出
			if lastSymbol.sblType == IDENT {
				panic("表达式错误,左括号出现了标识符")
			}
			ep.stack.Push(curSmb) //
		case RPAREN: //!!! 遇到右括号就不停弹出栈内操作符，直到弹出括号帧的帧底——左括号为止。
			if lastSymbol.sblType == LPAREN || lastSymbol.sblType == OPERATOR {
				panic("表达式错误,右括号前直接出现了左括号或者操作符")
			}
			for !ep.stack.IsEmpty() && ep.stack.Top().sblType != LPAREN {
				priorOpt := ep.stack.Pop() //弹出前面的操作符，因为它的右操作数已经找到
				ep.postFixSymbols = append(ep.postFixSymbols, priorOpt)
			}
			if ep.stack.IsEmpty() {
				panic("表达式错误，左右括号不匹配")
			}
			ep.stack.Pop() //弹出对应的左括号——括号帧的帧底
		case OPERATOR: //!!!如果单前操作符的优先级小于或等于栈里的操作符，就应先把栈里的操作符弹出到后缀表达式列表中，以便优先计算,否则就把自己压栈，待由后续操作符的优先级来决定。
			if lastSymbol.sblType == INVALID || lastSymbol.sblType == LPAREN || lastSymbol.sblType == OPERATOR {
				panic("表达式错误，操作符出现在错误位置")
			}
			if ep.stack.IsEmpty() || ep.stack.Top().sblType == LPAREN || //考虑到左括号这种特殊的操作符可能会在栈中的情况
				!isPrior(ep.stack.Top(), curSmb) {
				ep.stack.Push(curSmb)
			} else {
				for !ep.stack.IsEmpty() && isPrior(ep.stack.Top(), curSmb) {
					priorOpt := ep.stack.Pop()
					ep.postFixSymbols = append(ep.postFixSymbols, priorOpt)
				}
				ep.stack.Push(curSmb)
			}
		}
		lastSymbol = curSmb //更新上一个处理过的符号，以便遍历下一个符号时使用
	} //循环遍历表达式字符串结束
	for !ep.stack.IsEmpty() { //考虑到表达式最后的符号可能是标识符而不是操作符，将剩余操作符弹出处理
		symbOpt := ep.stack.Pop()
		if symbOpt.sblType == LPAREN {
			panic("表达式错误，左右括号不匹配")
		}
		ep.postFixSymbols = append(ep.postFixSymbols, symbOpt)
	}
}

type Number interface {
	~uint | ~uint32 | ~uint64 | ~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Compute[N Number](operator string, left, right N) N {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		panic("无法识别的操作符")
	}

}
func Evalueate[N Number](values map[string]N, postFixSymbols []Symbol) N {
	valueStack := SliceStackAny[N]{}
	for _, symbl := range postFixSymbols {
		switch symbl.sblType {
		case IDENT:
			symbValue := values[symbl.literal]
			valueStack.Push(symbValue)
		case OPERATOR:
			rightValue := valueStack.Pop()
			leftValue := valueStack.Pop()
			tempResult := Compute(symbl.literal, leftValue, rightValue)
			valueStack.Push(tempResult)
		}
	}
	return valueStack.Pop()
}

func TestInfixExpToPostfixExp(t *testing.T) {
	exp := ("a + (b - c) / (d * e)")
	//exp := "a +(b-c)*d/(e+g)*h-d*d"
	fmt.Println(exp)
	var result string
	result = InfixExpToPostfixExp(exp)
	fmt.Println(result)
	ep := NewExprParser(exp)
	ep.Execute()
	result = ep.GetPostFixExpr()
	fmt.Println(result)
	vm := map[string]int{
		"a": 1,
		"b": 10,
		"c": 5,
		"d": 4,
		"e": 2,
		"g": 3,
		"h": 6,
	}
	resultValue := Evalueate(vm, ep.postFixSymbols)
	println(resultValue)
	values := make(map[string]float64)
	values["a"] = 10
	values["b"] = 5
	values["c"] = 2
	values["d"] = 4
	values["e"] = 3
	fmt.Println(Evalueate(values, ep.postFixSymbols))

}

func TestFeature(t *testing.T) {

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	combinedErr := errors.Join(err1, err2)
	fmt.Println(combinedErr)
}
