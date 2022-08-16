

build: 
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o bin/annofileserver-win-amd64.exe .
	GOOS=linux GOARCH=amd64 go build -o bin/annofileserver-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/annofileserver-darwin-amd64 .
	
run :
	nodemon --exec go run main.go --signal SIGTERM

delete :
	cd uploaded && rm -rf * && cd ..
