package test

import "testing"

func TestNewTestEngine(t *testing.T) {
	testEngine := NewTestEngine()
	testEngine.CleanUp()
}
