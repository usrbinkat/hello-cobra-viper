#!/bin/bash -x
run_build () {
sudo /usr/bin/podman run \
    -it --rm --name higo-builder \
    --volume $(pwd)/bin:/tmp/bin:z \
    --entrypoint /root/dev/build.sh \
    --volume $(pwd):/root/dev:z \
  docker.io/codesparta/red-gotools; \
  return 0
}

run_build
ls -lah $(pwd)
ls -lah $(pwd)/bin
exit 0
