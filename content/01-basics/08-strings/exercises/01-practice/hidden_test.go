package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC08StringsE1(-5); got != 4 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 4) } }
func TestHidden2(t *testing.T) { if got := TransformC08StringsE1(0); got != 9 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 9) } }
func TestHidden3(t *testing.T) { if got := TransformC08StringsE1(1000000); got != 1000009 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000009) } }
