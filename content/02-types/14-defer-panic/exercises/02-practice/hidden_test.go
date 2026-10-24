package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC14DeferPanicE2(-5); got != 11 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 11) } }
func TestHidden2(t *testing.T) { if got := TransformC14DeferPanicE2(0); got != 16 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 16) } }
func TestHidden3(t *testing.T) { if got := TransformC14DeferPanicE2(1000000); got != 1000016 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000016) } }
