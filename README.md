# Konv

A generalised wrapper.

## Conversion Rules

See [rules](https://github.com/KevinZonda/Konv/blob/master/rules).

The first line is the csv header but also is the command excutor.

We use following syntax to present matching way:

| Symbol | Meaning                         | Syntax       |
|:-------|:--------------------------------|:-------------|
| `,`    | delimiter                       | `apt,pacman` |
| `;`    | command separator               | `-Syu;-S`    |
| `$`    | match 1 word                    | `upgrade $`  |
| `$$`   | match many words, i.e. `[word]` | `update $$`  |
| `@`    | do not append target            | `@ls -al`    |
| `@xxx` | `xxx` without any parse         | `@$` -> `$`  |


TODO

| Symbol  | Meaning             |
|:--------|:--------------------|
| `...`   | match to end        |
| `$var$` | match var           |

## Konv Parameter

`konv` has a built-in parameter which is between `arg[0]` and `arg[1]`.

E.g.

```bash
> apt :y upgrade
#  +  ++ ++++++++
#  |   |    +------- Arguments
#  |   +----- Konv Parameter
#  Konv
```

In most case, konv parameter can be ignored

| Parameter     | Meaning        |
|:--------------|:---------------|
| `:y` or `:sc` | ignore confirm |
| `:c`          | show confirm   |
| `:l`          | show run list  |
| `:v`          | verbose        |

## Paths

Usually, you should put your `cfg`s & `csv`s to

- `$HOME/.config`
- `/etc/konv` (for *nix & macOS)
- `%AppData%` (for Windows)