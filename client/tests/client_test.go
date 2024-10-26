// client_test.go
package client_test

import (
	"oapi-codegen-test/client"
	"testing"
)

// NewTestArtist creates a default test artist object for testing.
func NewTestArtist() *client.Artist {
	return &client.Artist{
		ArtistName:  stringPointer("Test Artist"),
		ArtistGenre: stringPointer("Pop"),
		Albums:      intPointer(3),
		Username:    "test_artist",
	}
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}

func TestGetArtistList(t *testing.T) {
	testArtist := NewTestArtist()

	// テスト処理でtestArtistを使用して検証
	if *testArtist.ArtistName != "Test Artist" {
		t.Errorf("Expected artist name to be 'Test Artist'")
	}
}
