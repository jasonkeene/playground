base:
  '*':
    #- utilities
    - hosts
  'roles:master':
    - match: grain
    - master
  'roles:proxy':
    - match: grain
    - proxy
  'roles:app':
    - match: grain
    - app
