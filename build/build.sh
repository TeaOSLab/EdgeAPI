#!/usr/bin/env bash

function build() {
	ROOT=$(dirname $0)
	NAME="edge-api"
	DIST=$ROOT/"../dist/${NAME}"
	OS=${1}
	ARCH=${2}

	if [ -z $OS ]; then
		echo "usage: build.sh OS ARCH"
		exit
	fi
	if [ -z $ARCH ]; then
		echo "usage: build.sh OS ARCH"
		exit
	fi

	VERSION=$(lookup-version $ROOT/../internal/const/const.go)
	ZIP="${NAME}-${OS}-${ARCH}-v${VERSION}.zip"

	# check edge-node
	NodeVersion=$(lookup-version $ROOT"/../../EdgeNode/internal/const/const.go")
	echo "building edge-node v${NodeVersion} ..."
	EDGE_NODE_BUILD_SCRIPT=$ROOT"/../../EdgeNode/build/build.sh"
	if [ ! -f $EDGE_NODE_BUILD_SCRIPT ]; then
		echo "unable to find edge-node build script 'EdgeNode/build/build.sh'"
		exit
	fi
	cd $ROOT"/../../EdgeNode/build"
	echo "=============================="
	architects=("amd64" "386" "arm64" "mips64" "mips64le")
	for arch in "${architects[@]}"; do
		./build.sh linux $arch
	done
	echo "=============================="
	cd -

	rm -f $ROOT/deploy/*.zip
	for arch in "${architects[@]}"; do
		cp $ROOT"/../../EdgeNode/dist/edge-node-linux-${arch}-v${NodeVersion}.zip" $ROOT/deploy/
	done

	# build sql
	echo "building sql ..."
	${ROOT}/sql.sh

	# copy files
	echo "copying ..."
	if [ ! -d $DIST ]; then
		mkdir $DIST
		mkdir $DIST/bin
		mkdir $DIST/configs
		mkdir $DIST/logs
	fi
	cp $ROOT/configs/api.template.yaml $DIST/configs/
	cp $ROOT/configs/db.template.yaml $DIST/configs/
	cp -R $ROOT/deploy $DIST/
	rm -f $dist/deploy/.gitignore
	cp -R $ROOT/installers $DIST/
	cp -R $ROOT/resources $DIST/
	rm -f $DIST/resources/ipdata/ip2region/global_region.csv
	rm -f $DIST/resources/ipdata/ip2region/ip.merge.txt

	# building installer
	echo "building installer ..."
	architects=("amd64" "386")
	for arch in "${architects[@]}"; do
		# TODO support arm, mips ...
		env GOOS=linux GOARCH=${arch} go build --ldflags="-s -w" -o $ROOT/installers/edge-installer-helper-linux-${arch} $ROOT/../cmd/installer-helper/main.go
	done

	# building api node
	env GOOS=$OS GOARCH=$ARCH go build --ldflags="-s -w" -o $DIST/bin/edge-api $ROOT/../cmd/edge-api/main.go

	# delete hidden files
	find $DIST -name ".DS_Store" -delete
	find $DIST -name ".gitignore" -delete

	echo "zip files"
	cd "${DIST}/../" || exit
	if [ -f "${ZIP}" ]; then
		rm -f "${ZIP}"
	fi
	zip -r -X -q "${ZIP}" ${NAME}/
	rm -rf ${NAME}
	cd - || exit

	echo "[done]"
}

function lookup-version() {
	FILE=$1
	VERSION_DATA=$(cat $FILE)
	re="Version[ ]+=[ ]+\"([0-9.]+)\""
	if [[ $VERSION_DATA =~ $re ]]; then
		VERSION=${BASH_REMATCH[1]}
		echo $VERSION
	else
		echo "could not match version"
		exit
	fi
}

build $1 $2
