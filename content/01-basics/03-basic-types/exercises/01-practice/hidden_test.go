package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC03BasicTypesE1(-5); got != -1 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, -1) } }
func TestHidden2(t *testing.T) { if got := TransformC03BasicTypesE1(0); got != 4 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 4) } }
func TestHidden3(t *testing.T) { if got := TransformC03BasicTypesE1(1000000); got != 1000004 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000004) } }
