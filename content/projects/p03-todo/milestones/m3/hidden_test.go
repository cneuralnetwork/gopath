package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP03TodoM3(-5); got != 28 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 28) } }
func TestHidden2(t *testing.T) { if got := TransformP03TodoM3(0); got != 33 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 33) } }
func TestHidden3(t *testing.T) { if got := TransformP03TodoM3(1000000); got != 1000033 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000033) } }
