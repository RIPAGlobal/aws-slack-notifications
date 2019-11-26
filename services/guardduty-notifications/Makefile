.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/guardduty-notifier ./aws-guardduty-notifier

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
