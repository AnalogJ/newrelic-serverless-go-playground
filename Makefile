.PHONY: clean deploy remove offline

clean:
	rm -rf ./bin ./vendor

deploy: clean domain
	npx serverless deploy --verbose

remove: clean build
	npx serverless remove --verbose

offline: clean build
	sls offline --useDocker --localEnvironment --noAuth
