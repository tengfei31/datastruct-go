project=datastruct-go
goPath=$(shell go env GOPATH)
#是否用gdb调试
debugParam=-gcflags "-N -l"
#压缩二进制，-s 去掉符号信息， -w 去掉DWARF调试信息
delDebug=-ldflags "-s -w"

#linux
execFile=$(project)
cleanExec=rm -rf $(goPath)/bin/$(execFile) $(execFile)
#压缩二进制
compress=upx

ifeq ($(LANG),)
	ifeq ($(shell uname), Darwin)
		#macOS
	else
		#windows
		execFile=$(project).exe
		cleanCommand=del
		ifeq ($(goPath)\bin\$(execFile), $(wildcard $(goPath)\bin\$(execFile)))
			delInstall=$(goPath)\bin\$(execFile)
		else
			delInstall=
		endif
		ifeq ($(execFile), $(wildcard $(execFile)))
			delBuild=$(execFile)
		else
			delBuild=
		endif
		cleanExec=$(cleanCommand) $(delBuild) $(delInstall)
		compress=upx.exe
	endif
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
	$(cleanExec)
