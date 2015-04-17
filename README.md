# dcolor
A Go package for finding the dominant colors in images.

[![Build Status](https://travis-ci.org/drjokepu/dcolor.svg?branch=master)](https://travis-ci.org/drjokepu/dcolor)

Based on the algorithm discussed at [http://stackoverflow.com/a/13675803/8954](http://stackoverflow.com/a/13675803/8954).

**GoDoc**: [https://godoc.org/github.com/drjokepu/dcolor](https://godoc.org/github.com/drjokepu/dcolor)

### Usage

```go
import "github.com/drjokepu/dcolor"
colors := dcolor.Get(myImg, 3)
```