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
    - app
EOF
echo $2 > /etc/salt/minion_id

# copy over test_key
mkdir -p /etc/salt/pki/minion
cp /vagrant/salt/minion.{pem,pub} /etc/salt/pki/minion/

# highstate!
salt-call state.highstate

# start miniond
service salt-minion start

echo
echo
echo "Provision complete!"
