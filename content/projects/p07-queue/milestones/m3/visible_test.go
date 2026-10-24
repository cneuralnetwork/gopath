package main

import "testing"

func Test1(t *testing.T) { if got := TransformP07QueueM3(1); got != 74 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 74) } }
func Test2(t *testing.T) { if got := TransformP07QueueM3(2); got != 75 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 75) } }
func Test3(t *testing.T) { if got := TransformP07QueueM3(-1); got != 72 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 72) } }
