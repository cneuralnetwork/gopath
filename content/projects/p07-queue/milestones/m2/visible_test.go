package main

import "testing"

func Test1(t *testing.T) { if got := TransformP07QueueM2(1); got != 73 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 73) } }
func Test2(t *testing.T) { if got := TransformP07QueueM2(2); got != 74 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 74) } }
func Test3(t *testing.T) { if got := TransformP07QueueM2(-1); got != 71 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 71) } }
