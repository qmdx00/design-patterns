GO = GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go

all: format test

.PHONY: all

format:
	@${GO} fmt **/*.go

test:
	@${GO} test **/*.go


