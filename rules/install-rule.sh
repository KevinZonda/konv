#!/bin/bash

if_empty_exit () {
  if [ -z "$1" ]
  then
    echo "Empty detected!"
    exit 1
  fi
}

ask_yes_or_exit () {
    read -r -p "Continue? [Y/n] " ask_yes
    case $ask_yes in
        [yY][eE][sS]|[yY])
            ;;
        *)
            echo "Abort..."
            exit 1
            ;;
    esac
}

from=$1
to=$2

if_empty_exit "$from"
echo "From: $from"
if_empty_exit "$to"
echo "To: $to"
echo "So we are going to copy ./$from-$to.csv to /etc/konv/$from.csv"
ask_yes_or_exit
mkdir -p /etc/konv
cp "./$from-$to.csv" "/etc/konv/$from.csv"

echo "Make a link so that you can use the konv?"
ask_yes_or_exit
ln -s "/usr/local/bin/konv" "/usr/local/bin/$from"