package main

import "testing"

func Test1(t *testing.T) { if got := TransformC33ReflectionE2(1); got != 36 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 36) } }
func Test2(t *testing.T) { if got := TransformC33ReflectionE2(2); got != 37 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 37) } }
func Test3(t *testing.T) { if got := TransformC33ReflectionE2(-1); got != 34 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 34) } }
