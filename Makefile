################## General Variables ##################

VERSION=v0.0.1

################# Install Variables ###################

BIN_PATH=$(HOME)/.local/bin
## the reslpath resolve the problem when we set pass with/without backslash on end of Installation path
INSTALATION_PATH=$(shell realpath $(BIN_PATH))
COMPLETION_FILE=completion.bash
BASH_COMPLETION_PATH=/etc/bash_completion.d

################## Build Variables ####################

BIN=ata
GO_C=go

SRC=main.go
LD_FLAGS:=-X '$(BIN).Version=$(VERSION)'

################# Build Targets #######################

$(BIN): $(SRC)
	$(GO_C) build -ldflags="$(LD_FLAGS)" -o $@

################ PHONY Targets ##################

.PHONY: clean run test install

install: $(BIN)
	mkdir -p $(INSTALATION_PATH)
	sudo install $(BIN) $(INSTALATION_PATH)/$(BIN)
	sudo install $(COMPLETION_FILE) $(BASH_COMPLETION_PATH)/$(BIN)_$(COMPLETION_FILE)

uninstall:
	test -e "$(INSTALATION_PATH)/$(BIN)" && sudo rm "$(INSTALATION_PATH)/$(BIN)" || true
	test -e "$(BASH_COMPLETION_PATH)/$(BIN)_$(COMPLETION_FILE)" && sudo rm "$(BASH_COMPLETION_PATH)/$(BIN)_$(COMPLETION_FILE)" || true

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
	
