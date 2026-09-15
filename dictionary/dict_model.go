package main

// DictRequest 表示向 API 发送的字典请求结构。
type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserId    string `json:"user_id"`
}

// DictResponse 表示从 API 返回的字典响应结构。
type DictResponse struct {
	Rc int `json:"rc"` // 响应代码

	Wiki struct {
	} `json:"wiki"` // 维基信息

	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"` // 美式发音
			En   string `json:"en"`    // 英式发音
		} `json:"prons"` // 发音
		Explanations []string      `json:"explanations"` // 单词解释
		Synonym      []string      `json:"synonym"`      // 同义词
		Antonym      []string      `json:"antonym"`      // 反义词
		WqxExample   [][]string    `json:"wqx_example"`  // 例句
		Entry        string        `json:"entry"`        // 单词条目
		Type         string        `json:"type"`         // 单词类型
		Related      []interface{} `json:"related"`      // 相关词
		Source       string        `json:"source"`       // 单词来源
	} `json:"dictionary"` // 字典信息
}
