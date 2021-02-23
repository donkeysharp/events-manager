generate:
	go generate

generate2:
	cd db/migrations && go-bindata -pkg migrations -ignore generated.go -o generated.go .

build: generate
	go build
