package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP04CrawlerM1(-5); got != 36 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 36) } }
func TestHidden2(t *testing.T) { if got := TransformP04CrawlerM1(0); got != 41 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 41) } }
func TestHidden3(t *testing.T) { if got := TransformP04CrawlerM1(1000000); got != 1000041 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000041) } }
