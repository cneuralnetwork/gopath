package main

import "testing"

func Test1(t *testing.T) { if got := TransformC13ErrorsE1(1); got != 15 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 15) } }
func Test2(t *testing.T) { if got := TransformC13ErrorsE1(2); got != 16 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 16) } }
func Test3(t *testing.T) { if got := TransformC13ErrorsE1(-1); got != 13 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 13) } }
