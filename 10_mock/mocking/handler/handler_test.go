package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"mocking/mock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockGetUser := mock.NewMockStore(ctl)
	gomock.InOrder(
		mockGetUser.EXPECT().GetUser(1).Return("Ushakov"),
	)

	handler := NewHandler(mockGetUser)
	actual := handler.GetUserHandler(1)

	require.Equal(t, "Ushakov", actual)
}

func TestUseCase_CreateUserHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockGetUser := mock.NewMockStore(ctl)
	gomock.InOrder(
		mockGetUser.EXPECT().CreateUser("1").Return("Ushakov"),
	)

	handler := NewHandler(mockGetUser)
	actual := handler.CreateUserHandler("1")

	require.Equal(t, "Ushakov", actual)
}
