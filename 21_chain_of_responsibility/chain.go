package chain

import (
	"fmt"
	"strings"
)

/*
	chain_of_responsibility 责任链模式：以消息过滤器为例
*/

//Handler 抽象责任链处理器
type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}

//AdHandler 广告过滤处理器
type AdHandler struct {
	handler Handler
}

//Handle 广告过滤处理方法
func (ad *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content, "广告", "**", 1)
	fmt.Println(newContent)
	ad.next(ad.handler, newContent)
}

//next 广告过滤逻辑递归处理
func (ad *AdHandler) next(handler Handler, content string) {
	if ad.handler != nil {
		ad.handler.Handle(content)
	}
}

//YellowHandler 涉黄过滤处理器
type YellowHandler struct {
	handler Handler
}

//Handle 涉黄过滤处理方法
func (yellow *YellowHandler) Handle(content string) {
	fmt.Println("执行涉黄过滤。。。")
	newContent := strings.Replace(content, "涉黄", "**", 1)
	fmt.Println(newContent)
	yellow.next(yellow.handler, newContent)
}

//next 涉黄过滤逻辑递归处理
func (yellow *YellowHandler) next(handler Handler, content string) {
	if yellow.handler != nil {
		yellow.handler.Handle(content)
	}
}

//SensitiveHandler 敏感词过滤处理器
type SensitiveHandler struct {
	handler Handler
}

//Handle 敏感词过滤处理方法
func (sensitive *SensitiveHandler) Handle(content string) {
	fmt.Println("执行敏感词过滤。。。")
	newContent := strings.Replace(content, "敏感词", "***", 1)
	fmt.Println(newContent)
	sensitive.next(sensitive.handler, newContent)
}

//next 敏感词过滤逻辑递归处理
func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if sensitive.handler != nil {
		sensitive.handler.Handle(content)
	}
}
