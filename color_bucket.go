package dcolor

import (
	"image"
	"image/color"
	"math"
)

type colorBucket []color.YCbCr
type colorBucketSlice []colorBucket

const defaultThreshold = 64.0

func gatherColorBuckets(img image.Image, threshold float64) colorBucketSlice {
	bounds := img.Bounds()
	x0 := bounds.Min.X
	y0 := bounds.Min.Y
	x1 := bounds.Max.X
	y1 := bounds.Max.Y

	var buckets colorBucketSlice

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a < 0x4000 {
				continue
			}

			y, cb, cr := color.RGBToYCbCr(uint8(r/256), uint8(g/256), uint8(b/256))
			yuv := color.YCbCr{Y: y, Cb: cb, Cr: cr}

			found := false
			for idx, bucket := range buckets {
				if dist(yuv, bucket[0]) <= threshold {
					buckets[idx] = append(bucket, yuv)
					found = true
					break
				}
			}

			if !found {
				bucket := make(colorBucket, 1)
				bucket[0] = yuv
				buckets = append(buckets, bucket)
			}
		}
	}

	return buckets
}

func (s colorBucketSlice) Len() int {
	return len(s)
}

func (s colorBucketSlice) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s colorBucketSlice) Swap(i, j int) {
	b := s[i]
	s[i] = s[j]
	s[j] = b
}

func dist(c0, c1 color.YCbCr) float64 {
	dy := float64(c0.Y) - float64(c1.Y)
	dcb := float64(c0.Cb) - float64(c1.Cb)
	dcr := float64(c0.Cr) - float64(c1.Cr)

	return math.Sqrt((dy * dy) + (dcb * dcb) + (dcr * dcr))
}

func (b colorBucket) mean() color.YCbCr {
	n := uint64(len(b))

	if n == 0 {
		return color.YCbCr{Y: 0, Cb: 0, Cr: 0}
	}

	var tY, tCb, tCr uint64

	for _, c := range b {
		tY += uint64(c.Y)
		tCb += uint64(c.Cb)
		tCr += uint64(c.Cr)
	}

	return color.YCbCr{Y: uint8(tY / n), Cb: uint8(tCb / n), Cr: uint8(tCr / n)}
}
