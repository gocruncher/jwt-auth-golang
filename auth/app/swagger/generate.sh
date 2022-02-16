#!/usr/bin/env bash

swagger-codegen generate -i ./swagger.yml -l go -o ./ -D models {opts}
rm -rf ./docs