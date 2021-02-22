project=datastruct-go
go_path=`go env GOPATH`
#是否用gdb调试
debugParam=-gcflags "-N -l"

build:
	go build -o $(project) main.go

debug:
	go build -o $(project) $(debugParam) main.go

install:
	go install

.PHONY: clean

clean:
	-rm -rf $(project) \
			$(go_path)/bin/$(project) \
			$(go_path)/bin/$(project).exe
