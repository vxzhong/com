package com

import (
	"testing"
)

func TestInetAtoN(t *testing.T) {
	v, err := InetAtoN("192.168.3.168")
	if err != nil || v != 3232236456 {
		t.Errorf("failed!")
	}
}

func TestInetNtoA(t *testing.T) {
	v, err := InetNtoA(3232236456)
	if err != nil || v != "192.168.3.168" {
		t.Errorf("failed!")
	}
}
