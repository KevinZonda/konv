# apt-pac

An apt syntax-like wrapper for pacman

## Conversion Rules

See [conv.csv](https://github.com/KevinZonda/apt-pac/blob/master/conv.csv).
We use following syntax to present matching way:

| Symbol | Meaning |
| :-- | :-- |
| `,` | delimiter |
| `;` | command separator |
| `$` | match 1 word |
| `$$` | match many words, i.e. `[word]` |

The first line is the csv header but also is the command excutor.
