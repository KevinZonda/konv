#!/bin/bash
# Copyright (C) 2022 KevinZonda
# This file is to install konv, a generalised wrapper
# Please run in sudo mod

# setup folder
mkdir /etc/apt-pac
chmod +r /etc/apt-pac

# copy rules file
cp ./conv.csv /etc/apt-pac/rule.csv

# compile konv
cd src
bash build.sh

# move konv into /usr/bin/
cp ./out/apt /usr/bin/apt

# set permission
chmod +x /usr/bin/apt

echo 'done.'
