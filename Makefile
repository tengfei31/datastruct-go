project=datastruct-go
go_path=`go env GOPATH`
#是否用gdb调试
debugParam=-gcflags "-N -l"

ifeq ($(LANG),)
	execFile=$(project).exe
else
	execFile=$(project)
endif

build:
	go build -o $(execFile) main.go


debug:
	go build -o $(execFile) $(debugParam) main.go

install:
	go install

.PHONY: clean

clean:
	-rm -rf $(execFile) \
		$(go_path)/bin/$(execFile)
