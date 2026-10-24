package main

import "testing"

func Test1(t *testing.T) { if got := TransformC08StringsE2(1); got != 11 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 11) } }
func Test2(t *testing.T) { if got := TransformC08StringsE2(2); got != 12 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 12) } }
func Test3(t *testing.T) { if got := TransformC08StringsE2(-1); got != 9 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 9) } }
