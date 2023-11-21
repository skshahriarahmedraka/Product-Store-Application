package handler_test

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.CreateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.DeleteProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetAllProducts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.UpdateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
