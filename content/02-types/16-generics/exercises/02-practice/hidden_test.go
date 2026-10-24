package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC16GenericsE2(-5); got != 13 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 13) } }
func TestHidden2(t *testing.T) { if got := TransformC16GenericsE2(0); got != 18 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 18) } }
func TestHidden3(t *testing.T) { if got := TransformC16GenericsE2(1000000); got != 1000018 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000018) } }
