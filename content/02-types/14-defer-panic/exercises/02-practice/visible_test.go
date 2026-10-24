package main

import "testing"

func Test1(t *testing.T) { if got := TransformC14DeferPanicE2(1); got != 17 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 17) } }
func Test2(t *testing.T) { if got := TransformC14DeferPanicE2(2); got != 18 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 18) } }
func Test3(t *testing.T) { if got := TransformC14DeferPanicE2(-1); got != 15 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 15) } }
