package main

import "testing"

func Test1(t *testing.T) { if got := TransformP04CrawlerM1(1); got != 42 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 42) } }
func Test2(t *testing.T) { if got := TransformP04CrawlerM1(2); got != 43 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 43) } }
func Test3(t *testing.T) { if got := TransformP04CrawlerM1(-1); got != 40 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 40) } }
