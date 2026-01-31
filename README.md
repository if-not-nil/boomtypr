# boomtypr

A minimal terminal-based typing test built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Features

- Clean, distraction-free TUI
- Multiple test modes: Time, Words, and Zen
- Configurable test duration and word count
- Results screen with WPM and accuracy display
- Responsive text wrapping

## Install

```bash
go install github.com/yagnikpt/boomtypr@latest
```

Or build from source:

```bash
git clone https://github.com/yagnikpt/boomtypr.git
cd boomtypr
make
```

## Usage

```bash
boomtypr
```

## Roadmap

- [x] WPM and accuracy stats display
- [x] Results screen after test completion
- [x] Configurable word count and test duration
- [ ] Multiple word lists
- [x] Test restart functionality
- [ ] Word erase (ctrl+w / ctlr+backspace)
- [ ] Realtime stats display / better stats

## Inspiration
The UI and flow is inspired by [ashish0kumar/typtea](https://github.com/ashish0kumar/typtea)

## License

MIT
