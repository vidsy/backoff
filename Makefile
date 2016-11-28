BRANCH = "master"
GO_BUILDER_IMAGE = "vidsyhq/go-builder"
PATH_BASE = "/go/src/github.com/vidsy"
REPONAME = "backoff"
VERSION = $(shell cat ./VERSION)

DEFAULT: test

check-version:
	git fetch && (! git rev-list ${VERSION})]

push-tag:
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@docker run \
	-it \
	--rm \
	-v "${CURDIR}":${PATH_BASE}/${REPONAME} \
	-w ${PATH_BASE}/${REPONAME} \
	--entrypoint=go \
	${GO_BUILDER_IMAGE} test .

test-ci:
	@docker run \
	-v "${CURDIR}":${PATH_BASE}/${REPONAME} \
	-w ${PATH_BASE}/${REPONAME} \
	--entrypoint=go \
	${GO_BUILDER_IMAGE} test . -cover

test-coverage:
	@go test -covermode=count -coverprofile=go.backoff.coverage.out .
	@go tool cover -html=go.backoff.coverage.out
