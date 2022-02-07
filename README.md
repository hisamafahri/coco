# COCO (Color Converter) for Go

COCO (Color Converter) is a color conversion library for Go. Heavily inspired by NPM's [color-convert](https://www.npmjs.com/package/color-convert). It converts all ways between `rgb`, `hsl`, `hsv`, `hwb`, `cmyk`, `ansi`, `ansi16`, `hex` strings, and CSS keywords (will round to closest)

## Install

Just run this:

```bash
go get github.com/hisamafahri/coco
```

## Usage

- RGB Base

You can change an RGB base value to other type of value. To use it:

```go
// RGB to CMYK
coco.Rgb2Cmyk(140, 200, 100) // Output [4]float64: [30 0 50 22]

// RGB to HSL
coco.Rgb2Hsl(140, 200, 100) // Output [3]float64: [96 48 59]

// RGB to HSV
coco.Rgb2Hsv(140, 200, 100) // Output [3]float64: [96 50 78]

// RGB to HWB
coco.Rgb2Hwb(140, 200, 100) // Output [3]float64: [96 39 22]

```


- *Other bases will come soon*

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)