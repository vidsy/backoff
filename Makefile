BRANCH = "master"
REPONAME = "backoff"
VERSION = $(shell cat ./VERSION)

DEFAULT: test

push-tag:
	@echo "=> New tag version: v${VERSION}"
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag v${VERSION}
	git push origin v${VERSION}

test:
	@GO111MODULE=on go test ./... -cover
