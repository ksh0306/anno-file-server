

gencert:
	# openssl ecparam -genkey -name secp384r1 -out server.key
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365
	# openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 # no good for apple
	
build: 
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o bin/annofileserver-win-amd64.exe .
	GOOS=linux GOARCH=amd64 go build -o bin/annofileserver-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/annofileserver-darwin-amd64 .

upload:
	scp -P 51123 bin/annofileserver-linux-amd64 server.crt server.key annowiz@222.110.65.138:app
	# ssh -p 51123 annowiz@222.110.65.138 chmod 644 app/annofileserver-linux-amd64

remote:
	ssh -p 51123 annowiz@222.110.65.138

run :
	nodemon --exec go run main.go -port=8888:8443 --signal SIGTERM

delete :
	cd uploaded && rm -rf * && cd ..
