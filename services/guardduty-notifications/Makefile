.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/guard-duty-notifer ./aws-guardduty-notifier

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
