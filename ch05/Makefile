.PHONY: fetch.out findlinks1.out outline.out findlinks1 outline ex5.1 ex5.2 ex5.3

fetch.out:
	go build -o fetch.out fetch/main.go

findlinks1.out:
	go build -o findlinks1.out findlinks1/main.go

outline.out:
	go build -o outline.out outline/main.go

findlinks1: fetch.out findlinks1.out
	./fetch.out https://golang.org | ./findlinks1.out

outline: fetch.out outline.out
	./fetch.out https://golang.org | ./outline.out

ex5.1:
	go build -o ex5.1.out ex5.1/main.go
	./fetch.out https://golang.org | ./ex5.1.out

ex5.2:
	go build -o ex5.2.out ex5.2/main.go
	./fetch.out https://golang.org | ./ex5.2.out

ex5.3:
	go build -o ex5.3.out ex5.3/main.go
	./fetch.out https://gopl.io | ./ex5.3.out
