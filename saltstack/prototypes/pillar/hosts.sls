hosts:
  - hostname: master
    address: 10.13.37.10
    aliases:
      - salt-master
  - hostname: proxy1
    address: 10.13.37.11
    aliases:
      - nginx1
  - hostname: proxy2
    address: 10.13.37.12
    aliases:
      - nginx2
  - hostname: app1
    address: 10.13.37.13
    aliases:
      - application1
  - hostname: app2
    address: 10.13.37.14
    aliases:
      - application2
