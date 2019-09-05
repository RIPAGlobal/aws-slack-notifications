.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/codebuild-notifier ./aws-codebuild-notifier
	env GOOS=linux go build -ldflags="-s -w" -o bin/codepipeline-notifier ./aws-codepipeline-notifier

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --force --verbose
