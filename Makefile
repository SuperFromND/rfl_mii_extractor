EXECNAME = rfl_mii_extractor

ifeq ($(OS),Windows_NT)
	EXECNAME := $(addsuffix .exe,$(EXECNAME))
endif

all:
	go build -trimpath -o bin/$(EXECNAME) src/main.go

release:
	go build -trimpath -o bin/$(EXECNAME) -ldflags "-X main.version=1.0.0" src/main.go
    
run:
	go build -trimpath -o bin/$(EXECNAME) src/main.go
	cd bin && ./rfl_mii_extractor.exe