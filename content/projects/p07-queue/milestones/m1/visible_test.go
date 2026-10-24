package main

import "testing"

func Test1(t *testing.T) { if got := TransformP07QueueM1(1); got != 72 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 72) } }
func Test2(t *testing.T) { if got := TransformP07QueueM1(2); got != 73 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 73) } }
func Test3(t *testing.T) { if got := TransformP07QueueM1(-1); got != 70 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 70) } }
