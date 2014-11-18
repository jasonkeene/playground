nginx:
  pkg:
    - installed
  service:
    - running
    - enable: True
    - watch:
      - file: /etc/nginx/sites-available/default

/etc/nginx/sites-available/default:
  file.managed:
    - source: salt://nginx/app.conf

uwsgi-plugin-python:
  pkg:
    - installed

uwsgi:
  pkg:
    - installed
  service:
    - running
    - enable: True
    - watch:
      - file: /etc/uwsgi/apps-enabled/app.json
      - file: /srv/wsgi.py

/etc/uwsgi/apps-enabled/app.json:
  file.managed:
    - source: salt://uwsgi/app.json

/srv/wsgi.py:
  file.managed:
    - source: salt://app/wsgi.py
