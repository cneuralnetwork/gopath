package main

import "testing"

func Test1(t *testing.T) { if got := TransformP04CrawlerM3(1); got != 44 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 44) } }
func Test2(t *testing.T) { if got := TransformP04CrawlerM3(2); got != 45 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 45) } }
func Test3(t *testing.T) { if got := TransformP04CrawlerM3(-1); got != 42 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 42) } }
