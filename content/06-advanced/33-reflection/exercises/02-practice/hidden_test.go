package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformC33ReflectionE2(-5); got != 30 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 30) } }
func TestHidden2(t *testing.T) { if got := TransformC33ReflectionE2(0); got != 35 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 35) } }
func TestHidden3(t *testing.T) { if got := TransformC33ReflectionE2(1000000); got != 1000035 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000035) } }
