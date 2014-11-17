base:
  'roles:master':
    - match: grain
    - utilities
    - master
  'roles:webserver':
    - match: grain
    - utilities
    - webserver
  'roles:application':
    - match: grain
    - utilities
    - application
