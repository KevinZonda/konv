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
| `@` | if a command starts with it means the command will not use the excutor which is mentioned at csv header | `update;@ls`|
