export GO111MODULE=on

default: test

ci: depsdev test

test:
	go test -v ./... -coverprofile=coverage.out -covermode=count

lint:
	golangci-lint run ./...

depsdev:
	go install github.com/Songmu/ghch/cmd/ghch@latest
	go install github.com/Songmu/gocredits/cmd/gocredits@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

prerelease:
	ghch -w -N ${VER}
	gocredits . > CREDITS
	git add CHANGELOG.md CREDITS
	git commit -m'Bump up version number'
	git tag ${VER}

.PHONY: default test
