CC=solc
CFLAGS=--combined-json abi,asm,ast,bin
SRC=contracts/src/*.sol
JSON=contracts.json
OUT=contracts.go

all: $(SRC)
		$(CC) $(CFLAGS) $(SRC) > $(JSON)
		abigen -combined-json $(JSON) -pkg evm -out $(OUT)

clean:
		rm -f $(JSON) $(OUT)
