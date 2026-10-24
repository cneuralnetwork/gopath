package main

import "testing"

func Test1(t *testing.T) { if got := TransformP06KvM3(1); got != 64 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 64) } }
func Test2(t *testing.T) { if got := TransformP06KvM3(2); got != 65 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 65) } }
func Test3(t *testing.T) { if got := TransformP06KvM3(-1); got != 62 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 62) } }
