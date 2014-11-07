base:
  'roles:master':
    - match: grain
    - master
  'roles:webserver':
    - match: grain
    - webserver
