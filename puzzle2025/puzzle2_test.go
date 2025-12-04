package puzzle2025

import (
	"reflect"
	"testing"
)

func TestSplitString(t *testing.T) {
	result := splitString(3, "123123123")
	expected := []string{"123", "123", "123"}
	result2 := splitString(2, "123456789101")
	expected2 := []string{"12", "34", "56", "78", "91", "01"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("expected %v, got %v", expected2, result2)
	}

}
