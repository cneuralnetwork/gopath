package main

import "testing"

func Test1(t *testing.T) { if got := TransformP05ShortenerM1(1); got != 52 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 52) } }
func Test2(t *testing.T) { if got := TransformP05ShortenerM1(2); got != 53 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 53) } }
func Test3(t *testing.T) { if got := TransformP05ShortenerM1(-1); got != 50 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 50) } }
