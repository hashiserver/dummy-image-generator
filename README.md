# dummy-image-generator

## Usage

```Text
NAME:
   dummy-image-generator
USAGE:
   dummy-image-generator [options]
VERSION:
   0.0.1
OPTIONS:
   -e, --extension string   Image Extension (default "png")
   -t, --height int         Image Height
   -h, --help               help for dummy
   -w, --width int          Image Width
```

## EXAMPLE

- create png

   ```Shell
   dummy-image-generator -w 1920 -t 1080
   ```

- create jpg

   ```Shell
   dummy-image-generator -w 920 -t 380 -e jpg
   ```
## Installation

```Shell
go install github.com/hashiserver/dummy-image-generator
```

## License

MIT

## Author

hashiserver