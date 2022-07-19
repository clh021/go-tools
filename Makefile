.PHONY: init
init: .git/hooks/pre-push

.git/hooks/pre-push:makefile
	@echo "#!/usr/bin/env bash" > $@
	@echo "set -e" >> $@
	@echo "make test" >> $@
	@echo "cd web && npm run lint" >> $@
	@chmod a+x $@

.PHONY: install-depend
install-depend:
	go get -u github.com/cosmtrek/air

.PHONY: .air
.air:
	touch .air

.PHONY: build-bindata
# rm -rf web/dist/js/*.map # 仅在最终发布包时用于优化软件包大小
build-bindata:
	cd web; npm run build
	go-bindata -fs -prefix "web/dist" --pkg "webassets" -o "./services/webAssets/bindata.go" web/dist/...

.PHONY: build
gitTime=$(shell date +%Y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD | cut -c1-8)
gitTag=$(shell git tag --list --sort=version:refname 'v*' | tail -1)
gitCount=$(shell git log --pretty=format:'' | wc -l)/$(shell git rev-list --all --count)
buildStr=${gitTime}.${gitCID}.${gitTag}.${gitCount}
build: 
	go mod tidy
	go generate
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-X main.build=${buildStr}" -o bin/app


APPINTO=uploadAdvanced

.PHONY: runn
runn: p1="./bin/child.amd64"
runn: p2="./bin/main"
runn:
	APPINTO=${APPINTO} ./bin/app
# @echo 'import subprocess; [p.wait() for p in subprocess.Popen(${p1}),subprocess.Popen(${p2})]' | python2

.PHONY: dev
dev: p1="air"
dev: p2=["sh", "-c", "APPINTO=${APPINTO} ./bin/app"]
dev: .air bin
	air

.PHONY: bin
bin:
	mkdir -p bin