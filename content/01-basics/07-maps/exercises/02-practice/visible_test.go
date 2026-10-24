package main

import "testing"

func Test1(t *testing.T) { if got := TransformC07MapsE2(1); got != 10 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 10) } }
func Test2(t *testing.T) { if got := TransformC07MapsE2(2); got != 11 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 11) } }
func Test3(t *testing.T) { if got := TransformC07MapsE2(-1); got != 8 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 8) } }
