package netplan

import (
	"testing"

	yamlnillable "github.com/ploynomail/go-yaml-nillable"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyBond(t *testing.T) {
	given := Bond{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal Bond
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeBond(t *testing.T) {
	given := Bond{
		Device: Device{
			DHCP4: yamlnillable.BoolOf(true),
			DHCP6: yamlnillable.BoolOf(false),
		},
		Parameters: &BondParameters{
			Mode: ActiveBackupBondMode(),
		},
		Interfaces: []string{"inter-1"},
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
parameters:
  mode: active-backup
interfaces:
- inter-1
`), marshal)

	var unmarshal Bond
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
