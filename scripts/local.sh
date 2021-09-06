#!/bin/bash

source ./build/package/.food.env

DATABASE_HOST=localhost

go run main.go $*
