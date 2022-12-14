#!/bin/bash
ITUN=chain0
TUNIP=10.111.111.2/24

ip tuntap add dev $ITUN mod tun
ip addr add $TUNIP dev $ITUN
ip link set $ITUN up