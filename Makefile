all: push

TAG = 1.1.0
PREFIX = quay.io/mgoodness/app-healthz

build: main.go clean
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix netgo \
		--ldflags '-extldflags "-static"' -tags netgo .

clean:
	rm -f app-healthz

container: build
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)
