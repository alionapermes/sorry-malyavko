SOURCE = cmd/main.go
TARGET = sorry-malyavko

app: clear
	go build -o ${TARGET} ${SOURCE}

clear:
	rm -f ${TARGET}
