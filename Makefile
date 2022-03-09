################## General Variables ##################

VERSION=v0.0.1

################## Build Variables ####################

BIN=main
GO_C=go

SRC=main.go
LD_FLAGS:=-X '$(BIN).Version=$(VERSION)'

################# Build Targets #######################

$(BIN): $(SRC)
	$(GO_C) build -ldflags="$(LD_FLAGS)"

################ PHONY Targets ##################

.PHONY: clean run test

clean:
	rm $(BIN)

run: $(BIN)
	./$(BIN)

test: $(BIN)
	@ echo "Testing option \"version\""
	./$(BIN) version
	@ echo
	@ echo "Testing option \"-v\""
	./$(BIN) -v
	@ echo
	@ echo "Testing option \"help\""
	./$(BIN) help
	@ echo
	@ echo "Testing option \"-h\""
	./$(BIN) -h
	@ echo
	@ echo "Testing option \"adb\""
	./$(BIN) adb
	@ echo
	@ echo "Testing option \"log\""
	./$(BIN) log
	@ echo
	@ echo "Testing a invalid option:"
	./$(BIN) invalid && false || true
	
