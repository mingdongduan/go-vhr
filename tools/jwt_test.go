package tools

import (
	"fmt"
	"testing"
	"time"
)

const s = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MTY3MTY0MDA2NjAxMCwidXNlcm5hbWUiOiJhZG1pbiJ9.g9cEIIcld8YOhak01bQIIzcg_o5kltk-SCrmowDZbXc"


func TestParseToken(t *testing.T) {
	token := ParseToken(s)
	fmt.Println(token)
	fmt.Println(token["expire_time"])
	fmt.Println(token["username"])

	fmt.Println("----")
	milli := time.Now().UnixMilli()
	print(milli)
}
