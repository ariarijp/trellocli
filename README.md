trellocli
=====================

## Features

### Show Cards

```
$ trellocli
* How to Use Trello for Android (https://trello.com/b/9dnaRkNt)
  * Getting Started
    * Welcome to Trello! This is a card.
    * Tap on a card to open it up.
{blue }
    * Color-coded labels can be used to classify cards.
    * Create as many cards as you want. We've got an unlimited supply!
    * Here is a picture of Taco, our Siberian Husky.
  * Diving In
    * Press and hold this card to drag it to another list.
    * Make as many lists and boards as you need. We'll make more!
  * Mastering Trello
    * Finished with a card? Drag it to the top of the board to archive it.
    * You can reorder lists too.
    * Invite team members to collaborate on this board.
  * More Info
    * Want updates on new features?
    * You can also view your boards on trello.com
```

### TODO

* Add Card

## Installation

```shell
$ go get github.com/ariarijp/trellocli
```

## Configuration

Create `.trelloclirc` file in your home directory.

```toml
app_key = "YOUR_TRELLO_APP_KEY"
token = "YOUR_TRELLO_TOKEN"
username = "YOUR_TRELLO_USERNAME"
```

Then, you can use `trellocli` command.

```
$ trellocli
```

## License

MIT

## Author

[ariarijp](https://github.com/ariarijp)
