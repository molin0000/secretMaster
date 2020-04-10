#!/bin/bash

# Build Web

cd ui/webpage
umi build
cd ..
rm rice-box.go
rice embed-go
cd ..

docker start -a build-gocqplg
./cqcfg -c .
rm -rf ~/coolq/dev/me.cqp.molin.secretmaster
mkdir -p ~/coolq/dev/me.cqp.molin.secretmaster
cp app.dll ~/coolq/dev/me.cqp.molin.secretmaster/
cp app.json ~/coolq/dev/me.cqp.molin.secretmaster/

rm -rf ~/coolqPro/dev/me.cqp.molin.secretmaster
mkdir -p ~/coolqPro/dev/me.cqp.molin.secretmaster
cp app.dll ~/coolqPro/dev/me.cqp.molin.secretmaster/
cp app.json ~/coolqPro/dev/me.cqp.molin.secretmaster/

