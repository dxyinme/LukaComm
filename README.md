# Luka - comm

base libraries are used in luka-im

[![Build Status](https://travis-ci.com/dxyinme/LukaComm.svg?branch=master)](https://travis-ci.com/dxyinme/LukaComm)

#### Protocol design

```
 -- LukaMsg
|checkpoint - 1byte|From - 32 byte|Target - 32 byte|MsgType - 2 byte|MsgContentType - 2 byte| content ... 
```
