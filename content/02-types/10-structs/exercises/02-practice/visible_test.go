package main

import "testing"

func Test1(t *testing.T) { if got := TransformC10StructsE2(1); got != 13 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 13) } }
func Test2(t *testing.T) { if got := TransformC10StructsE2(2); got != 14 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 14) } }
func Test3(t *testing.T) { if got := TransformC10StructsE2(-1); got != 11 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 11) } }
