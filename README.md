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
import (
    "fmt"
    "github.com/hisamafahri/coco"
)

func main() {
    result := coco.Rgb2Hsl(140, 200, 100) // [96 48 59]
}

```


- *Other bases will come soon*

## Author

[Hisam A Fahri](https://hisamafahri.com): [@hisamafahri](https://github.com/hisamafahri)

## License

[MIT](LICENSE)