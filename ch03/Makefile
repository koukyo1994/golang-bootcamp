.PHONY: surface mandelbrot ex3.1 ex3.2 ex3.3 ex3.4 ex3.5 ex3.6

surface:
	go run surface/main.go > assets/surface.svg

mandelbrot:
	go run mandelbrot/main.go > assets/mandelbrot.png

julia:

ex3.1:
	go run ex3.1/main.go > assets/ex3.1.svg

ex3.2:
	go run ex3.2/main.go -shape="saddle" > assets/ex3.2-saddle.svg
	go run ex3.2/main.go -shape="bump" > assets/ex3.2-bump.svg
	go run ex3.2/main.go -shape="egg pack" > assets/ex3.2-eggpack.svg

ex3.3:
	go run ex3.3/main.go > assets/ex3.3.svg

ex3.4:
	go run ex3.4/main.go

ex3.5:
	go run ex3.5/main.go > assets/ex3.5.png

ex3.6:
	go run ex3.6/main.go > assets/ex3.6.png
