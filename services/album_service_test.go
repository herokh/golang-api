package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/herokh/golang-api/configs"
	"github.com/herokh/golang-api/models"
	mock_repositories "github.com/herokh/golang-api/repositories/mocks"
	"github.com/herokh/golang-api/services"
)

func TestGetAlbums(t *testing.T) {
	want := []models.Album{
		models.Album{ID: 1},
		models.Album{ID: 2},
	}

	ctrl := gomock.NewController(t)
	repoMock := mock_repositories.NewMockAlbumRepository(ctrl)
	repoMock.EXPECT().GetAll().Return(want)

	service := services.NewAlbumService(&configs.Logger{},
		repoMock)

	got := service.GetAlbums()

	if len(got) != 2 {
		t.Errorf("Expected %d albums but got %d", len(want), len(got))
	}
}
