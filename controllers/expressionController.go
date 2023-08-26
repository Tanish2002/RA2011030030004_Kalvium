package controllers

import (
	"fmt"
	"kalvium/models"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
)

type Controller struct {
	Model models.Model
}

func (c *Controller) ExpressionController(paths []string) (float64, string, error) {
	// Create the expression
	expression, err := expressionMaker(paths)
	if err != nil {
		return math.NaN(), "", err
	}

	// Evaluate the expression
	evaluatedExprInterface, err := expressionEvaluator(expression)
	if err != nil {
		return math.NaN(), "", err
	}
	evaluatedExpr, ok := evaluatedExprInterface.(float64)
	if !ok {
		return math.NaN(), "", err
	}

	// Add to Request to DB
	c.Model.AddToRequestLog(&models.RequestLog{
		Question:  expression,
		Answer:    evaluatedExpr,
		Timestamp: time.Now(),
	})

	return evaluatedExpr, expression, err
}

func expressionEvaluator(expression string) (interface{}, error) {
	evalExpress, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return math.NaN(), err
	}
	result, err := evalExpress.Evaluate(nil)
	if err != nil {
		return math.NaN(), err
	}
	return result, nil
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
		case "modulo":
			paths[i] = "%"
		case "lbrack":
			paths[i] = "("
		case "rbrack":
			paths[i] = ")"
		case "pow":
			paths[i] = "**"
		case "lshift":
			paths[i] = "<<"
		case "rshift":
			paths[i] = ">>"
		default:
			return "", fmt.Errorf("Invalid arithmetic operator %s", paths[i])
		}

	}
	return strings.Join(paths, ""), nil
}
