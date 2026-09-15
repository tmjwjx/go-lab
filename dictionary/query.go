package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func query(word string) string {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个包含 JSON 数据的 Reader
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)

	// 创建一个新的 HTTP POST 请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}

	// 设置 HTTP 请求头
	{
		req.Header.Set("Connection", "keep-alive")                                                                                                               // 保持连接
		req.Header.Set("DNT", "1")                                                                                                                               // 请勿追踪
		req.Header.Set("os-version", "")                                                                                                                         // 操作系统版本
		req.Header.Set("sec-ch-ua-mobile", "?0")                                                                                                                 // 是否为移动设备
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36") // 用户代理
		req.Header.Set("app-name", "xy")                                                                                                                         // 应用名称
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")                                                                                         // 内容类型
		req.Header.Set("Accept", "application/json, text/plain, */*")                                                                                            // 接受的内容类型
		req.Header.Set("device-id", "")                                                                                                                          // 设备 ID
		req.Header.Set("os-type", "web")                                                                                                                         // 操作系统类型
		req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")                                                                                          // 授权令牌
		req.Header.Set("Origin", "https://fanyi.caiyunapp.com")                                                                                                  // 请求来源
		req.Header.Set("Sec-Fetch-Site", "cross-site")                                                                                                           // 资源获取站点
		req.Header.Set("Sec-Fetch-Mode", "cors")                                                                                                                 // 资源获取模式
		req.Header.Set("Sec-Fetch-Dest", "empty")                                                                                                                // 资源获取目标
		req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")                                                                                                // 引用页
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")                                                                                                      // 接受的语言
		req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	}

	// 发送 HTTP 请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 检查响应状态码
	if resp.StatusCode != 200 {
		log.Fatal("bad status code:", resp.StatusCode, "body:", string(bodyText))
	}

	// 打印响应体内容
	//fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dictResponse)

	//fmt.Println("____________________________________________________")
	//fmt.Println("单词:", dictResponse.Dictionary.Entry)
	//fmt.Println("UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	//for _, i := range dictResponse.Dictionary.Explanations {
	//	fmt.Println(i)
	//}

	output := "____________________________________________________\n" +
		fmt.Sprintf("单词: %s\n", dictResponse.Dictionary.Entry) +
		fmt.Sprintf("UK: %s US: %s\n", dictResponse.Dictionary.Prons.En, dictResponse.Dictionary.Prons.EnUs)

	for _, explanation := range dictResponse.Dictionary.Explanations {
		output += fmt.Sprintf("%s\n", explanation)
	}
	output += fmt.Sprintf("----------------------------------------------------\n")
	return output

}
