.PHONY: mock
mock:
	mocker --prefix "" --dst mock/api_default.go --pkg mock --selfpkg github.com/confluentinc/schema-registry-sdk-go api_default.go DefaultApi
