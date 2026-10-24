package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC21ChannelsE2(-5); got != 18 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 18) } }
func TestHidden2(t *testing.T) { if got := TransformC21ChannelsE2(0); got != 23 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 23) } }
func TestHidden3(t *testing.T) { if got := TransformC21ChannelsE2(1000000); got != 1000023 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000023) } }
