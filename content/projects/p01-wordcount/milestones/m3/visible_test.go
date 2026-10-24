package main

import "testing"

func Test1(t *testing.T) { if got := TransformP01WordcountM3(1); got != 14 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 14) } }
func Test2(t *testing.T) { if got := TransformP01WordcountM3(2); got != 15 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 15) } }
func Test3(t *testing.T) { if got := TransformP01WordcountM3(-1); got != 12 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 12) } }
