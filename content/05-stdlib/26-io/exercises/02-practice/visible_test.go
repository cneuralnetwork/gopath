package main

import "testing"

func Test1(t *testing.T) { if got := TransformC26IoE2(1); got != 29 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 29) } }
func Test2(t *testing.T) { if got := TransformC26IoE2(2); got != 30 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 30) } }
func Test3(t *testing.T) { if got := TransformC26IoE2(-1); got != 27 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 27) } }
