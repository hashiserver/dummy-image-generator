# dummy-image-generator

## Usage

```Text
NAME:
   dummy-image-generator
USAGE:
   dummy-image-generator [global options]
VERSION:
   0.0.1
GLOBAL OPTIONS:
   --width, -w       width.
   --height, -t      height.
   --extension, -e   extension. default png.
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