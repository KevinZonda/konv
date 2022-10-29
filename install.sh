#!/bin/bash
# Copyright (C) 2022 KevinZonda
# This file is to install apt-pac, an apt
# syntax-like wraper for pacman
# Please run in sudo mod

# setup folder
mkdir /etc/apt-pac
chmod +r /etc/apt-pac

# copy rule file
cp ./conv.csv /etc/apt-pac/rule.csv

# compile apt-pac
cd src
bash build.sh

# move apt-pac into /usr/bin/
cp ./out/apt /usr/bin/apt

# set permission
chmod +x /usr/bin/apt

echo 'done.'
