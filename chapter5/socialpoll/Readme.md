Start nsqlookupd so that our nsqd instances are discoverable

Start nsqd and tell it which nsqlookupd to use

Start mongod for data services

---
nsqlookupd

nsqd --lookupd-tcp-address=localhost:4160


Run mongod, will need dir db creatd beforehand.
mongod --dbpath ./db

--
Run ./socialpoll/setup.sh to load twitter env variables