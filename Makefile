EXEC=georacle
EVM=chain/evm
CHAINS=evm

build: $(CHAINS)
	go build ./cmd/...

install: $(CHAINS)
	go install ./cmd/...

start: build
	./$(EXEC) start

evm: $(EVM)
	cd $(EVM) && $(MAKE)

docker: build
	docker build -t $(EXEC) .

clean: $(EVM)
	rm -f $(EXEC)

	cd $(EVM) && $(MAKE) clean
