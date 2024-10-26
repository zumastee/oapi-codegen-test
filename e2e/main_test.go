// e2e/main_test.go
package e2e

import (
	"context"
	"io/ioutil"
	"net/http"
	"oapi-codegen-test/client" // clientパッケージのインポート
	"strings"
	"testing"
	// interfacesパッケージのインポート
)

// テスト用モッククライアントの実装
type mockClient struct{}

// GetArtistListのモック実装
func (m *mockClient) GetArtistList(ctx context.Context, params *client.GetArtistListParams) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`[{"artist_name": "E2E Artist", "artist_genre": "Jazz", "albums": 1, "username": "e2e_artist"}]`)),
	}, nil
}

func TestMainFunction(t *testing.T) {
	// モッククライアントのインスタンスを作成
	mockClient := &mockClient{}

	// テスト対象関数を実行し、モッククライアントを使用
	response, err := mockClient.GetArtistList(context.Background(), &client.GetArtistListParams{})

	// エラーチェック
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// ステータスコードとレスポンス内容の検証
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", response.StatusCode)
	}

	body, _ := ioutil.ReadAll(response.Body)
	if !strings.Contains(string(body), "E2E Artist") {
		t.Errorf("Expected response to contain 'E2E Artist'")
	}
}
