# apt-pac

An apt syntax-like wrapper for pacman

## Conversion Rules

See [conv.csv](https://github.com/KevinZonda/apt-pac/blob/master/conv.csv).

The first line is the csv header but also is the command excutor.

We use following syntax to present matching way:

| Symbol | Meaning | Syntax |
| :-- | :-- | :-- |
| `,` | delimiter | `apt,pacman` |
| `;` | command separator | `-Syu;-S` |
| `$` | match 1 word | `upgrade $` |
| `$$` | match many words, i.e. `[word]` | `update $$` |

## Some Tool

- `apt y xxx`: it will skip confirm of pacman command
  check, but will not say yes to pacman
