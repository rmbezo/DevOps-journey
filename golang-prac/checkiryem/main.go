package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type tokenType int

const (
	tNumber tokenType = iota
	tOp
	tLParen
	tRParen
)

type token struct {
	typ tokenType
	num float64
	op  string
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение (пустая строка — выход): ")
		line, err := in.ReadString('\n')
		if err != nil && len(line) == 0 {
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			return
		}

		val, err := evalExpr(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Println(val)
	}
}

func evalExpr(s string) (float64, error) {
	toks, err := tokenize(s)
	if err != nil {
		return 0, err
	}
	rpn, err := toRPN(toks)
	if err != nil {
		return 0, err
	}
	return evalRPN(rpn)
}

func tokenize(s string) ([]token, error) {
	var toks []token

	// prev: нужен для определения унарного минуса
	prevType := tOp // как будто перед началом был оператор (значит '-' в начале будет унарным)

	i := 0
	for i < len(s) {
		r := rune(s[i])

		if unicode.IsSpace(r) {
			i++
			continue
		}

		// число: [0-9]... или .[0-9]...
		if unicode.IsDigit(r) || r == '.' {
			start := i
			dotSeen := false
			for i < len(s) {
				c := rune(s[i])
				if unicode.IsDigit(c) {
					i++
					continue
				}
				if c == '.' {
					if dotSeen {
						return nil, fmt.Errorf("неверное число возле: %q", s[start:i+1])
					}
					dotSeen = true
					i++
					continue
				}
				break
			}
			numStr := s[start:i]
			if numStr == "." {
				return nil, errors.New("одна точка — не число")
			}
			v, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return nil, fmt.Errorf("не удалось прочитать число %q: %w", numStr, err)
			}
			toks = append(toks, token{typ: tNumber, num: v})
			prevType = tNumber
			continue
		}

		switch s[i] {
		case '(':
			toks = append(toks, token{typ: tLParen})
			prevType = tLParen
			i++
		case ')':
			toks = append(toks, token{typ: tRParen})
			prevType = tRParen
			i++
		case '+', '-', '*', '/':
			op := string(s[i])

			// унарные +/-
			if (op == "+" || op == "-") && (prevType == tOp || prevType == tLParen) {
				if op == "+" {
					// унарный плюс можно просто игнорировать
					i++
					continue
				}
				op = "u-"
			}

			toks = append(toks, token{typ: tOp, op: op})
			prevType = tOp
			i++
		default:
			return nil, fmt.Errorf("неизвестный символ: %q", string(s[i]))
		}
	}
	return toks, nil
}

func toRPN(toks []token) ([]token, error) {
	prec := map[string]int{
		"u-": 3,
		"*":  2,
		"/":  2,
		"+":  1,
		"-":  1,
	}
	rightAssoc := map[string]bool{
		"u-": true,
	}

	var out []token
	var st []token // стек операторов и скобок

	pop := func() token {
		t := st[len(st)-1]
		st = st[:len(st)-1]
		return t
	}

	for _, tk := range toks {
		switch tk.typ {
		case tNumber:
			out = append(out, tk)

		case tOp:
			for len(st) > 0 {
				top := st[len(st)-1]
				if top.typ != tOp {
					break
				}
				p1 := prec[tk.op]
				p2 := prec[top.op]

				// если лев.ассоц: p1 <= p2; если прав.ассоц: p1 < p2
				shouldPop := (!rightAssoc[tk.op] && p1 <= p2) || (rightAssoc[tk.op] && p1 < p2)
				if !shouldPop {
					break
				}
				out = append(out, pop())
			}
			st = append(st, tk)

		case tLParen:
			st = append(st, tk)

		case tRParen:
			found := false
			for len(st) > 0 {
				top := pop()
				if top.typ == tLParen {
					found = true
					break
				}
				out = append(out, top)
			}
			if !found {
				return nil, errors.New("несовпадающие скобки: лишняя ')'")
			}
		}
	}

	for len(st) > 0 {
		top := pop()
		if top.typ == tLParen || top.typ == tRParen {
			return nil, errors.New("несовпадающие скобки: лишняя '('")
		}
		out = append(out, top)
	}
	return out, nil
}

func evalRPN(rpn []token) (float64, error) {
	var st []float64

	pop := func() (float64, error) {
		if len(st) == 0 {
			return 0, errors.New("некорректное выражение")
		}
		v := st[len(st)-1]
		st = st[:len(st)-1]
		return v, nil
	}

	for _, tk := range rpn {
		if tk.typ == tNumber {
			st = append(st, tk.num)
			continue
		}
		if tk.typ != tOp {
			return 0, errors.New("некорректное выражение")
		}

		switch tk.op {
		case "u-":
			a, err := pop()
			if err != nil {
				return 0, err
			}
			st = append(st, -a)

		case "+", "-", "*", "/":
			b, err := pop()
			if err != nil {
				return 0, err
			}
			a, err := pop()
			if err != nil {
				return 0, err
			}

			var r float64
			switch tk.op {
			case "+":
				r = a + b
			case "-":
				r = a - b
			case "*":
				r = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("деление на ноль")
				}
				r = a / b
			}
			st = append(st, r)

		default:
			return 0, fmt.Errorf("неизвестный оператор: %q", tk.op)
		}
	}

	if len(st) != 1 {
		return 0, errors.New("некорректное выражение (лишние значения)")
	}
	return st[0], nil
}
