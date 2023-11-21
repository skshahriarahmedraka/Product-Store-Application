package handler_test

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
)

func TestCreateBrand(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.CreateBrand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBrand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteBrand(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.DeleteBrand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteBrand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllBrands(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetAllBrands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllBrands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBrand(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetBrand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBrand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateBrand(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.UpdateBrand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateBrand() = %v, want %v", got, tt.want)
			}
		})
	}
}
