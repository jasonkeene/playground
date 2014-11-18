nginx:
  pkg:
    - installed
  service:
    - running
    - enable: True
uwsgi:
  pkg:
    - installed
  service:
    - running
    - enable: True
