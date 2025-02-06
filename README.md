# Zennハッカソン　チームDDD API

## 立ち上げ方法

```
docker-compose build
docker-compose up -d
docker-compose down
```

または

```
make up
make down
```

## WIP
一旦`localhost:8080`でsuccessが出力されるところまで


## X API認証方法

 - 開発者ポータルから，ClientID, ClientSecretを取得してメモしておく．
 - `[User authentication settings] > [GENERAL AUTHENTICATION SETTINGS]`からリダイレクト後の戻り先URL(redirect_uri)を登録しておく．（今回は適当にhttps://twitter.com/にした）

### Authorization Requestについて
Authorization Requestは以下のような形式になる

```
https://twitter.com/i/oauth2/authorize?response_type=code&client_id=<Client ID>&redirect_uri=https://127.0.0.1:3000/cb&scope=tweet.read%20users.read%20offline.access&state=abc&code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM&code_challenge_method=s256
```
各パラメータの詳細は以下の通り

 - response_type
    どのフローを利用するかを決定するパラメーター。Twitter OAuth2.0ではcode固定。
 - client_id
    Developer Portalで確認したクライアント識別子。
 - redirect_uri
    Developer Portalで登録したURL。
 - scope
    要求するアクセス範囲を明示するパラメーター。OAuth 2.0 Authorization Code Flow with PKCEのScopesに記載されているものから選択する。
 - state
    CSRF対策用のパラメーター。(今回はサンプルなので推測可能な値を設定していますが、実際に利用する際はランダムな値を指定してください。)
 - code_challenge
    code_verifierを後述のcode_challenge_methodで計算したパラメーター。なお、code_verifierは推測不可能なランダムな文字列であり、RFC 7636には43文字から128文字の範囲で指定する必要がある旨が記載されています。(今回はRFC 7636の「Appendix B. Example for the S256 code_challenge_method」に記載されている値をそのまま利用しています。)
 - code_challenge_method
    code_verifierからcode_challengeを導出する際に利用するアルゴリズム。"plain"または"s256"が指定可能。

Authorization RequestのURLにアクセスすると、指定したscopeパラメーターにもとづいて認証画面が表示される．`[Authorize app]`を押して認証する．
すると以下のような事前に設定したリダイレクトURLにリダイレクトされる．

```
https://twitter.com/?state=abc&code=<Authorization Code>
```

ここで取得される`AuthorizationCode`をメモ

そして`POST:/2/oauth2/token`エンドポイントに以下のようなリクエストを送るとAccessTokenを取得することができる

```
curl --location --request POST 'https://api.twitter.com/2/oauth2/token' \
                  --basic -u '<Client ID>:<Client Secret>' \
                  --header 'Content-Type: application/x-www-form-urlencoded' \
                  --data-urlencode 'code=<Authorization Code>' \
                  --data-urlencode 'grant_type=authorization_code' \
                  --data-urlencode 'client_id=<Client ID>' \
                  --data-urlencode 'redirect_uri=https://twitter.com/' \
                  --data-urlencode 'code_verifier=dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk'
```

ここで取得したAccessTokenをheaderに貼り付けると，以下のように自分のユーザーIDが取得でき，そのユーザーIDを用いると，自分のアカウントのポストを一覧で取得することができる．

```
curl --location --request GET 'https://api.X.com/2/users/me' \
--header 'Authorization: Bearer <accessToken>'
```

```
{"data":{"id":"1486606447220002818","name":"えふじ","username":"tsunufu_f"}}
```


```
curl --location 'https://api.twitter.com/2/users/1486606447220002818/tweets' \
--header 'Authorization: Bearer <accessToken>'
```


```
{"data":[{"id":"1885152132191707483","edit_history_tweet_ids":["1885152132191707483"],"text":"https://t.co/belqCsVFD8"},{"id":"1885152074381615146","edit_history_tweet_ids":["1885152074381615146"],"text":"またもや中華AI\nアリババ発「Qwen」\n使ってみる\nhttps://t.co/BMi7DiOTqb"},{"id":"1883786132779147636","edit_history_tweet_ids":["1883786132779147636"],"text":"RT @CyberAgent_PR: 【モデル公開のお知らせ】\nDeepSeek-R1-Distill-Qwen-14B/32Bをベースに日本語データで追加学習を行ったLLMを公開いたしました。今後もモデル公開や産学連携を通じて国内の自然言語処理技術の発展に貢献してまいります。…"},{"id":"1883180865490211238","edit_history_tweet_ids":["1883180865490211238"],"text":"テキストメッセージを通話のようにやりとりができるSNS\nお互いが接続している時にやり取りできるので既読の概念がない\nおもしろい https://t.co/m6BVLKDJ9z"},{"id":"1882433929946517546","edit_history_tweet_ids":["1882433929946517546"],"text":"Think in English output in Japanese つけないと精度結構変わる気がする"},{"id":"1882384439931723970","edit_history_tweet_ids":["1882384439931723970"],"text":"DeepSeek V3と比較するとR1は速度もすごい早くなってる"},{"id":"1881696740493128016","edit_history_tweet_ids":["1881696740493128016"],"text":"RT @ytiskw: 『中国のサム・アルトマン』とも呼ばれるDeepSeek創業者、梁文峰（Liang Wenfeng）は何者か？\n\n今、AI業界を語るうえで外せない存在として注目を集めているのが、スタートアップ「DeepSeek（深度求索）」の創業者、梁文峰（Liang W…"},{"id":"1881330888073699686","edit_history_tweet_ids":["1881330888073699686"],"text":"マネタイズをちゃんと考えたい"},{"id":"1881266869040255101","edit_history_tweet_ids":["1881266869040255101"],"text":"https://t.co/nIoffTHNYu"},{"id":"1880457854676791699","edit_history_tweet_ids":["1880457854676791699"],"text":"SWEAIエージェント\n\nhttps://t.co/bOp5jfLiRE"}],"meta":{"next_token":"7140dibdnow9c7btw4b38ggdejbbk1sy0255d6pmwck6x","result_count":10,"newest_id":"1885152132191707483","oldest_id":"1880457854676791699"}}
```

