package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP05ShortenerM1(-5); got != 46 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 46) } }
func TestHidden2(t *testing.T) { if got := TransformP05ShortenerM1(0); got != 51 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 51) } }
func TestHidden3(t *testing.T) { if got := TransformP05ShortenerM1(1000000); got != 1000051 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000051) } }
