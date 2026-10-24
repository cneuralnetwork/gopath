package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC32DesignInterfacesE2(-5); got != 29 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 29) } }
func TestHidden2(t *testing.T) { if got := TransformC32DesignInterfacesE2(0); got != 34 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 34) } }
func TestHidden3(t *testing.T) { if got := TransformC32DesignInterfacesE2(1000000); got != 1000034 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000034) } }
