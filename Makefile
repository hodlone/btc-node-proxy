build:
	go build -o main .

run: build
	./main

watch:
	ulimit -n 1000 #increase the file watch limit
	reflex -s -r '\.go$$' make run