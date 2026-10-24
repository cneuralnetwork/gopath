package main

import "testing"

func Test1(t *testing.T) { if got := TransformP03TodoM3(1); got != 34 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 34) } }
func Test2(t *testing.T) { if got := TransformP03TodoM3(2); got != 35 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 35) } }
func Test3(t *testing.T) { if got := TransformP03TodoM3(-1); got != 32 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 32) } }
