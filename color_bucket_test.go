package dcolor

import (
	"image"
	"image/color"
	"testing"
)

func TestDistZero(t *testing.T) {
	c0 := color.YCbCr{Y: 0, Cb: 0, Cr: 0}
	c1 := color.YCbCr{Y: 0, Cb: 0, Cr: 0}

	d := dist(c0, c1)
	assertFloat64(t, 0.0, d)
}

func TestDistNonZero(t *testing.T) {
	c0 := color.YCbCr{Y: 1, Cb: 1, Cr: 1}
	c1 := color.YCbCr{Y: 3, Cb: 4, Cr: 7}

	d := dist(c0, c1)
	assertFloat64(t, 7.0, d)
}

func TestMeanZero(t *testing.T) {
	b := colorBucket{
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
		color.YCbCr{Y: 0, Cb: 0, Cr: 0},
	}

	m := b.mean()
	assertYUVEqual(t, m, color.YCbCr{Y: 0, Cb: 0, Cr: 0})
}

func TestMeanConstant(t *testing.T) {
	b := colorBucket{
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
		color.YCbCr{Y: 100, Cb: 40, Cr: 220},
	}

	m := b.mean()
	assertYUVEqual(t, m, color.YCbCr{Y: 100, Cb: 40, Cr: 220})
}

func TestMeanVarying(t *testing.T) {
	b := colorBucket{
		color.YCbCr{Y: 10, Cb: 90, Cr: 30},
		color.YCbCr{Y: 20, Cb: 70, Cr: 60},
		color.YCbCr{Y: 30, Cb: 20, Cr: 95},
		color.YCbCr{Y: 40, Cb: 160, Cr: 95},
		color.YCbCr{Y: 50, Cb: 200, Cr: 20},
	}

	m := b.mean()
	assertYUVEqual(t, m, color.YCbCr{Y: 30, Cb: 108, Cr: 60})
}

func TestGatherColorBucketsAllBlack(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for i := 0; i < len(img.Pix); i++ {
		if i%4 < 3 {
			img.Pix[i] = 0
		} else {
			img.Pix[i] = 255
		}
	}

	buckets := gatherColorBuckets(img, defaultThreshold)
	assertInt(t, 1, len(buckets))
	assertYUVEqual(t, rgb2yuv(0, 0, 0), buckets[0].mean())
}

func TestGatherColorBuckets2(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	l0 := (len(img.Pix) / 4) * 3
	l1 := (len(img.Pix) / 4) * 1

	for i := 0; i < l0; i++ {
		switch i % 4 {
		case 0:
			img.Pix[i] = 200
			break
		case 1:
			img.Pix[i] = 10
			break
		case 2:
			img.Pix[i] = 30
			break
		case 3:
			img.Pix[i] = 255
			break
		}
	}

	for i := l0; i < l0+l1; i++ {
		switch i % 4 {
		case 0:
			img.Pix[i] = 40
			break
		case 1:
			img.Pix[i] = 10
			break
		case 2:
			img.Pix[i] = 230
			break
		case 3:
			img.Pix[i] = 255
			break
		}
	}

	assertInt(t, 200, int(img.Pix[0]))
	assertInt(t, 10, int(img.Pix[1]))
	assertInt(t, 30, int(img.Pix[2]))
	assertInt(t, 255, int(img.Pix[3]))

	buckets := gatherColorBuckets(img, defaultThreshold)
	assertInt(t, 2, len(buckets))
}

func assertYUVEqual(t *testing.T, c0, c1 color.YCbCr) {
	if c0.Y != c1.Y || c0.Cb != c1.Cb || c0.Cr != c1.Cr {
		t.Errorf("Expected: %v, actual: %v", c0, c1)
		t.FailNow()
	}
}

func assertInt(t *testing.T, v0, v1 int) {
	if v0 != v1 {
		t.Errorf("Expected: %d, actual: %d", v0, v1)
		t.FailNow()
	}
}

func assertFloat64(t *testing.T, v0, v1 float64) {
	if v0 != v1 {
		t.Errorf("Expected: %f, actual: %f", v0, v1)
		t.FailNow()
	}
}

func rgb2yuv(r, g, b uint8) color.YCbCr {
	y, cb, cr := color.RGBToYCbCr(r, g, b)
	return color.YCbCr{y, cb, cr}
}
