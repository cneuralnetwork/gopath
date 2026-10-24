package main

import "testing"

func Test1(t *testing.T) { if got := TransformP04CrawlerM2(1); got != 43 { t.Fatalf("hint: check the transformation for 1: got %d, want %d", got, 43) } }
func Test2(t *testing.T) { if got := TransformP04CrawlerM2(2); got != 44 { t.Fatalf("hint: check the transformation for 2: got %d, want %d", got, 44) } }
func Test3(t *testing.T) { if got := TransformP04CrawlerM2(-1); got != 41 { t.Fatalf("hint: check the transformation for -1: got %d, want %d", got, 41) } }
