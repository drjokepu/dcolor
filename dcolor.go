// Package dcolor finds the dominant colors in images.
package dcolor

import (
	"image"
	"image/color"
	"sort"
)

// Gets the top n dominant colors in img.
func Get(img image.Image, n int) []color.Color {
	buckets := gatherColorBuckets(img, defaultThreshold)
	sort.Sort(buckets)

	if len(buckets) > n {
		buckets = buckets[:n]
	}

	colors := make([]color.Color, len(buckets))
	for idx, b := range buckets {
		colors[idx] = b.mean()
	}

	return colors
}
