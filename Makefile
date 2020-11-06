
docker-build:
	docker build --no-cache -t gcr.io/mchirico/go-spreadsheet:test -f Dockerfile .

push:
	docker push gcr.io/mchirico/go-spreadsheet:test

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/mchirico/go-spreadsheet:test
