all:
	go build clock.go 
	./clock 8000 &
	./clock 5000 &
	./clock 3000 &
	./clock 3400 

kill:
	killall clock

clean:
	rm clock
