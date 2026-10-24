package main

import "testing"

func Test1(t *testing.T) { if got := TransformC03BasicTypesE2(1); got != 6 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 6) } }
func Test2(t *testing.T) { if got := TransformC03BasicTypesE2(2); got != 7 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 7) } }
func Test3(t *testing.T) { if got := TransformC03BasicTypesE2(-1); got != 4 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 4) } }
