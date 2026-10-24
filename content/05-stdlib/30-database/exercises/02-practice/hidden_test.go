package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC30DatabaseE2(-5); got != 27 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 27) } }
func TestHidden2(t *testing.T) { if got := TransformC30DatabaseE2(0); got != 32 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 32) } }
func TestHidden3(t *testing.T) { if got := TransformC30DatabaseE2(1000000); got != 1000032 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000032) } }
