# COCO (Color Converter) for Go

COCO (Color Converter) is a color conversion library for Go. Heavily inspired by NPM's [color-convert](https://www.npmjs.com/package/color-convert). It converts all ways between `rgb`, `hsl`, `hsv`, `hwb`, `cmyk`, `ansi`, `ansi16`, `hex` strings, and CSS keywords (will round to closest)

## Install

Just run this:

```bash
go get github.com/hisamafahri/coco
```

## Usage

- RGB Base

You can change an RGB base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// RGB to CMYK
coco.Rgb2Cmyk(140, 200, 100) // Output [4]float64: [30 0 50 22]

// RGB to HSL
coco.Rgb2Hsl(140, 200, 100) // Output [3]float64: [96 48 59]

// RGB to HSV
coco.Rgb2Hsv(140, 200, 100) // Output [3]float64: [96 50 78]

// RGB to HWB
coco.Rgb2Hwb(140, 200, 100) // Output [3]float64: [96 39 22]

// RGB to LAB
coco.Rgb2Lab(140, 200, 100) // Output [3]float64: [75 -37 44]

// RGB to XYZ
coco.Rgb2Hwb(140, 200, 100) // Output [3]float64: [34 48 20]

```

- HSL Base

You can change an HSL base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// HSL to RGB
coco.Hsl2Rgb(136, 54, 43) // Output [3]float64: [50 169 82]

// HSL to HSV
coco.Hsl2Hsv(136, 54, 43) // Output [3]float64: [136 70 66]

```

- HSV Base

You can change an HSV base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// HSV to RGB
coco.Hsv2Rgb(4, 78, 92) // Output [3]float64: [235 64 52]

// HSV to HSL
coco.Hsv2Hsl(4, 78, 92) // Output [3]float64: [4 82 56]

```

- HWB Base

You can change an HWB base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// HWB to RGB
coco.Hwb2Rgb(136, 0, 43) // Output [3]float64: [0 145 107]

```

- CMYK Base

You can change an CMYK base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// CMYK to RGB
coco.Cmyk2Rgb(70, 0, 51, 34) // Output [3]float64: [50 168 82]

```

- XYZ Base

You can change an XYZ base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// XYZ to RGB
coco.Xyz2Rgb(21, 18, 5) // Output [3]float64: [166 103 46]

// XYZ to LAB
coco.Xyz2Lab(21, 18, 5) // Output [3]float64: [49 20 41]

```

- LAB Base

You can change an LAB base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// LAB to XYZ
coco.Lab2Xyz(50, 21, 38) // Output [3]float64: [22 18 6]

// LAB to LCH
coco.Lab2Lch(50, 21, 38) // Output [3]float64: [50 43 61]

```

## Pending Implementation

1. RGB to keyword
2. Keyword to RGB

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)