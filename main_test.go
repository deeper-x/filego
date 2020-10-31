package main

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestOpen(t *testing.T) {
	ctrl := gomock.NewController(t)

	rh := NewMockFileManager(ctrl)

	rh.EXPECT().WriteLine("100").Return(errors.New("demo"))
}
