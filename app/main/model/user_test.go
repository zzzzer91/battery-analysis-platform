package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUser_JsonTag(t *testing.T) {
	user := &User{}
	b, _ := json.Marshal(user)
	fmt.Printf("%s\n", b)
}
