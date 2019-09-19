package main

import (
	"context"
	"testing"

	pb "github.com/q3k/is-odd/proto"
)

func TestIsOdd(t *testing.T) {
	ctx := context.Background()
	s := &server{}

	tests := []struct {
		number int64
		want   pb.IsOddResponse_Result
	}{
		{1337, pb.IsOddResponse_RESULT_ODD},
		{-1337, pb.IsOddResponse_RESULT_ODD},
		{42, pb.IsOddResponse_RESULT_NON_ODD},
		{-42, pb.IsOddResponse_RESULT_NON_ODD},
		{0, pb.IsOddResponse_RESULT_NON_ODD},
	}

	for i, test := range tests {
		req := &pb.IsOddRequest{
			Number: test.number,
		}
		res, err := s.IsOdd(ctx, req)
		if err != nil {
			t.Fatalf("Case %d: error from service: %v", i, err)
		}

		if want, got := test.want, res.Result; want != got {
			t.Errorf("Case %d: want %v, got %v", i, want, got)
		}
	}
}
