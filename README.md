<div align="center">
  <img src="https://github.com/user-attachments/assets/97abc570-10c9-4f0a-bf7f-2e80b4a84d17" height="60" alt="ditto logo"/>

  <br>

[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](https://go.dev)
[![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-624fff)](https://pkg.go.dev/charm.land/bubbletea/v2)
[![Lipgloss](https://img.shields.io/badge/Lipgloss-DB7093)](https://pkg.go.dev/charm.land/lipgloss/v2)
[![evdev](https://img.shields.io/badge/evdev-000000)](https://pkg.go.dev/github.com/gvalkov/golang-evdev)

</div>

<img width="1280" height="685" alt="ditto demo" src="https://github.com/user-attachments/assets/bf71030a-eacd-4eaa-9be5-06adb86de192" />

Ditto is a system-wide ASCII keyboard visualizer that captures global key presses. Even if you're outside the terminal session, the visualizer will still keep updating. Hit the repo with a star if you like this little interactive eye candy app! ⭐

## Table of Contents

- [Motivation](#motivation)
- [Installation](#installation)
- [Permissions](#permissions)
- [Usage](#usage)
  - [Lists](#lists)
  - [Custom Layouts](#custom-layouts)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [Acknowledgments](#acknowledgments)
- [Note on AI](#note-on-ai)
- [License](#license)

## Motivation

If you've seen rices over at places like [r/unixp\*rn](https://www.reddit.com/r/unixporn/), every now and then you'll see some developer layouts where you'd have a code editor, then some eye candy to fill up dead space. I thought it would be nice to have an interactive keyboard visualizer that updates no matter which window has focus.

Practically, it could have a niche use as well, like sharing your screen with other people so ~~you can flex your Vim skills~~ they can see the keys you press as you navigate through your workflow.

## Installation

I recommend installing this program through Go. So make sure you have Go installed first:

```bash
# Check if Go is installed
go version

# If not, install it from https://go.dev/dl/
# or via your package manager, e.g.:

# sudo pacman -S go            (Arch)
# sudo apt install golang-go   (Debian/Ubuntu)
```

Once you have Go installed, you can install it directly, or clone if you'd like to mess around with it:

```bash

# Install directly
go install github.com/arvingarciabtw/ditto@latest

# Or clone and build
git clone https://github.com/arvingarciabtw/ditto.git
cd ditto
go build
```

You can also install it from the AUR, though note that this one isn't maintained by me. You can see the [package](https://aur.archlinux.org/packages/ditto) here. For now, I recommend installing it with Go.

## Permissions

> [!IMPORTANT]
> Ditto reads raw evdev events from /dev/input/event\*, which isn't readable by normal users by default. Grant the binary read access with:
>
> `sudo setcap cap_dac_read_search=ep "$(which ditto)"`
>
> This adds a single Linux capability (`cap_dac_read_search`) to the binary — it only bypasses the DAC read check on `/dev/input/event*`, nothing else. The binary still runs as your user, not as root. Re-run after rebuilding. Revoke anytime with:
>
> `sudo setcap -r "$(which ditto)"`

## Usage

There are three commands you need to be aware of: `l`, `s`, and `d`. Pressing `l` opens up the layout list, `s` opens up the size list, and `d` alternates between the ANSI standard and ISO standard.

### Lists

| Size   | Form Factor                        |
| ------ | ---------------------------------- |
| `60%`  | Compact, no F-row or arrows        |
| `65%`  | 60% + arrow keys + nav cluster     |
| `75%`  | Has F-row, compact layout          |
| `80%`  | TKL — F-row, arrows, no numpad     |
| `96%`  | Compact full-size, includes numpad |
| `100%` | Full-size with everything          |

| Layout       | Description                |
| ------------ | -------------------------- |
| `qwerty`     | Standard US layout         |
| `qwerty uk`  | Standard UK layout         |
| `dvorak`     | Dvorak simplified          |
| `dvorak uk`  | Dvorak UK layout           |
| `colemak`    | Colemak modern alternative |
| `colemak-dh` | Colemak with angle mod     |
| `workman`    | Workman layout             |
| `azerty`     | French AZERTY              |

If you wish, you can add your own key maps for your custom layouts!

### Custom Layouts

Custom layouts are loaded from JSON files placed in `~/.config/ditto/layouts/`.

Each `.json` file becomes a named layout (the filename without extension). Format:

```jsonc
{
  "map": {
    "A": "A",
    "S": "O",
    "D": "E",
    "F": "U",
    "G": "I",
    "H": "D",
    "J": "H",
    "K": "T",
    "L": "N",
    "Q": "'",
    "W": ",",
    "E": ".",
    // ...
  },
  "shift": {
    "1": "!",
    "2": "@", // ...
  },
}
```

- map required — maps physical key labels → remapped labels (same format as the built-in layouts in layouts.go)
- shift optional — shifted state mappings; falls back to US QWERTY shift if omitted

The layout will automatically appear in the layout list the next time you launch Ditto.

## Roadmap

Some features I'm thinking of implementing in the future, not in order.

- [x] UK layouts
- [x] ANSI and ISO standards
- [ ] Windows support
- [ ] Mac support
- [ ] Custom layouts via TUI
- [ ] Custom finger zones
- [ ] Additional niche standards (JIS, ABNT, KS)

## Contributing

I'm certainly not much of an experienced programmer, so any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". If you're liking Ditto, I'd appreciate a star! Thanks again! ⭐

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add super amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a pull request

When writing my commit messages, I follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) spec, so I'd be glad if you followed this convention as well for your contributions for the sake of consistency.

## Acknowledgments

Here is a list of useful resources that I always refer to when developing:

- [effective go](https://go.dev/doc/effective_go)
- [bubble tea docs](https://pkg.go.dev/charm.land/bubbletea/v2)
- [lipgloss docs](https://pkg.go.dev/charm.land/lipgloss/v2)
- [evdev docs](https://pkg.go.dev/github.com/gvalkov/golang-evdev)
- [input interfaces](https://www.kernel.org/doc/html/latest/input/input.html)
- [setcap](https://man.archlinux.org/man/setcap.8.en)
- [capabilities](https://man7.org/linux/man-pages/man7/capabilities.7.html)

## Note on AI

I'm only a beginner programmer and I've only recently picked up Go, so I'm well aware that the codebase is subpar. I still don't have an eye for what is considered idiomatic Go or not. I've used AI to assist me with making this project, especially in cases where I got really stuck, but rest assured I also tried my best in reviewing it.

I'm always eager to learn and improve, so I'd be happy to hear any feedback or suggestions. If you'd like to contribute, you can refer to the [Contributing](#contributing) section above.

## License

Distributed under the MIT License. See `LICENSE` for more information.
