package main

import "testing"

func Test1(t *testing.T) { if got := TransformC17TestingE2(1); got != 20 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 20) } }
func Test2(t *testing.T) { if got := TransformC17TestingE2(2); got != 21 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 21) } }
func Test3(t *testing.T) { if got := TransformC17TestingE2(-1); got != 18 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 18) } }
