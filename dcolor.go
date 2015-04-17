// Package dcolor finds the dominant colors in images. Based on the algorithm discussed at http://stackoverflow.com/a/13675803/8954.
package dcolor

import (
	"image"
	"image/color"
	"sort"
)

// Gets the top n dominant colors in img. The number of colors returned may be less than the number of colors requested.
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
