package netplan

import (
	"testing"

	yamlnillable "github.com/ploynomail/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyNetwork(t *testing.T) {
	given := Network{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`network: null
`), marshal)

	var unmarshal Network
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeNetworkExample1(t *testing.T) {
	// https://netplan.io/reference#examples
	given := Network{
		Network: &NetworkConfig{
			Ethernets: Ethernets{
				"eno1": &Ethernet{
					Device: Device{
						DHCP4: yamlnillable.BoolOf(true),
					},
				},
			},
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`network:
  version: 2
  ethernets:
    eno1:
      dhcp4: true
`), marshal)

	var unmarshal Network
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeNetworkExample2(t *testing.T) {
	// https://netplan.io/reference#examples
	given := Network{
		Network: &NetworkConfig{
			Renderer: NetworkdRenderer(),
			Ethernets: Ethernets{
				"eno1": &Ethernet{
					Device: Device{
						Addresses: []*Address{
							{
								Address:   "10.0.0.10",
								PrefixLen: yamlnillable.Uint8Of(24),
							},
							{
								Address:   "11.0.0.11",
								PrefixLen: yamlnillable.Uint8Of(24),
							},
						},
						NameServers: &Nameservers{
							Addresses: []string{"8.8.8.8", "8.8.4.4"},
						},
						Routing: Routing{
							Routes: []*Route{
								{
									To: &Address{
										Address:   "0.0.0.0",
										PrefixLen: yamlnillable.Uint8Of(0),
									},
									Via: &Address{
										Address: "10.0.0.1",
									},
									Metric: yamlnillable.Uint64Of(100),
								},
								{
									To: &Address{
										Address:   "0.0.0.0",
										PrefixLen: yamlnillable.Uint8Of(0),
									},
									Via: &Address{
										Address: "11.0.0.1",
									},
									Metric: yamlnillable.Uint64Of(100),
								},
							},
						},
					},
				},
			},
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`network:
  version: 2
  renderer: networkd
  ethernets:
    eno1:
      addresses:
      - 10.0.0.10/24
      - 11.0.0.11/24
      nameservers:
        addresses:
        - 8.8.8.8
        - 8.8.4.4
      routes:
      - to: 0.0.0.0/0
        via: 10.0.0.1
        metric: 100
      - to: 0.0.0.0/0
        via: 11.0.0.1
        metric: 100
`), marshal)

	var unmarshal Network
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeNetworkExample3(t *testing.T) {
	// https://netplan.io/reference#examples
	given := Network{
		Network: &NetworkConfig{
			Renderer: NetworkManagerRenderer(),
			Ethernets: Ethernets{
				"id0": &Ethernet{
					PhysicalDevice: PhysicalDevice{
						Match: &Match{
							MacAddress: yamlnillable.StringOf("00:11:22:33:44:55"),
						},
						WakeOnLAN: yamlnillable.BoolOf(true),
					},
					Device: Device{
						DHCP4: yamlnillable.BoolOf(true),
						Addresses: []*Address{
							{
								Address:   "192.168.14.2",
								PrefixLen: yamlnillable.Uint8Of(24),
							},
							{
								Address:   "192.168.14.3",
								PrefixLen: yamlnillable.Uint8Of(24),
							},
							{
								Address:   "2001:1::1",
								PrefixLen: yamlnillable.Uint8Of(64),
							},
						},
						Gateway4: yamlnillable.StringOf("192.168.14.1"),
						Gateway6: yamlnillable.StringOf("2001:1::2"),
						NameServers: &Nameservers{
							Search:    []string{"foo.local", "bar.local"},
							Addresses: []string{"8.8.8.8"},
						},
						Routing: Routing{
							Routes: []*Route{
								{
									To: &Address{
										Address:   "0.0.0.0",
										PrefixLen: yamlnillable.Uint8Of(0),
									},
									Via: &Address{
										Address: "11.0.0.1",
									},
									Table:  yamlnillable.Uint64Of(70),
									OnLink: yamlnillable.BoolOf(true),
									Metric: yamlnillable.Uint64Of(3),
								},
							},
							RoutingPolicy: []*RoutingPolicy{
								{
									From: &Address{
										Address:   "192.168.14.2",
										PrefixLen: yamlnillable.Uint8Of(24),
									},
									To: &Address{
										Address:   "10.0.0.0",
										PrefixLen: yamlnillable.Uint8Of(8),
									},
									Table:    yamlnillable.Uint64Of(70),
									Priority: yamlnillable.Uint32Of(100),
								},
								{
									From: &Address{
										Address:   "192.168.14.3",
										PrefixLen: yamlnillable.Uint8Of(24),
									},
									To: &Address{
										Address:   "20.0.0.0",
										PrefixLen: yamlnillable.Uint8Of(8),
									},
									Table:    yamlnillable.Uint64Of(70),
									Priority: yamlnillable.Uint32Of(50),
								},
							},
						},
						Renderer: NetworkdRenderer(),
					},
				},
				"lom": &Ethernet{
					PhysicalDevice: PhysicalDevice{
						Match: &Match{
							Driver: yamlnillable.StringOf("ixgbe"),
						},
						SetName: yamlnillable.StringOf("lom1"),
					},
					Device: Device{
						DHCP6: yamlnillable.BoolOf(true),
					},
				},
				"switchports": &Ethernet{
					PhysicalDevice: PhysicalDevice{
						Match: &Match{
							Name: yamlnillable.StringOf("enp2*"),
						},
					},
					Device: Device{
						MTU: yamlnillable.Uint64Of(1280),
					},
				},
			},
			Wifis: Wifis{
				"all-wlans": &Wifi{
					AccessPoints: AccessPoints{
						"Joe's home": &AccessPoint{
							Password: yamlnillable.StringOf("s3kr1t"),
						},
					},
				},
				"wlp1s0": &Wifi{
					AccessPoints: AccessPoints{
						"guest": &AccessPoint{
							Mode: APAccessPointMode(),
						},
					},
				},
			},
			Bridges: Bridges{
				"br0": &Bridge{
					Device: Device{
						DHCP4: yamlnillable.BoolOf(true),
					},
					Interfaces: []string{"wlp1s0", "switchports"},
				},
			},
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`network:
  version: 2
  renderer: NetworkManager
  ethernets:
    id0:
      dhcp4: true
      addresses:
      - 192.168.14.2/24
      - 192.168.14.3/24
      - 2001:1::1/64
      gateway4: 192.168.14.1
      gateway6: 2001:1::2
      nameservers:
        search:
        - foo.local
        - bar.local
        addresses:
        - 8.8.8.8
      renderer: networkd
      routes:
      - to: 0.0.0.0/0
        via: 11.0.0.1
        on-link: true
        metric: 3
        table: 70
      routing-policy:
      - from: 192.168.14.2/24
        to: 10.0.0.0/8
        table: 70
        priority: 100
      - from: 192.168.14.3/24
        to: 20.0.0.0/8
        table: 70
        priority: 50
      match:
        macaddress: "00:11:22:33:44:55"
      wakeonlan: true
    lom:
      dhcp6: true
      match:
        driver: ixgbe
      set-name: lom1
    switchports:
      mtu: 1280
      match:
        name: enp2*
  wifis:
    all-wlans:
      access-points:
        Joe's home:
          password: s3kr1t
    wlp1s0:
      access-points:
        guest:
          mode: ap
  bridges:
    br0:
      dhcp4: true
      interfaces:
      - wlp1s0
      - switchports
`), marshal)

	var unmarshal Network
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
