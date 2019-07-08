GCC=go
GCMD=run
GPATH=main.go


run:
	make build
	$(GCC) $(GCMD) $(GPATH)

build:
	make build_db

build_db:
	db-migrate up
	rm pkg/db/db_structs.go
	go run pkg/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db/

install: 
	make install_routes

install_encryption:
	go get -u golang.org/x/crypto/bcrypt
install_routes:
	go get -u github.com/gorilla/mux
install_db:
	go get -u github.com/go-xorm/xorm
install_jwt:
	go get -u github.com/dgrijalva/jwt-go

create_keys:
	ssh-keygen -t rsa -b $(RA512) -f keys/app.rsa
	openssl rsa -in keys/app.rsa -pubout -outform PEM -out keys/app.rsa.pub
	