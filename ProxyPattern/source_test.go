package ProxyPattern

import "testing"

func TestProxyPattern(t *testing.T) {
	girl := SchoolGirl{name: "jiaojiao"}
	pursuit := PurSuit{girl}
	proxy := Proxy{pursuit}
	proxy.GiveDolls()
	proxy.GiveFollowers()
	proxy.GiveChocolate()
}
