.PHONY: fetch.out findlinks1.out outline.out findlinks1 outline ex5.1 ex5.2 ex5.3 ex5.4 findlinks2 ex5.5 ex5.6 ex5.10 ex5.11 ex5.12 ex5.13

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

findlinks2:
	go run findlinks2/main.go https://gopl.io https://golang.org/

ex5.1: fetch.out
	go build -o ex5.1.out ex5.1/main.go
	./fetch.out https://golang.org | ./ex5.1.out

ex5.2: fetch.out
	go build -o ex5.2.out ex5.2/main.go
	./fetch.out https://golang.org | ./ex5.2.out

ex5.3: fetch.out
	go build -o ex5.3.out ex5.3/main.go
	./fetch.out https://gopl.io | ./ex5.3.out

ex5.4: fetch.out
	go build -o ex5.4.out ex5.4/main.go
	./fetch.out https://gopl.io | ./ex5.4.out

ex5.5:
	go run ex5.5/main.go https://golang.org https://gopl.io

ex5.6:
	go run ex5.6/main.go > ./assets/ex5.6.svg

ex5.10:
	go run ex5.10/main.go

ex5.11:
	go run ex5.11/main.go

ex5.12:
	go run ex5.12/main.go https://gopl.io

ex5.13:
	go run ex5.13/main.go https://gopl.io
