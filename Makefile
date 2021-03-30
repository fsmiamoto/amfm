BIN=amfm

$(BIN):
	go build -o $(BIN)

clean:
	rm $(BIN)
