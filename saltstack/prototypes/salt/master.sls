salt-master:
  pkg:
    - installed
  service:
    - running
    - enable: True
