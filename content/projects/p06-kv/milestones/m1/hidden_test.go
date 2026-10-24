package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP06KvM1(-5); got != 56 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 56) } }
func TestHidden2(t *testing.T) { if got := TransformP06KvM1(0); got != 61 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 61) } }
func TestHidden3(t *testing.T) { if got := TransformP06KvM1(1000000); got != 1000061 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000061) } }
