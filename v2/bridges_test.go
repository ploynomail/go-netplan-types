package netplan

import (
	"testing"

	yamlnillable "github.com/ploynomail/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyBridge(t *testing.T) {
	given := Bridge{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Bridge
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeBridge(t *testing.T) {
	given := Bridge{
		Device: Device{
			DHCP4: yamlnillable.BoolOf(true),
			DHCP6: yamlnillable.BoolOf(false),
		},
		Interfaces: []string{"vlan1", "vlan2"},
		Parameters: &BridgeParameters{
			STP: yamlnillable.BoolOf(false),
		},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
interfaces:
- vlan1
- vlan2
parameters:
  stp: false
`), marshal)

	var unmarshal Bridge
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
