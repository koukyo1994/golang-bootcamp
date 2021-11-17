package main

import (
	"bootcamp/ch07/ex7.16/eval"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func dropDuplicates(vars []string) map[string]bool {
	varSet := make(map[string]bool)
	for _, v := range vars {
		varSet[v] = true
	}
	return varSet
}

func evaluateExpression(w http.ResponseWriter, r *http.Request) {
	exprStr := r.URL.Query().Get("expr")
	exprStr, err := url.QueryUnescape(exprStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse query expr %s: %v", r.URL.Query().Get("expr"), err), http.StatusBadRequest)
		return
	}

	var envVars map[string]float64
	err = json.Unmarshal([]byte(r.URL.Query().Get("env")), &envVars)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse query env %s: %v", r.URL.Query().Get("env"), err), http.StatusBadRequest)
		return
	}

	expr, err := eval.Parse(exprStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse expression %s: %v", exprStr, err), http.StatusBadRequest)
		return
	}
	requiredVars := dropDuplicates(expr.Vars())
	for key := range requiredVars {
		if _, ok := envVars[key]; !ok {
			http.Error(w, fmt.Sprintf("insufficient variable: %s", key), http.StatusBadRequest)
			return
		}
	}

	variables := make(map[eval.Var]float64)
	for k, v := range envVars {
		variables[eval.Var(k)] = v
	}
	fmt.Fprintf(w, "%v", expr.Eval(eval.Env(variables)))
}

func main() {
	http.HandleFunc("/evaluate", evaluateExpression)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
