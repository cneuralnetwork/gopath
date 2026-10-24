package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC34ShippingE2(-5); got != 31 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 31) } }
func TestHidden2(t *testing.T) { if got := TransformC34ShippingE2(0); got != 36 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 36) } }
func TestHidden3(t *testing.T) { if got := TransformC34ShippingE2(1000000); got != 1000036 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000036) } }
