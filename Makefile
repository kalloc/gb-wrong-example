all: deps test build


deps:
	gb vendor restore

        
build:
	env GOOS=darwin GOARCH=amd64 gb build
	env GOOS=linux GOARCH=amd64 gb build

test:
	cd src/example && GOPATH=${PWD}:${PWD}/vendor/ go test -cover -test.coverprofile=../../reports/cover.out

