package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		IsSubject string `json:"is_subject"`
		Item      struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL string `json:"image_url"`
		Sitelink string `json:"sitelink"`
		ID       string `json:"id"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func main() {
	translation("robot")
}

func translation(word string) {
	client := &http.Client{}

	body := DictRequest{"en2zh", word}
	buf, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	data := bytes.NewReader(buf)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")

	resp, err := client.Do(req) // 发起请求
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result DictResponse
	err = json.Unmarshal(bodyText, &result) // 反序列化记得加&, 才能成功修改结构体的值
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", result.Dictionary.Prons.En, "US:", result.Dictionary.Prons.EnUs)
	for _, item := range result.Dictionary.Explanations {
		fmt.Println(item)
	}
}
