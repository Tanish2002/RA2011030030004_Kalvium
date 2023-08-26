package controllers

import (
	"fmt"
	"go/token"
	"go/types"
	"math"
	"strconv"
	"strings"
)

type Controller struct {
	DB string // db connection
}

func (c *Controller) ExpressionController(paths []string) (float64, error) {
	expression, err := expressionMaker(paths)
	if err != nil {
		return math.NaN(), err
	}
	evaluatedExpr, err := expressionEvaluator(expression)
	return evaluatedExpr, err
}

func expressionEvaluator(expression string) (float64, error) {
	fset := token.NewFileSet()
	eval, err := types.Eval(fset, nil, token.NoPos, expression)
	if err != nil {
		return math.NaN(), err
	}
	value, err := strconv.ParseFloat(eval.Value.ExactString(), 64)
	if err != nil {
		return math.NaN(), err
	}
	return value, nil

}

func expressionMaker(paths []string) (string, error) {
	for i, v := range paths {
		_, err := strconv.ParseFloat(v, 64)
		// If value is a number then skip it
		if err == nil {
			continue
		}

		// replace textual operators with real operators
		switch v {
		case "plus":
			paths[i] = "+"
		case "minus":
			paths[i] = "-"
		case "into":
			paths[i] = "*"
		case "div":
			paths[i] = "/"
		default:
			return "", fmt.Errorf("Invalid arithmetic operator %s", paths[i])
		}

	}
	return strings.Join(paths, ""), nil
}
