package test

import (
	"fmt"
	"sharemusic/models/api"
	"testing"
)

func TestSongDetail(t *testing.T) {
	query := map[string]interface{}{
		"ids": "347230,347231",
	}
	fmt.Println(api.SongDetail(query))
}
