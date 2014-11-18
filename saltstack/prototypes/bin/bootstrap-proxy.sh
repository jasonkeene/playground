#!/usr/bin/env bash

# update packages
apt-get update
apt-get upgrade --yes

# install salt minion
/vagrant/bin/install-minion.sh

# configure salt minion to pull from master
cat << EOF > /etc/salt/minion
master: $1
grains:
  roles:
    - proxy
EOF
echo $2 > /etc/salt/minion_id

# copy over test_key
mkdir -p /etc/salt/pki/minion
cp /vagrant/salt/test_key.pem /etc/salt/pki/minion/minion.pem
cp /vagrant/salt/test_key.pub /etc/salt/pki/minion/minion.pub

# start miniond
service salt-minion start

# highstate!
salt-call state.highstate

# make sure miniond is running
service salt-minion start

echo
echo
echo "Provision complete!"
