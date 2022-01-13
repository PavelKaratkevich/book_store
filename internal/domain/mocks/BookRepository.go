// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "book_store/internal/domain"

	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// DeleteBook provides a mock function with given fields: ctx, id
func (_m *BookRepository) DeleteBook(ctx *gin.Context, id int) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(*gin.Context, int) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBook provides a mock function with given fields: ctx, id
func (_m *BookRepository) GetBook(ctx *gin.Context, id int) (*domain.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Book
	if rf, ok := ret.Get(0).(func(*gin.Context, int) *domain.Book); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields: ctx
func (_m *BookRepository) GetBooks(ctx *gin.Context) ([]domain.Book, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Book
	if rf, ok := ret.Get(0).(func(*gin.Context) []domain.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBook provides a mock function with given fields: ctx, req
func (_m *BookRepository) NewBook(ctx *gin.Context, req domain.Book) (int, error) {
	ret := _m.Called(ctx, req)

	var r0 int
	if rf, ok := ret.Get(0).(func(*gin.Context, domain.Book) int); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context, domain.Book) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: ctx, req
func (_m *BookRepository) UpdateBook(ctx *gin.Context, req domain.Book) (int, error) {
	ret := _m.Called(ctx, req)

	var r0 int
	if rf, ok := ret.Get(0).(func(*gin.Context, domain.Book) int); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context, domain.Book) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}