# Hero

Hero is a command line tool for creating [Hero images](https://en.wikipedia.org/wiki/Hero_image)

## Installation

`go get github.com/sinnott74/hero`

The command `hero` will be installed at $GOPATH/bin/

## Usage

`hero --help`

```
Hero creates banner images to be displayed on a website

Usage:
  hero [flags]

Flags:
  -c, --color string        Background color (default "red")
  -y, --height int          Max y coordinate of the hero - height of the hero (default 480)
  -h, --help                help for hero
  -i, --icons stringArray   Icons
  -s, --iconsize int        Width and height of the incons on the hero (default 300)
  -o, --output string       File to output (default "./hero.png")
  -x, --width int           Max x coordinate of the hero - width of the hero (default 960)
```

## Example

`hero -c purple -i .typescript.png -i ./apache.png -o hero.png`

will create the following hero image
![Purple hero image containing Typescript & Apache icons](https://i.imgur.com/elUw241.png "Purple hero image containing Typescript & Apache icons")
