package controllers

import (
	"fmt"
	"strconv"
	"strings"
)

type Controller struct {
	DB string // db connection
}

func (c *Controller) ExpressionController(paths []string) (string, error) {
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
