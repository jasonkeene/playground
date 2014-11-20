/etc/hosts:
  file.managed:
    - template: jinja
    - source: salt://network/hosts
