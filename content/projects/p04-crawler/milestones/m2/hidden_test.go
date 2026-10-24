package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP04CrawlerM2(-5); got != 37 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 37) } }
func TestHidden2(t *testing.T) { if got := TransformP04CrawlerM2(0); got != 42 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 42) } }
func TestHidden3(t *testing.T) { if got := TransformP04CrawlerM2(1000000); got != 1000042 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000042) } }
