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

func TestGetValidDivisions(t *testing.T) {
	result1 := GetValidDivisions("1345")
	expected1 := []int{1, 2}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("expected %v, got %v", expected1, result1)
	}
	result2 := GetValidDivisions("123")
	expected2 := []int{1}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("expected %v, got %v", expected2, result2)
	}
	result3 := GetValidDivisions("123123123")
	expected3 := []int{1, 3}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("expected %v, got %v", expected3, result3)
	}
}
