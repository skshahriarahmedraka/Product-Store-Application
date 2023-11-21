package handler_test

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
)

func TestGetProductByCategory(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductByCategory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductByMulBrand(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductByMulBrand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByMulBrand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductByName(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductByName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductByPrice(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductByPrice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductBySupplier(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductBySupplier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductBySupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductByVerifiedSupplier(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductByVerifiedSupplier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByVerifiedSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
