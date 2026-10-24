package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC13ErrorsE2(-5); got != 10 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 10) } }
func TestHidden2(t *testing.T) { if got := TransformC13ErrorsE2(0); got != 15 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 15) } }
func TestHidden3(t *testing.T) { if got := TransformC13ErrorsE2(1000000); got != 1000015 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000015) } }
