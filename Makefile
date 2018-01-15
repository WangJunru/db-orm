all:

init:
	go get -u github.com/jteeuwen/go-bindata/...

debugTpl:
	go-bindata -nometadata -o tpl/bindata.go -ignore bindata.go -pkg tpl -debug tpl

buildTpl:
	go-bindata -nometadata -o tpl/bindata.go -ignore bindata.go -pkg tpl tpl

build:
	go build

install:
	go install

sql:
	go install
	db-orm sql -i ./example/yaml/ -o ./example/script/

test:
	go install
	db-orm code -i ./example/yaml/ -o ./example/model/
	# go test -v ./... ???
	#go test -v ./orm/sqlbuilder
	go test -v ./example/model/...
