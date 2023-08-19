package main

import (
	"testing"
	"time"

	"github.com/ozline/tiktok/kitex_gen/follow"
)

func testFollowerList(t *testing.T) {
	_, err := followService.FollowerList(&follow.FollowerListRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func BenchmarkFollowerList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := followService.FollowerList(&follow.FollowerListRequest{
			UserId: id,
			Token:  token,
		})

		if err != nil {
			b.Errorf("err: [%v] \n", err)
		}

		time.Sleep(100 * time.Millisecond) // Add a sleep to simulate some processing time
	}
}