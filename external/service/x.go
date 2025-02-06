package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/DDD-Zenn/api/external/serviceIF"
)

type xsvc struct {
}

func NewXService() serviceIF.X {
	return &xsvc{}
}

type TweetResponse struct {
	Data []struct {
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
	} `json:"meta"`
}

func (svc *xsvc) GetMyPost() string {
	//最新の一件を取得したいときは/recentEPを叩く
	url := "https://api.x.com/2/tweets/search/recent?query=from:tsunufu_f"
	//10件まとめて取得したい時は/:id/tweetsを叩く
	//ここでは例として@tsunufu_fのツイートを取得
	//今後はユーザーidも動的にいじりたい
	// id := "1486606447220002818"
	// url := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", id)

	token := os.Getenv("X_BEARER_TOKEN")
	if token == "" {
		log.Fatal("X_BEARER_TOKEN is not set")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("リクエスト作成エラー: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Cookie", `guest_id=v1%3A173875786411475703; guest_id_ads=v1%3A173875786411475703; guest_id_marketing=v1%3A173875786411475703; personalization_id="v1_Yyoic55er/VnGt/CDMj4XQ=="'`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("リクエスト送信エラー: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("レスポンス読み込みエラー: %v", err)
	}

	var tweetResponse TweetResponse
	if err := json.Unmarshal(body, &tweetResponse); err != nil {
		log.Fatalf("JSONパースエラー: %v", err)
	}

	var texts []string
	for _, tweet := range tweetResponse.Data {
		texts = append(texts, tweet.Text)
	}

	result := strings.Join(texts, "\n")
	return result
}
