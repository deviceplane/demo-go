#!/bin/bash
set -e

echo $VERSION

docker build -t deviceplane/demo-go:amd64-${VERSION} -f dockerfiles/amd64/Dockerfile --build-arg version=${VERSION} .
docker build -t deviceplane/demo-go:arm32v6-${VERSION} -f dockerfiles/arm32v6/Dockerfile --build-arg version=${VERSION} .
docker build -t deviceplane/demo-go:arm64v8-${VERSION} -f dockerfiles/arm64v8/Dockerfile --build-arg version=${VERSION} .

docker push deviceplane/demo-go:amd64-${VERSION}
docker push deviceplane/demo-go:arm32v6-${VERSION}
docker push deviceplane/demo-go:arm64v8-${VERSION}

docker manifest create deviceplane/demo-go:${VERSION} deviceplane/demo-go:amd64-${VERSION} deviceplane/demo-go:arm32v6-${VERSION} deviceplane/demo-go:arm64v8-${VERSION}
docker manifest annotate deviceplane/demo-go:${VERSION} deviceplane/demo-go:arm32v6-${VERSION} --os linux --arch arm
docker manifest annotate deviceplane/demo-go:${VERSION} deviceplane/demo-go:arm64v8-${VERSION} --os linux --arch arm64
