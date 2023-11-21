package handler

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestCreateSupplier(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSupplier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSupplier(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSupplier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllSuppliers(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllSuppliers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllSuppliers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSupplier(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSupplier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateSupplier(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateSupplier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
