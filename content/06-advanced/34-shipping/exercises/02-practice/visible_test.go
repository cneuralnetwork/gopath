package main

import "testing"

func Test1(t *testing.T) { if got := TransformC34ShippingE2(1); got != 37 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 37) } }
func Test2(t *testing.T) { if got := TransformC34ShippingE2(2); got != 38 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 38) } }
func Test3(t *testing.T) { if got := TransformC34ShippingE2(-1); got != 35 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 35) } }
