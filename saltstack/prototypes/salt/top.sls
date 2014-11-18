base:
  'roles:master':
    - match: grain
    - utilities
    - master
  'roles:proxy':
    - match: grain
    - utilities
    - proxy
  'roles:app':
    - match: grain
    - utilities
    - app
