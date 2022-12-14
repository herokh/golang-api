// Code generated by MockGen. DO NOT EDIT.
// Source: repositories\album_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/herokh/golang-api/models"
)

// MockAlbumRepository is a mock of AlbumRepository interface.
type MockAlbumRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumRepositoryMockRecorder
}

// MockAlbumRepositoryMockRecorder is the mock recorder for MockAlbumRepository.
type MockAlbumRepositoryMockRecorder struct {
	mock *MockAlbumRepository
}

// NewMockAlbumRepository creates a new mock instance.
func NewMockAlbumRepository(ctrl *gomock.Controller) *MockAlbumRepository {
	mock := &MockAlbumRepository{ctrl: ctrl}
	mock.recorder = &MockAlbumRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlbumRepository) EXPECT() *MockAlbumRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockAlbumRepository) GetAll() []models.Album {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Album)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAlbumRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAlbumRepository)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockAlbumRepository) GetByID(id uint) models.Album {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.Album)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAlbumRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAlbumRepository)(nil).GetByID), id)
}

// Post mocks base method.
func (m *MockAlbumRepository) Post(album models.Album) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", album)
	ret0, _ := ret[0].(uint)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockAlbumRepositoryMockRecorder) Post(album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockAlbumRepository)(nil).Post), album)
}
