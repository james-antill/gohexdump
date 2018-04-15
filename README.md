[![Build Status](https://semaphoreci.com/api/v1/james-antill/gohexdump/branches/master/shields_badge.svg)](https://semaphoreci.com/james-antill/gohexdump)
[![Go Report Card](https://goreportcard.com/badge/github.com/james-antill/gohexdump)](https://goreportcard.com/report/github.com/james-antill/gohexdump)

# Hexdump

 Simple hexdump cmd and API for Go.

## Command

 Outputs 16 bytes at a time, cmd also reads from stdin if no arguments given.

 **Install** with: go get github.com/james-antill/cmd/hexdump

 Roughly equivalent to:

    % rpm -qf /usr/bin/hexdump
    util-linux-2.11r-10

    % hexdump -e '"%08_ax:"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"
                  " " 2/1 "%02X"'
              -e '"  " 16 "%_p" "\n"'

