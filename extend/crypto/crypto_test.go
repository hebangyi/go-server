package sCrypto

import (
	"fmt"
	"testing"
)

func TestBase64Decode(t *testing.T) {

	fmt.Println(Base64Encode("{\"typ\":\"JWT\",\"alg\":\"HS256\"}"))

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9"

	val, err := Base64Decode(token)
	fmt.Println(val)
	fmt.Println(err)

}
