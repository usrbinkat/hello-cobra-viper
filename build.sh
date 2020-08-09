#!/bin/bash
# cobra init --pkg-name github.com/usrbinkat/go-hello-viper
# cobra add hi
# cobra add bye
# go build
# gitup devel

goCmd=$(which go)

rm /bin/higo 2>/dev/null
rm -rf /root/higo 2>/dev/null
mkdir -p /tmp/bin

plugins="
  "github.com/usrbinkat/hello-cobra-viper/cmd" \
"
# "github.com/spf13/viper" \
# "github.com/spf13/cobra" \

for i in ${plugins}; do
  ${goCmd} get -u ${i};
done

${goCmd} build

cp -f ./dev /tmp/bin/higo 2>/dev/null
ls -lah /tmp/bin

