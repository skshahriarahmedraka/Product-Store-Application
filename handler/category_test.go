package handler_test

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
)

func TestCreateCategory(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.CreateCategory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.DeleteCategory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCategories(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetAllCategories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCategory(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetCategory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.UpdateCategory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
