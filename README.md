 Simple hexdump cmd and API for Go.

 Outputs 16 bytes at a time, cmd also reads from stdin if no arguments given.

 Install with: go get github.com/james-antill/cmd/hexdump

 Roughly equivalent to:

hexdump -e '"%08_ax:"
            " " 2/1 "%02X"
            " " 2/1 "%02X"
            " " 2/1 "%02X"
            " " 2/1 "%02X"
            " " 2/1 "%02X"
            " " 2/1 "%02X"
            " " 2/1 "%02X"
            " " 2/1 "%02X"'
        -e '"  " 16 "%_p" "\n"'

