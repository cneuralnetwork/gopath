package main

import "testing"

func Test1(t *testing.T) { if got := TransformP06KvM2(1); got != 63 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 63) } }
func Test2(t *testing.T) { if got := TransformP06KvM2(2); got != 64 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 64) } }
func Test3(t *testing.T) { if got := TransformP06KvM2(-1); got != 61 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 61) } }
