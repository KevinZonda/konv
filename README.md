# Kov

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

## Some Tool

- `apt :y xxx`: it will skip confirm of pacman command
  check, but will not say yes to pacman
