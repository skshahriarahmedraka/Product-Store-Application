package handler_test

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
)

func TestGetAllProductStocks(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetAllProductStocks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllProductStocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductStock(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetProductStock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductStock() = %v, want %v", got, tt.want)
			}
		})
	}
}
