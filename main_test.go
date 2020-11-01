package filego

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestWriteLine(t *testing.T) {
	ctrl := gomock.NewController(t)

	rh := NewMockFileManager(ctrl)

	rh.EXPECT().WriteLine("100").Return(errors.New("demo"))
}

func TestDeleteLine(t *testing.T) {
	ctrl := gomock.NewController(t)

	rh := NewMockFileManager(ctrl)

	rh.EXPECT().DeleteLine("100").Return(errors.New("demo"))
}
