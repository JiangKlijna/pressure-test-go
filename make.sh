
function cross_compile() {
	export GOOS=$1
	export GOARCH=$2
	TARGET_FILE=rainx_$1_$2
	if [ $GOOS = windows ]; then
		TARGET_FILE=$TARGET_FILE.exe
	fi
	go build -i -o out/$TARGET_FILE -ldflags "-s -w"
	echo compile $TARGET_FILE done
}

function publish() {
	cross_compile darwin 386
	cross_compile darwin amd64

	cross_compile linux 386
	cross_compile linux amd64
	cross_compile linux arm
	cross_compile linux arm64
	cross_compile linux mips
	cross_compile linux mips64

	cross_compile windows 386
	cross_compile windows amd64
}

if [ "$1" == "" ]; then
    go build -ldflags "-s -w"
elif [ $1 == run ]; then
    go build -o out/tmp.out && out/tmp.out
elif [ $1 == test ]; then
    go test
elif [ $1 == clean ]; then
    rm -rf out
elif [ $1 == publish ]; then
    publish
else
    echo "Example:"
	echo "	make"
	echo "	make run"
	echo "	make test"
	echo "	make clean"
	echo "	make publish"
fi