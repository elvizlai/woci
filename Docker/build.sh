#!/usr/bin/env bash
# swap
# dd if=/dev/zero of=/swapfile bs=1024 count=1024k
# mkswap /swapfile && chmod 0600 /swapfile && swapon /swapfile
# echo "/swapfile       swap    swap    defaults      0       0" >> /etc/fstab

# if docker not exist, install one
# curl -fsSL https://get.docker.com/ | sh
docker build -t woci .