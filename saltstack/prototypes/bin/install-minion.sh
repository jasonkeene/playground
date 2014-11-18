#!/usr/bin/env bash

# install salt minion
add-apt-repository --yes ppa:saltstack/salt-testing
#echo deb http://ppa.launchpad.net/saltstack/salt-testing/ubuntu `lsb_release -sc` main | tee /etc/apt/sources.list.d/saltstack.list
wget -q -O- "http://keyserver.ubuntu.com:11371/pks/lookup?op=get&search=0x4759FA960E27C0A6" | apt-key add -
apt-get update
apt-get install --yes salt-minion=2014.7.0+ds-2trusty1

# stop miniond
service salt-minion stop
