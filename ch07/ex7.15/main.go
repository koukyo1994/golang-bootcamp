package main

import (
	"bootcamp/ch07/ex7.15/eval"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func dropDuplicates(vars []string) map[string]bool {
	varSet := make(map[string]bool)
	for _, v := range vars {
		varSet[v] = true
	}
	return varSet
}

func main() {
	var s string
	fmt.Printf("expr: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()

	expr, err := eval.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing expr %s: %v\n", s, err)
		os.Exit(1)
	}

	requiredVars := dropDuplicates(expr.Vars())
	variablesMap := make(map[eval.Var]float64)
	for key := range requiredVars {
	again:
		fmt.Printf("%s: ", key)
		fmt.Scan(&s)
		value, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("invalid input: %s\nretry...\n", s)
			goto again
		}
		variablesMap[eval.Var(key)] = value
	}
	env := eval.Env(variablesMap)
	result := fmt.Sprintf("%.6g", expr.Eval(env))
	fmt.Printf("%s in %v => %s\n", expr, env, result)
}
