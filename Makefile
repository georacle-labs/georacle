EXEC=georacle

build:
	go build -o $(EXEC)

start: build
	./$(EXEC) start

clean:
	rm -f $(EXEC)
