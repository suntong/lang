#!/bin/bash

set -ex

#go build -o app

cat >1.json<<EOF
{
	"host": "127.0.0.1",
	"port": 8080
}
EOF

echo ===run root command
go run 027-global-option.go -h

echo =case 1: read from file
go run 027-global-option.go

echo =case 2: read from ENV variables
HOST=10.0.0.1 go run 027-global-option.go

echo =case 3: read from command line
HOST=10.0.0.1 go run 027-global-option.go -H 168.0.0.1

echo ===run sub command
go run 027-global-option.go sub -h

echo =case 1: read from file
go run 027-global-option.go sub

echo =case 2: read from ENV variables
HOST=10.0.0.1 go run 027-global-option.go sub

echo =case 3: read from command line
HOST=10.0.0.1 go run 027-global-option.go sub -H 168.0.0.1
HOST=10.0.0.1 go run 027-global-option.go sub -H 168.0.0.2 -w something

rm 1.json
#rm app
