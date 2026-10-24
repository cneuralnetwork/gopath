package main

import "testing"

func Test1(t *testing.T) { if got := TransformC23SyncE1(1); got != 25 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 25) } }
func Test2(t *testing.T) { if got := TransformC23SyncE1(2); got != 26 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 26) } }
func Test3(t *testing.T) { if got := TransformC23SyncE1(-1); got != 23 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 23) } }
