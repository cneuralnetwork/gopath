package main

import "testing"

func TestHidden1(t *testing.T) { if got := TransformP04CrawlerM3(-5); got != 38 { t.Fatalf("hint: check the transformation for -5: got %d, want %d", got, 38) } }
func TestHidden2(t *testing.T) { if got := TransformP04CrawlerM3(0); got != 43 { t.Fatalf("hint: check the transformation for 0: got %d, want %d", got, 43) } }
func TestHidden3(t *testing.T) { if got := TransformP04CrawlerM3(1000000); got != 1000043 { t.Fatalf("hint: check the transformation for 1000000: got %d, want %d", got, 1000043) } }
