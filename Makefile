run :
	nodemon --exec go run main.go --signal SIGTERM

delete :
	cd uploaded && rm -rf * && cd ..
