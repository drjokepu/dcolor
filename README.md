# dcolor
A Go package for finding the dominant colors in images.

[![GoDoc](https://godoc.org/github.com/drjokepu/dcolor?status.svg)](https://godoc.org/github.com/drjokepu/dcolor) [![Build Status](https://travis-ci.org/drjokepu/dcolor.svg?branch=master)](https://travis-ci.org/drjokepu/dcolor)

Based on the algorithm discussed at [http://stackoverflow.com/a/13675803/8954](http://stackoverflow.com/a/13675803/8954).

### Usage

```go
import "github.com/drjokepu/dcolor"
colors := dcolor.Get(myImg, 3)
```