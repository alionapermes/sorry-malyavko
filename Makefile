SOURCE = cmd/main.go
TARGET = sorry-malyavko

linux: clear
	go build -o ${TARGET} ${SOURCE}

shitos: clear
	GOOS=windows go build -o ${TARGET}.exe ${SOURCE}

all: linux shitos

clear:
	rm -f ${TARGET}
