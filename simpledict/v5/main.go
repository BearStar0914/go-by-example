package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DictRequest struct {
	Text string `json:"text"`
	Language    string `json:"language"`
	// UserID    string `json:"user_id"`
}

type DictResponse struct {
	Words []struct {
		Source int `json:"source"`
		Text string `json:"text"`
		PosList []struct {
			Type int `json:"type"`
			Phonetics []struct {
				Type int `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text string `json:"text"`
				Examples []struct {
					Type int `json:"type"`
					Sentences []struct {
						Text string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode int `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{Text: word, Language: "en"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzKwGLQDGX3WfISWss3e9WX7J6&_signature=_02B4Z6wo00001TpTTtQAAIDBulG0l0ZN26k6Q0pAACzlpAVtKndaoOG-2ZhaYYWAJjHIKtj07Atdc3KSaeXJtFV1zQAie7xkK.5UAbiH8.wixCVWEZdd0CwIW0jYO8B0OZNoKYp51FSJQLTd1e", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16519362511056980; i18next=zh-CN; ttcid=ad16d065e3204a5aad6852c073aa63f519; tt_scid=3Sd5MG2R0tCldnnfNMup8fPsMrpn5R0yjJC5QrU0DRMk5.UN8I8QO3YvkMulhBMR6860; s_v_web_id=verify_866e748b9ff549bdc33c4fb1a7dda5b5; _tea_utm_cache_2018=undefined")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=detect&target_language=zh&text=good%0A%0A")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Mobile Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	// fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	for _, word := range dictResponse.Words{
		for _, item := range word.PosList{
			for _, explantion := range item.Explanations{
				fmt.Println(explantion.Text)
			}
		}
	}
	// fmt.Printf("%#v\n", dictResponse) 
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
