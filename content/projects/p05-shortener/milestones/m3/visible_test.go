package main

import "testing"

func Test1(t *testing.T) { if got := TransformP05ShortenerM3(1); got != 54 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 54) } }
func Test2(t *testing.T) { if got := TransformP05ShortenerM3(2); got != 55 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 55) } }
func Test3(t *testing.T) { if got := TransformP05ShortenerM3(-1); got != 52 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 52) } }
