TARGET=fetchall

all: 
	go build -o $(TARGET)  main.go 
	./fetchall https://golang.org http://gopl.io https://godoc.org \

clean:
	rm $(TARGET) output.txt

