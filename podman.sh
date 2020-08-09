#!/bin/bash
sudo podman run \
    -it --rm --net=host \
    --volume $(pwd):/root/dev \
    --volume ~/.ssh:/root/.ssh \
    --volume ~/.bashrc:/root/.bashrc \
    --volume ~/.gitconfig:/root/.gitconfig \
    --name higo --hostname higo \
  docker.io/ocpredshift/red-gotools -c /usr/bin/tmux
