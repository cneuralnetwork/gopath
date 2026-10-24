package main

import "testing"

func Test1(t *testing.T) { if got := TransformC25ContextE2(1); got != 28 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 28) } }
func Test2(t *testing.T) { if got := TransformC25ContextE2(2); got != 29 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 29) } }
func Test3(t *testing.T) { if got := TransformC25ContextE2(-1); got != 26 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 26) } }
