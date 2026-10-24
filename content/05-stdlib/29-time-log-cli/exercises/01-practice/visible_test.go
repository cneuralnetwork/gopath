package main

import "testing"

func Test1(t *testing.T) { if got := TransformC29TimeLogCliE1(1); got != 31 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 31) } }
func Test2(t *testing.T) { if got := TransformC29TimeLogCliE1(2); got != 32 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 32) } }
func Test3(t *testing.T) { if got := TransformC29TimeLogCliE1(-1); got != 29 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 29) } }
