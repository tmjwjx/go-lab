package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Content struct {
	Content string `json:"content"`
}

// 定义与 JSON 数据相对应的结构体
type ApiResponse struct {
	Result   Result        `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

type Result struct {
	Response string `json:"response"`
	Usage    Usage  `json:"usage"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func api(w http.ResponseWriter, r *http.Request) {

	// 解析 JSON 数据
	var content Content
	err := json.NewDecoder(r.Body).Decode(&content)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "解析 JSON 数据失败", http.StatusBadRequest)
		return
	}
	//fmt.Println(content.Content)

	// 发送 HTTP 请求
	url := "https://api.cloudflare.com/client/v4/accounts/84bbf5f1b19d1f86a63f71761e9acd62/ai/run/@cf/meta/llama-3-8b-instruct"
	method := "POST"

	str := fmt.Sprintf(`{
	 "model": "@cf/meta/llama-3.1-8b-instruct",
	 "messages": [
	   {
	     "role": "user",
	     "content": "%s"
	   }
	 ]
	}`, content.Content)

	payload := strings.NewReader(str)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer G-PpN6ne8PU3eV89rAgvc78Ofebpt_-cKYXKVdDa")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.cloudflare.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "__cf_bm=.fcEaI4xRTqBluYt.BvQ_tvkmJfGQqz2ITjYHvhcMQA-1739189033-1.0.1.1-uIiXS2tl9VwTpqau8m_TD4PQT61Nmho0U8mZlSwkEl4LCcT1C_X0rL37AJkaoQAWWTkI_U.UiP20PRZk7phwqQ; __cfruid=f3cfee64ec6ea2c45bd836f04310f468a37a5bbb-1739187618; _cfuvid=DlHjvyKvSYzzkL0knVF_YPSFnGSKucpKmiv2dqrEsAc-1739187618610-0.0.1.1-604800000")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 转化为json
	var api_response ApiResponse
	json.Unmarshal([]byte(body), &api_response)
	fmt.Println(api_response.Result.Response)

	// 返回 JSON 数据
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(api_response.Result.Response)
}

func main() {
	// 创建 CORS 中间件
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 允许所有来源，生产环境建议配置具体域名
		AllowedMethods:   []string{"POST"},
		AllowCredentials: true,
	})

	// 监听 HTTP 请求
	http.HandleFunc("/api", api)

	handler := c.Handler(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}
