# COCO (Color Converter) for Go

COCO (Color Converter) is a color conversion library for Go. Heavily inspired by NPM's [color-convert](https://www.npmjs.com/package/color-convert). It converts all ways between `rgb`, `hsl`, `hsv`, `hwb`, `cmyk`, and `hex` strings.

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

// RGB to HEX
coco.Rgb2Hex(140, 200, 100) // Output string: 8CC864

// RGB to HCG
coco.Rgb2Hcg(140, 200, 100) // Output [3]float64: [96 39 65]

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

// HSL to HCG
coco.Hsl2Hsv(136, 54, 43) // Output [3]float64: [136 46 37]

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

// HSV to HCG
coco.Hsv2Hcg(4, 78, 92) // Output [3]float64: [4 72 72]

```

- HWB Base

You can change an HWB base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// HWB to RGB
coco.Hwb2Rgb(136, 0, 43) // Output [3]float64: [0 145 107]

// HWB to HCG
coco.Hwb2Hcg(136, 0, 43) // Output [3]float64: [136 57 0]

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

- LCH Base

You can change an LCH base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// LCH to LAB
coco.Lab2Xyz(30, 79, 50) // Output [3]float64: [30 51 61]

```

- HEX Base

You can change an HEX base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// HEX to RGB
coco.Hex2Rgb("4287F5") // Output [3]uint8: [66 135 245]

```

- HCG Base

You can change an HCG base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// HCG to RGB
coco.Hcg2Rgb(44, 98, 50) // Output [3]float64: [252 186 3]

// HCG to HSV
coco.Hcg2Hsv(44, 98, 50) // Output [3]float64: [44 99 99]

// HCG to HSL
coco.Hcg2Hsl(44, 98, 50) // Output [3]float64: [44 98 50]

// HCG to HWB
coco.Hcg2Hwb(44, 98, 50) // Output [3]float64: [44 1 1]

```

- Apple Base

You can change an Apple base value to other types of value. To use it:

```go

import (
	"github.com/hisamafahri/coco"
)

// Apple to RGB
coco.Apple2Rgb(44, 98, 50) // Output [3]float64: [0 1 0]

```

## Pending Implementation

1. RGB to keyword
2. Keyword to RGB
3. `Ansi16` & `Ansi256`

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)