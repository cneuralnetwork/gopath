package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC21ChannelsE1(-5); got != 17 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 17) } }
func TestHidden2(t *testing.T) { if got := TransformC21ChannelsE1(0); got != 22 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 22) } }
func TestHidden3(t *testing.T) { if got := TransformC21ChannelsE1(1000000); got != 1000022 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000022) } }
