SHARED_OBJECT=libketos-hookroot.so
LIBRARY_PATH=/usr/lib

hookroot:
	go build -buildmode=c-shared -o ${SHARED_OBJECT} .

install:
	cp ${SHARED_OBJECT} ${LIBRARY_PATH}/
