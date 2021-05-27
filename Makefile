project=datastruct-go
goPath=$(shell go env GOPATH)
#是否用gdb调试
debugParam=-gcflags "-N -l"
#压缩二进制，-s 去掉符号信息， -w 去掉DWARF调试信息
delDebug=-ldflags "-s -w"

ifeq ($(LANG),)
	execFile=$(project).exe
	cleanExec=del $(goPath)\bin\$(execFile)
	compress=upx.exe
else
	execFile=$(project)
	cleanExec=rm -rf $(goPath)/bin/$(execFile)
	#压缩二进制
    compress=upx
endif
install=$(goPath)/bin/$(execFile)

build:
	go build -o $(execFile) $(delDebug) main.go
	$(compress) $(execFile)

debug:
	go build -o $(execFile) $(debugParam) main.go

install:
	go install $(delDebug)
	$(compress) $(install)

.PHONY: clean

clean:
	$(cleanExec) $(execFile)
