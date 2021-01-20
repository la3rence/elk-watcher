APP=elk-watcher

.PHONY: build
build: clean
	go build -ldflags '-w -s' -o release/${APP}

.PHONY: run
run:
	env go run watcher.go

.PHONY: clean
clean:
    ifneq ($(wildcard release/${APP}),)
		rm -rf release/
    endif

.PHONY: docker
docker:
	docker build . --file Dockerfile
