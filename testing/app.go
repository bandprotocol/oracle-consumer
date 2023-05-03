package testing

import (
	"encoding/json"

	oracleapp "github.com/bandprotocol/oracle-consumer/app"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	ibctesting "github.com/cosmos/ibc-go/v5/testing"
	ibctestingtypes "github.com/cosmos/ibc-go/v5/testing/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}

type TestingApp struct {
	*oracleapp.App
}

func (app *TestingApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper implements the TestingApp interface.
func (app *TestingApp) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *TestingApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *TestingApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *TestingApp) GetTxConfig() client.TxConfig {
	return oracleapp.MakeEncodingConfig().TxConfig
}

func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := oracleapp.MakeEncodingConfig()
	app := &TestingApp{
		App: oracleapp.New(
			log.NewNopLogger(),
			db,
			nil,
			true,
			map[int64]bool{},
			oracleapp.DefaultNodeHome,
			5,
			encCdc,
			EmptyAppOptions{},
		),
	}

	return app, oracleapp.NewDefaultGenesisState(encCdc.Marshaler)
}

func GetTestingApp(chain *ibctesting.TestChain) *TestingApp {
	app, ok := chain.App.(*TestingApp)
	if !ok {
		panic("not testing app")
	}

	return app
}

func init() {
	ibctesting.DefaultTestingAppInit = SetupTestingApp
}
