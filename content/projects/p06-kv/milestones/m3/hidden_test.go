package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP06KvM3(-5); got != 58 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 58) } }
func TestHidden2(t *testing.T) { if got := TransformP06KvM3(0); got != 63 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 63) } }
func TestHidden3(t *testing.T) { if got := TransformP06KvM3(1000000); got != 1000063 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000063) } }
