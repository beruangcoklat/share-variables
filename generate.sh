#!/bin/bash

protoc proto/sharevariable.proto --go_out=plugins=grpc:.;