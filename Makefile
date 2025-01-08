kt:
	@kitex -module github.com/lianzhilu/chat-paper -gen-path=cp-core/kitex/kitex_gen cp-core/kitex/idl/service.thrift

build-cp-core:
	export CPCORE_CONF_DIR=$(shell pwd)/cp-core/conf/
	@go build -o output/cp-core ./cp-core/cmd