# dcolor
A Go package for finding the dominant colors in images.

[![Build Status](https://travis-ci.org/drjokepu/dcolor.svg?branch=master)](https://travis-ci.org/drjokepu/dcolor)

### Usage

```go
import "github.com/drjokepu/dcolor"
colors := dcolor.Get(myImg, 3)
```