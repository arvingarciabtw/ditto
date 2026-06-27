<div align="center">
  <img src="https://github.com/user-attachments/assets/97abc570-10c9-4f0a-bf7f-2e80b4a84d17" height="60" alt="ditto logo"/>

  <br>

[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](https://go.dev)
[![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-624fff)](https://pkg.go.dev/charm.land/bubbletea/v2)
[![Lipgloss](https://img.shields.io/badge/Lipgloss-DB7093)](https://pkg.go.dev/charm.land/lipgloss/v2)
[![evdev](https://img.shields.io/badge/evdev-000000)](https://pkg.go.dev/github.com/gvalkov/golang-evdev)

</div>

<img width="1280" height="683" alt="ditto v1.0.3 demo" src="https://github.com/user-attachments/assets/f6cda363-045b-4313-9ae2-10dfda03ced8" />

Ditto is a system-wide ASCII keyboard visualizer that captures global key presses. Even if you're outside the terminal session, the visualizer will still keep updating. Hit the repo with a star if you like this little interactive eye candy app! ⭐

## Showcase

|                                              60%                                              |                                              65%                                              |
| :-------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------: |
| ![60percent](https://github.com/user-attachments/assets/2001a816-b8d7-460b-b63f-6168cb5150e6) | ![65percent](https://github.com/user-attachments/assets/6b445af6-f482-4ecc-ac04-3e36666f194f) |

|                                              75%                                              |                                              80%                                              |
| :-------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------: |
| ![75percent](https://github.com/user-attachments/assets/738c2d9d-f4d7-4e0a-a06b-d99433945350) | ![80percent](https://github.com/user-attachments/assets/6e5cef0f-6498-4a1c-981e-15f67a3bcb06) |

|                                              96%                                              |                                              100%                                              |
| :-------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------------: |
| ![96percent](https://github.com/user-attachments/assets/6fd42732-de0e-4038-8beb-595bf3c0f771) | ![100percent](https://github.com/user-attachments/assets/4a916cb2-c433-4ac3-b9c7-05c6e666610b) |

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

If you've seen rices over at places like [r/unixp\*rn](https://www.reddit.com/r/unixporn/), every now and then you'll see some developer layouts where you'd have a code editor, then some eye candy to fill up dead space. I thought it would be nice to have some eye candy that was interactive.

And so, I thought a keyboard visualizer that updates no matter which window has focus would be pretty neat. You code away on your editor, and the ASCII keyboard on the corner lights up!

Practically, it could have a niche use as well, like sharing your screen with other people so ~~you can flex your Vim skills~~ they can see the keys you press as you navigate through your workflow.

## Installation

For now, this program is only installed through Go. So make sure you have Go installed first:

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
go install github.com/arvingarciabtw/ditto/cmd/ditto@latest

# Or clone and build
git clone https://github.com/arvingarciabtw/ditto.git
cd ditto
go build -o ditto ./cmd/ditto/
```

## Permissions

> [!IMPORTANT]
> Ditto reads raw evdev events from /dev/input/event\*, which isn't readable by normal users by default. You can grant the binary read access with:
>
> `sudo setcap cap_dac_read_search=ep "$(which ditto)"`
>
> This adds a single Linux capability (`cap_dac_read_search`) to the binary, so it only bypasses the DAC read check on `/dev/input/event*`, nothing else. The binary still runs as your user, not as root. You'd need to re-run it if you rebuilt the binary. You can revoke anytime with:
>
> `sudo setcap -r "$(which ditto)"`
>
> To my knowledge, this is the safer way of granting permissions. You could technically add the user to the input group and it would work, but this is more unsafe since this would grant full control over all devices under `/dev/input`.

## Usage

There are four commands you need to be aware of: `l`, `s`, `d`, and `c`. Pressing `l` opens up the layout list, `s` opens up the size list, and `d` opens up the standard list. If your active standard is either JIS or KS, you can press `c` to toggle between the Latin alphabet and the standard's logograms.

> [!NOTE]
> For the JIS and KS standards to render the logograms properly, you need to have a compatible font installed in your system. I recommend Noto Sans CJK JP and Noto Sans CJK KR.

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

- **map is required** — maps physical key labels → remapped labels (same format as the built-in layouts in layouts.go)
- **shift is optional** — shifted state mappings; falls back to US QWERTY shift if omitted

The layout will automatically appear in the layout list the next time you launch Ditto.

## Roadmap

Some features I'm thinking of implementing in the future, not in order.

- [x] UK layouts
- [x] ANSI and ISO standards
- [x] Additional niche standards (JIS, ABNT, KS)
- [ ] Windows support
- [ ] Mac support
- [ ] Custom layouts via TUI
- [ ] Custom finger zones

## Contributing

I'm certainly not much of an experienced programmer, so any contributions you make are **greatly appreciated**.

As I kept developing this program, I've learned that keyboards actually get pretty deep. For instance, check out this [list](https://en.wikipedia.org/wiki/List_of_QWERTY_keyboard_language_variants) of QWERTY keyboard language variants. If you'd like to add a specific layout variant to the layout list, I'd appreciate a pull request or opening up an issue about it.

A simple workflow:

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add super amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a pull request

When writing my commit messages, I follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) spec, so I'd be glad if you followed this convention as well for your contributions for the sake of consistency.

If you're liking Ditto, I'd appreciate a star! Thanks again! ⭐

## Acknowledgments

Here is a list of useful resources that I always refer to when developing:

- [effective go](https://go.dev/doc/effective_go)
- [bubble tea docs](https://pkg.go.dev/charm.land/bubbletea/v2)
- [lipgloss docs](https://pkg.go.dev/charm.land/lipgloss/v2)
- [evdev docs](https://pkg.go.dev/github.com/gvalkov/golang-evdev)
- [input interfaces](https://www.kernel.org/doc/html/latest/input/input.html)
- [setcap](https://man.archlinux.org/man/setcap.8.en)
- [standards reference](https://en.wikipedia.org/wiki/Keyboard_layout#/media/File:Physical_keyboard_layouts_comparison_ANSI_ISO_KS_ABNT_JIS.png)

## Note on AI

I'm not an experienced programmer and I've only recently picked up Go, so I'm well aware that the codebase is subpar. I still don't have an eye for what is considered idiomatic Go or not.

I've used AI to assist me with making this project, notably for pointing me towards reliable resources like documentation and for easily doing tedious, repetitive tasks (especially so for when I was adding the key matrices for each keyboard standard's size).

In the cases where I couldn't figure out an implementation despite researching on my own, I did lean into AI for suggestions, and sometimes an initial prototype of an implementation. Rest assured that I did my best in reviewing these suggestions, and making appropriate changes along the way.

Of course, I'm certain the code can still be improved a lot more, so I'd be happy to hear any feedback or suggestions. If you'd like to contribute, you can refer to the [Contributing](#contributing) section above.

## License

Distributed under the MIT License. See `LICENSE` for more information.
