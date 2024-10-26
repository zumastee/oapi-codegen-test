package main

import (
	"context"
	"log"
	"net/http"

	// APIクライアントのパッケージをインポート
	// このパッケージはoapi-codegenで生成されたもの
	// モジュール＋パッケージ名で指定
	"oapi-codegen-test/client"
)

func main() {
	// クライアントを作成
	appClient, err := client.NewClientWithResponses("https://api.example.com", client.WithHTTPClient(&http.Client{}))

	// クエリパラメータの設定
	params := &client.GetArtistListParams{
		Limit:  intPointer(10),
		Offset: intPointer(0),
	}

	// APIリクエストの送信
	response, err := appClient.GetArtistListWithResponse(context.Background(), params)
	if err != nil {
		log.Fatalf("Error calling API: %v", err)
	}

	log.Printf("Response: %v", response.JSON200) // 成功時のレスポンスを表示
}

// intをポインタで渡すためのヘルパー関数
func intPointer(i int) *int {
	return &i
}
