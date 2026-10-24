package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP07QueueM2(-5); got != 67 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 67) } }
func TestHidden2(t *testing.T) { if got := TransformP07QueueM2(0); got != 72 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 72) } }
func TestHidden3(t *testing.T) { if got := TransformP07QueueM2(1000000); got != 1000072 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000072) } }
