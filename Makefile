MAIN=cmd/main.go
DARWIN=bin/turnipin.darwin
EXE=bin/turnipin.exe

build_mac:
	go build -o $(DARWIN) $(MAIN)

build_win:
	go build -o $(EXE) $(MAIN)

run:
	go run $(MAIN)

gen:
	sqlc generate

tidy:
	go mod tidy

.PHONY: build_mac, build_win, run, gen