#!/bin/bash
# Copyright (C) 2022 KevinZonda
# This file is to install konv, a generalised wrapper
# Please run in sudo mod

# setup folder
mkdir /etc/konv
chmod +r /etc/konv

# compile konv
cd src
bash build.sh

# move konv into /usr/bin/
cp ./out/konv /usr/local/bin/konv

# set permission
chmod +x /usr/local/bin/konv

echo 'done.'
