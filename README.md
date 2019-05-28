# ESClang v0.1.0

ESClang is an esoteric programming language where instructions are represented with [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code).

## Example

in [examples](/examples) directory

## Usage

If you have golang:
```sh
go get github.com/nakario/esclang/cmd/esc
esc your_script.esc
```

or if you have docker:
```sh
git clone https://github.com/nakario/esclang.git
cd esclang
docker build -t esc
docker run -it -v /path/to/dir:/code:ro esc /code/your_script.esc
```

## Syntax

ESClang v0.1 uses SGR - SELECT GRAPHIC RENDITION control functions, especially functions which change display colors and background colors.

The runtime treats a colored character as one or two instruction(s). If the background is colored, a `background` instruction is executed. Then, if the text is colored, a `foreground` instruction is executed. Instructions are as described below.

Texts not colored or default colored are treated as comments.

## System

There are 2 memories, `data memory` and `pointer memory`. And the program has one pointer `primal pointer`.

`data memory` is responsible for calculation, IO and condition for jump.

`pointer memory` contains indices for cells of `data memory`.

Both memories use zero-based numbering for indices, accessing indices less than 0 causes undefined behaviour.

`primal pointer` represents the current index for `pointer memory`.

The runtime treats every character as a unicode codepoint, so 'A' is 65, 'あ' is 12354.

## Instructions

In this section, these shorthanded representations are used:
```
data: data memory
ptr: pointer memory
pp: primal pointer
(c): the colored character
```

### foreground

| sequence | color     | instruction  | description |
| -------- | --------- | ------------ | ----------- |
| \033[30m | black     | copy         | `data[ptr[pp]] = (c)` |
| \033[31m | red       | increment    | `data[ptr[pp]]++` |
| \033[32m | green     | input        | `data[ptr[pp]] = input()` |
| \033[33m | yellow    | rotate left  | circular shift left on `data[ptr[pp]]` |
| \033[34m | blue      | rotate right | circular shift right on `data[ptr[pp]]` |
| \033[35m | magenta   | output       | `print(data[ptr[pp]])` This outputs utf-8 encoded character(s) |
| \033[36m | cyan      | decrement    | `data[ptr[pp]]--` |
| \033[37m | white     | swap         | `swap(data[ptr[pp]], ptr[pp])` |
| \033[38m | (ext)     | (error)      |  |
| \033[39m | (default) | nop          | no operation |

### background

| sequence | color     | instruction         | description |
| -------- | --------- | ------------------- | ----------- |
| \033[40m | black     | call                | call external module named `(c).esc` (*NOT IMPLEMENTED IN v0.1.0*) |
| \033[41m | red       | increment ptr       | `ptr[pp]++` |
| \033[42m | green     | jump if zero        | `goto (c) if data[ptr[pp]]==0` |
| \033[43m | yellow    | increment primalPtr | `pp++` |
| \033[44m | blue      | decrement primalPtr | `pp--` |
| \033[45m | magenta   | label               | set a label `(c)`. multiple labels with the same name are prohibited |
| \033[46m | cyan      | decrement ptr       | `ptr[pp]--` |
| \033[47m | white     | exit                | exit the module |
| \033[48m | (ext)     | (error)             |  |
| \033[49m | (default) | nop                 | no operation |
