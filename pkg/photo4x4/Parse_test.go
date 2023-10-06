package photo4x4

import (
	"fmt"
	"testing"
)

func TestParseCategory(t *testing.T) {
	links, err := ParseCategory("https://4x4photo.ru/foto-kartinki/", 2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(links))
}
