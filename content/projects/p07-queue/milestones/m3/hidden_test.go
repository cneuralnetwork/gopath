package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP07QueueM3(-5); got != 68 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 68) } }
func TestHidden2(t *testing.T) { if got := TransformP07QueueM3(0); got != 73 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 73) } }
func TestHidden3(t *testing.T) { if got := TransformP07QueueM3(1000000); got != 1000073 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000073) } }
