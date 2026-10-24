package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP07QueueM1(-5); got != 66 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 66) } }
func TestHidden2(t *testing.T) { if got := TransformP07QueueM1(0); got != 71 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 71) } }
func TestHidden3(t *testing.T) { if got := TransformP07QueueM1(1000000); got != 1000071 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000071) } }
