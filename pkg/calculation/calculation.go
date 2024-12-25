package calculation

import (
	//"errors"
	"strconv"
	"strings"
	"unicode"
)

func SumStr(exp []string) string {
	res := ""
	for _, elem := range exp {
		res += elem
	}
	return res
}
func Split(expression string) ([]string, error) {
	exp := []rune(expression)
	var err error = nil
	var res []string
	var cur []rune
	for _, elem := range exp {
		if unicode.IsDigit(elem) {
			cur = append(cur, elem)
		} else {
			if len(cur) == 0 && elem != '(' && res[len(res)-1] != ")" {
				//err = errors.New("invalid operation")
				err = ErrInvalidExpression
			}
			res = append(res, string(cur))
			res = append(res, string(elem))
			cur = cur[0:0]
		}
	}
	if len(cur) == 0 {
		//err = errors.New("invalid operation")
		err = ErrInvalidExpression
	}
	res = append(res, string(cur))
	return res, err
}
func Math(exp []string) ([]string, error) {
	var tmp_e string
	var err error
	var tmp_m []string
	t := []string{""}
	exp = append(t, exp...)
	exp = append(exp, t...)
	for strings.Contains(SumStr(exp), "(") {
		start := 0
		end := 0
		for ind, elem := range exp {
			if elem == "(" {
				start = ind

			} else if elem == ")" {
				end = ind

				break
			}

		}
		tmp := exp[end+1:]

		tmp_m, err = Math(exp[start+1 : end])
		if err != nil {
			return nil, err
		}
		exp = append(exp[:start-1], tmp_m...)
		exp = append(exp, tmp[1:]...)

	}
	for strings.Contains(SumStr(exp), "*") || strings.Contains(SumStr(exp), "/") {
		start := 0
		end := 0
		oper := -1
		for ind, elem := range exp {
			if elem == "*" {
				start = ind - 1
				end = ind + 1
				oper = 1
				break

			} else if elem == "/" {
				start = ind - 1
				end = ind + 1
				oper = 2
				break
			}

		}
		tmp := exp[end+1:]
		if oper == 1 {
			tmp_e, err = Mult(exp[start : end+1])
			if err != nil {
				return nil, err
			}
			exp = append(exp[:start], tmp_e)
		} else if oper == 2 {
			tmp_e, err = Div(exp[start : end+1])
			if err != nil {
				return nil, err
			}
			exp = append(exp[:start], tmp_e)
		}

		exp = append(exp, tmp...)

	}
	for strings.Contains(SumStr(exp), "+") || strings.Contains(SumStr(exp), "-") {

		start := 0
		end := 0
		oper := -1
		for ind, elem := range exp {
			if elem == "+" {
				start = ind - 1
				end = ind + 1
				oper = 1
				break

			} else if elem == "-" {
				start = ind - 1
				end = ind + 1
				oper = 2
				break
			}

		}
		tmp := exp[end+1:]
		if oper == 1 {
			tmp_e, err = Sum(exp[start : end+1])
			if err != nil {
				return nil, err
			}
			exp = append(exp[:start], tmp_e)
		} else if oper == 2 {
			tmp_e, err = Diff(exp[start : end+1])
			if err != nil {
				return nil, err
			}
			exp = append(exp[:start], tmp_e)
		}
		exp = append(exp, tmp...)

	}
	result, err := Split(strings.Replace(SumStr(exp), " ", "", -1))
	return result, err
}

func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0.0, ErrEmptyExp
	}
	var tmp []string
	s, err := Split(expression)
	if err != nil {
		return 0.0, err
	}

	tmp, err = (Math(s))
	if err != nil {
		return 0.0, err
	}
	res := SumStr(tmp)
	r, err := strconv.ParseFloat(res, 64)
	return r, err
}
func ParseNums(exp []string) (float64, float64, error) {
	tmp1, err1 := strconv.ParseFloat(exp[0], 64)
	if err1 != nil {
		return 0.0, 0.0, err1
	}
	tmp2, err2 := strconv.ParseFloat(exp[2], 64)
	if err2 != nil {
		return 0.0, 0.0, err2
	}
	return tmp1, tmp2, nil
}
func Mult(exp []string) (string, error) {

	tmp1, tmp2, err := ParseNums(exp)
	if err != nil {
		return "", err
	}
	return (strconv.FormatFloat(tmp1*tmp2, 'f', -1, 64)), nil
}
func Div(exp []string) (string, error) {

	tmp1, tmp2, err := ParseNums(exp)
	if err != nil {
		return "", err
	}
	if tmp2 == 0 {
		return "", ErrDivisionByZero
	}
	return (strconv.FormatFloat(tmp1/tmp2, 'f', 2, 64)), nil

}
func Sum(exp []string) (string, error) {
	tmp1, tmp2, err := ParseNums(exp)
	if err != nil {
		return "", err
	}

	return (strconv.FormatFloat(tmp1+tmp2, 'f', -1, 64)), nil
}
func Diff(exp []string) (string, error) {
	tmp1, tmp2, err := ParseNums(exp)
	if err != nil {
		return "", err
	}
	return (strconv.FormatFloat(tmp1-tmp2, 'f', -1, 64)), nil
}
