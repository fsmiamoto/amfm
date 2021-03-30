BIN=amfm

$(BIN): *.go
	go build -o $(BIN)

clean:
	rm $(BIN)
