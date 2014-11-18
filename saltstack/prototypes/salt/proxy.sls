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
    - source: salt://nginx/proxy-pass.conf
