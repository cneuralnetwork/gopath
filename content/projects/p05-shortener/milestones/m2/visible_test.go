package main

import "testing"

func Test1(t *testing.T) { if got := TransformP05ShortenerM2(1); got != 53 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 53) } }
func Test2(t *testing.T) { if got := TransformP05ShortenerM2(2); got != 54 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 54) } }
func Test3(t *testing.T) { if got := TransformP05ShortenerM2(-1); got != 51 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 51) } }
