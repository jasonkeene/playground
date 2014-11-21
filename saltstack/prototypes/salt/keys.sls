/home/vagrant/.ssh/authorized_keys:
  file.managed:
    - template: jinja
    - source: salt://ssh/authorized_keys
