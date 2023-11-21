package handler_test

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
)

func TestHandlerHealth(t *testing.T) {
	
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.HandlerHealth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
