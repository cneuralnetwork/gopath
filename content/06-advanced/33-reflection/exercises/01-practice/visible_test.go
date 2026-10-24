package main

import "testing"

func Test1(t *testing.T) { if got := TransformC33ReflectionE1(1); got != 35 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 35) } }
func Test2(t *testing.T) { if got := TransformC33ReflectionE1(2); got != 36 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 36) } }
func Test3(t *testing.T) { if got := TransformC33ReflectionE1(-1); got != 33 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 33) } }
