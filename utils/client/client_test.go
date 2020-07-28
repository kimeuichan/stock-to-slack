package client

import (
	"reflect"
	"testing"
)

func TestNaverGetClient(t *testing.T) {
	naverClient := GetClient("naver")

	if reflect.TypeOf(naverClient) != reflect.TypeOf(&NaverClient{}) {
		t.Error("this client is not naver client")
	}
}
