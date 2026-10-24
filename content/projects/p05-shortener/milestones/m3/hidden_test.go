package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP05ShortenerM3(-5); got != 48 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 48) } }
func TestHidden2(t *testing.T) { if got := TransformP05ShortenerM3(0); got != 53 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 53) } }
func TestHidden3(t *testing.T) { if got := TransformP05ShortenerM3(1000000); got != 1000053 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000053) } }
