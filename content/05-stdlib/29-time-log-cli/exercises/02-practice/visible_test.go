package main

import "testing"

func Test1(t *testing.T) { if got := TransformC29TimeLogCliE2(1); got != 32 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 32) } }
func Test2(t *testing.T) { if got := TransformC29TimeLogCliE2(2); got != 33 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 33) } }
func Test3(t *testing.T) { if got := TransformC29TimeLogCliE2(-1); got != 30 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 30) } }
