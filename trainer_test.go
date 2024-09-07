package tdd_pokemon

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TrainerTestSuite struct {
	suite.Suite
	Trainer
}

func (suite *TrainerTestSuite) SetupTest() {
	suite.Trainer = *NewTrainer()
}

func (suite *TrainerTestSuite) TestQuickFilterIVTrue() {
	pokemon := Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 13}}
	suite.True(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterIVFalse() {
	pokemon := Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 12}}
	suite.False(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterPVPTrue() {
	pokemon := Pokemon{iv: IV{Attack: 6, Defense: 12, HP: 12}}
	suite.True(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterPVPFalse() {
	pokemon := Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 6}}
	suite.False(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterLegendaryTrue() {
	pokemon := Pokemon{no: 144}
	suite.True(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterLegendaryFalse() {
	pokemon := Pokemon{no: 1}
	suite.False(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterShinyTrue() {
	pokemon := Pokemon{no: 25, shiny: true}
	suite.True(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterShinyFalse() {
	pokemon := Pokemon{no: 25, shiny: false}
	suite.False(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterEventTrue() {
	pokemon := Pokemon{no: 25, event: true}
	suite.True(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestQuickFilterEventFalse() {
	pokemon := Pokemon{no: 25, event: false}
	suite.False(suite.Trainer.QuickFilter(pokemon))
}

func (suite *TrainerTestSuite) TestPreciseFilterIVTrue() {
	pokemon := Pokemon{iv: IV{Attack: 14, Defense: 14, HP: 14}}
	suite.True(suite.Trainer.PreciseFilter(pokemon))
}

func (suite *TrainerTestSuite) TestPreciseFilterIVFalse() {
	pokemon := Pokemon{iv: IV{Attack: 13, Defense: 13, HP: 13}}
	suite.False(suite.Trainer.PreciseFilter(pokemon))
}

func (suite *TrainerTestSuite) TestPreciseFilterPVPTrue() {
	suite.Trainer.storage.add(Pokemon{no: 1, iv: IV{Attack: 3, Defense: 14, HP: 14}})
	suite.Trainer.storage.add(Pokemon{no: 979, iv: IV{Attack: 2, Defense: 15, HP: 15}})
	suite.True(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestPreciseFilterPVPFalse() {
	suite.Trainer.storage.add(Pokemon{no: 979, iv: IV{Attack: 3, Defense: 14, HP: 14}})
	suite.Trainer.storage.add(Pokemon{no: 979, iv: IV{Attack: 2, Defense: 15, HP: 15}})
	suite.False(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestPreciseFilterLegendaryTrue() {
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 13, Defense: 13, HP: 13}})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 14, Defense: 14, HP: 14}})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 15, Defense: 15, HP: 15}})
	suite.True(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestPreciseFilterLegendaryFalse() {
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 12, Defense: 12, HP: 12}})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 13, Defense: 13, HP: 13}})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 14, Defense: 14, HP: 14}})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 15, Defense: 15, HP: 15}})
	suite.False(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestPreciseFilterShinyTrue() {
	suite.Trainer.storage.add(Pokemon{no: 25, iv: IV{Attack: 11, Defense: 11, HP: 11}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 25, iv: IV{Attack: 12, Defense: 12, HP: 12}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 25, iv: IV{Attack: 13, Defense: 13, HP: 13}})
	suite.True(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[1]))
}

func (suite *TrainerTestSuite) TestPreciseFilterShinyFalse() {
	suite.Trainer.storage.add(Pokemon{no: 25, iv: IV{Attack: 11, Defense: 11, HP: 11}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 25, iv: IV{Attack: 12, Defense: 12, HP: 12}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 25, iv: IV{Attack: 13, Defense: 13, HP: 13}, shiny: true})
	suite.False(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestPreciseFilterShinyLegendaryTrue() {
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 11, Defense: 10, HP: 10}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 11, Defense: 11, HP: 11}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 12, Defense: 12, HP: 12}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 13, Defense: 13, HP: 13}})
	suite.True(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestPreciseFilterShinyLegendaryFalse() {
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 10, Defense: 10, HP: 10}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 11, Defense: 11, HP: 11}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 12, Defense: 12, HP: 12}, shiny: true})
	suite.Trainer.storage.add(Pokemon{no: 144, iv: IV{Attack: 13, Defense: 13, HP: 13}, shiny: true})
	suite.False(suite.Trainer.PreciseFilter(suite.Trainer.storage.pokemon[0]))
}

func (suite *TrainerTestSuite) TestManage() {
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 12}})
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 13, Defense: 13, HP: 13}})
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 14, Defense: 14, HP: 14}})
	suite.Trainer.Manage()
	suite.Equal(1, len(suite.Trainer.storage.pokemon))
	suite.Equal(Pokemon{iv: IV{Attack: 14, Defense: 14, HP: 14}}, suite.Trainer.storage.pokemon[0])
}

func (suite *TrainerTestSuite) TestGotcha() {
	p, i := suite.Trainer.Gotcha()
	suite.IsType(Pokemon{}, p)
	suite.Equal(1, len(suite.Trainer.storage.pokemon))
	suite.Equal(0, i)
}

func (suite *TrainerTestSuite) TestTransferSuccess() {
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 12}})
	err := suite.Trainer.Transfer(0)
	suite.NoError(err)
	suite.Equal(0, len(suite.Trainer.storage.pokemon))
}

func (suite *TrainerTestSuite) TestTransferTwiceSuccess() {
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 12}})
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 14, Defense: 14, HP: 14}})
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 13, Defense: 13, HP: 13}})
	_ = suite.Trainer.Transfer(2)
	_ = suite.Trainer.Transfer(0)
	suite.Equal(1, len(suite.Trainer.storage.pokemon))
	suite.Equal(Pokemon{iv: IV{Attack: 14, Defense: 14, HP: 14}}, suite.Trainer.storage.pokemon[0])
}

func (suite *TrainerTestSuite) TestTransferFail() {
	suite.Trainer.storage.add(Pokemon{iv: IV{Attack: 12, Defense: 12, HP: 12}})
	err := suite.Trainer.Transfer(1)
	suite.Error(err)
}

func TestTrainerTestSuite(t *testing.T) {
	suite.Run(t, new(TrainerTestSuite))
}
