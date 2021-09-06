#!/bin/bash

source .food.env

DATABASE_HOST=localhost

go run main.go $*
