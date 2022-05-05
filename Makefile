EXEC=georacle
EVM=chain/evm

build:
	go build -o $(EXEC)

start: build
	./$(EXEC) start

evm: $(EVM)
	cd $(EVM) && $(MAKE)

clean: $(EVM)
	rm -f $(EXEC)

	cd $(EVM) && $(MAKE) clean
