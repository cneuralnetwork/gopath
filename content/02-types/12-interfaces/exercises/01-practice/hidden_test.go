package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC12InterfacesE1(-5); got != 8 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 8) } }
func TestHidden2(t *testing.T) { if got := TransformC12InterfacesE1(0); got != 13 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 13) } }
func TestHidden3(t *testing.T) { if got := TransformC12InterfacesE1(1000000); got != 1000013 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000013) } }
