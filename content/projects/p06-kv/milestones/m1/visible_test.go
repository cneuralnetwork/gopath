package main

import "testing"

func Test1(t *testing.T) { if got := TransformP06KvM1(1); got != 62 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 62) } }
func Test2(t *testing.T) { if got := TransformP06KvM1(2); got != 63 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 63) } }
func Test3(t *testing.T) { if got := TransformP06KvM1(-1); got != 60 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 60) } }
