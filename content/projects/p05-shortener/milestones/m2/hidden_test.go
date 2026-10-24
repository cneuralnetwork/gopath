package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP05ShortenerM2(-5); got != 47 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 47) } }
func TestHidden2(t *testing.T) { if got := TransformP05ShortenerM2(0); got != 52 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 52) } }
func TestHidden3(t *testing.T) { if got := TransformP05ShortenerM2(1000000); got != 1000052 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000052) } }
