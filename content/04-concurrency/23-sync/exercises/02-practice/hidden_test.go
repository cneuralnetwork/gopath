package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC23SyncE2(-5); got != 20 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 20) } }
func TestHidden2(t *testing.T) { if got := TransformC23SyncE2(0); got != 25 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 25) } }
func TestHidden3(t *testing.T) { if got := TransformC23SyncE2(1000000); got != 1000025 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000025) } }
