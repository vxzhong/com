package com

import (
	"errors"
	"math"
	"math/big"
	"net"
)

func InetNtoA(number uint) (string, error) {
	if number > math.MaxUint32 {
		return "", errors.New("too large ipv4 number")
	}

	var bytes [4]byte
	bytes[0] = byte(number & 0xFF)
	bytes[1] = byte((number >> 8) & 0xFF)
	bytes[2] = byte((number >> 16) & 0xFF)
	bytes[3] = byte((number >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String(), nil
}

func InetAtoN(ip string) (int64, error) {
	to4 := net.ParseIP(ip).To4()
	if to4 == nil {
		return 0, errors.New("not a valid ipv4 address")
	}

	ipv4Int := big.NewInt(0)
	ipv4Int.SetBytes(to4)
	return ipv4Int.Int64(), nil
}
