package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP06KvM2(-5); got != 57 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 57) } }
func TestHidden2(t *testing.T) { if got := TransformP06KvM2(0); got != 62 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 62) } }
func TestHidden3(t *testing.T) { if got := TransformP06KvM2(1000000); got != 1000062 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000062) } }
