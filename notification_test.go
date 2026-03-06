package goesi_test

import (
	"testing"

	"github.com/fnt-eve/goesi-openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const MercenaryDenAttackedText = `
aggressorAllianceName: <a href="showinfo:16159//3011">Bad Alliance</a>
aggressorCharacterID: 1011
aggressorCorporationName: <a href="showinfo:2//2011">Bad Corp</a>
armorPercentage: 99.0
hullPercentage: 100.0
itemID: &id001 1000000000099
mercenaryDenShowInfoData:
- showinfo
- 85230
- *id001
planetID: 40161469
planetShowInfoData:
- showinfo
- 11
- 40161469
shieldPercentage: 0.0
solarsystemID: 30002537
typeID: 85230
`

func TestNotification_MercenaryDenAttacked(t *testing.T) {
	var got goesi.MercenaryDenAttacked
	err := yaml.Unmarshal([]byte(MercenaryDenAttackedText), &got)
	require.NoError(t, err)
	assert.Equal(t, int64(1011), got.AggressorCharacterID)
	assert.Equal(t, 99.0, got.ArmorPercentage)
	assert.Equal(t, 100.0, got.HullPercentage)
	assert.Equal(t, int64(40161469), got.PlanetID)
	assert.Equal(t, 0.0, got.ShieldPercentage)
	assert.Equal(t, int64(30002537), got.SolarSystemID)
	assert.Equal(t, int64(85230), got.TypeID)
}

const MercenaryDenReinforcedText = `
aggressorAllianceName: <a href="showinfo:16159//3011">Bad Alliance</a>
aggressorCharacterID: 1011
aggressorCorporationName: <a href="showinfo:2//2011">Bad Corp</a>
itemID: &id001 1000000000099
mercenaryDenShowInfoData:
- showinfo
- 85230
- *id001
planetID: 40161469
planetShowInfoData:
- showinfo
- 11
- 40161469
solarsystemID: 30002537
timestampEntered: 134026459143594820
timestampExited: 134027264983594820
typeID: 85230
`

func TestNotification_MercenaryDenReinforced(t *testing.T) {
	var got goesi.MercenaryDenReinforced
	err := yaml.Unmarshal([]byte(MercenaryDenReinforcedText), &got)
	require.NoError(t, err)
	assert.Equal(t, int64(1011), got.AggressorCharacterID)
	assert.Equal(t, int64(40161469), got.PlanetID)
	assert.Equal(t, int64(30002537), got.SolarSystemID)
	assert.Equal(t, int64(85230), got.TypeID)
	assert.Equal(t, int64(134026459143594820), got.TimestampEntered)
	assert.Equal(t, int64(134027264983594820), got.TimestampExited)
}
