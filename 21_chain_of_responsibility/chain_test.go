package chain

import "testing"

func TestChain(t *testing.T) {
	adHandler := &AdHandler{}
	yellowHandler := &YellowHandler{}
	sensitiveHandler := &SensitiveHandler{}
	// 责任链串联
	adHandler.handler = yellowHandler
	yellowHandler.handler = sensitiveHandler

	adHandler.Handle("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}
