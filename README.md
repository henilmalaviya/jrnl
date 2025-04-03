# jrnl

_A simple and elegant journaling tool for the terminal._

This project was created primarily for personal use, but you're welcome to use and modify it as needed.

## Features

- Minimalist and distraction-free journaling
- Automatic record organization by date
- Markdown rendering with rich formatting
- File picker for browsing past entries
- Lightweight and fast

## Installation

Clone the repository and build it manually:

```sh
git clone https://github.com/henilmalaviya/jrnl.git
cd jrnl
go build -o jrnl
```

Then, add it to your PATH:

```sh
export PATH=$PATH:$(pwd)
```

## Usage

Start a new journal entry:

```sh
jrnl
```

Browse past entries:

```sh
jrnl -l
```

Preview an entry before saving:

```sh
jrnl -p
```

## Configuration

Entries are stored in `~/.jrnl/records`. The editor is `nano`.

## License

MIT

---

_Built with ❤️ by Henil Malaviya_
