package main

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	var scores = []int{1, 6, 3, 8, 4, 9, 0, 3}
	var resultScores = []int{0, 1, 3, 3, 4, 6, 8, 9}
	isEqual := reflect.DeepEqual(sort(scores, len(scores)), resultScores)
	if !isEqual {
		t.Errorf("recebido %v esperado %v", false, true)
	}
}
