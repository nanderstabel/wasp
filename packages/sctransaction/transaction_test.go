package sctransaction

import (
	"bytes"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	valuetransaction "github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/transaction"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const (
	testAddress = "mtNnGt72bZd25v291TjEzw5uTonExip24cAjtB38x4tq"
)

//
//func TestGenData(t *testing.T) {
//	addr1 := address.RandomOfType(address.VersionED25519)
//	t.Logf("addrEC = %s", addr1.String())
//	addr2 := address.RandomOfType(address.VersionBLS)
//	t.Logf("addrBLS = %s", addr2.String())
//	color := RandomColor()
//	t.Logf("color = %s", color.String())
//}

func randomColor() (ret balance.Color) {
	if _, err := rand.Read(ret[:]); err != nil {
		panic(err)
	}
	return
}

func TestTransactionStateBlockOrigin(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb := NewTransactionBuilder()
	tx, err := txb.Finalize()
	assert.Error(t, err)

	txb = NewTransactionBuilder()
	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)
	bal := balance.New(balance.ColorNew, 1)
	txb.AddBalanceToOutput(addr, bal)

	txb.AddStateBlock(NewStateBlockParams{
		Color:      balance.ColorNew,
		StateIndex: 0,
	})

	tx, err = txb.Finalize()
	assert.NoError(t, err)
	origin, err := tx.ValidateBlocks(&addr)
	assert.NoError(t, err)
	assert.Equal(t, origin, true)

	txb = NewTransactionBuilder()
	o1 = valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)
	bal = balance.New(balance.ColorIOTA, 1)
	txb.AddBalanceToOutput(addr, bal)

	txb.AddStateBlock(NewStateBlockParams{
		Color:      balance.ColorNew,
		StateIndex: 42,
	})

	tx, err = txb.Finalize()
	assert.NoError(t, err)
	_, err = tx.ValidateBlocks(&addr)
	assert.Error(t, err)
}

func TestTransactionStateBlock1(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb := NewTransactionBuilder()
	_, err = txb.Finalize()
	assert.Error(t, err)

	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)
	color := randomColor()
	bal := balance.New(color, 1)
	txb.AddBalanceToOutput(addr, bal)

	txb.AddStateBlock(NewStateBlockParams{
		Color:      color,
		StateIndex: 42,
	})

	tx, err := txb.Finalize()
	assert.NoError(t, err)

	origin, err := tx.ValidateBlocks(&addr)
	assert.NoError(t, err)
	assert.Equal(t, origin, false)

	_, err = txb.Finalize()
	assert.Error(t, err)
}

func TestTransactionStateBlock2(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb := NewTransactionBuilder()
	_, err = txb.Finalize()
	assert.Error(t, err)

	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)
	bal := balance.New(balance.ColorIOTA, 1)
	txb.AddBalanceToOutput(addr, bal)

	txb.AddStateBlock(NewStateBlockParams{
		Color:      balance.ColorNew,
		StateIndex: 42,
	})

	txb.AddRequestBlock(NewRequestBlock(addr, 0))

	tx, err := txb.Finalize()
	assert.NoError(t, err)

	_, err = tx.ValidateBlocks(&addr)
	assert.Error(t, err)
}

func TestTransactionRequestBlock1(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb := NewTransactionBuilder()
	_, err = txb.Finalize()
	assert.Error(t, err)

	txb = NewTransactionBuilder()
	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)
	bal := balance.New(balance.ColorIOTA, 1)
	txb.AddBalanceToOutput(addr, bal)

	reqBlk := NewRequestBlock(addr, 0)
	txb.AddRequestBlock(reqBlk)

	tx, err := txb.Finalize()
	assert.NoError(t, err)

	_, err = tx.ValidateBlocks(&addr)
	assert.Error(t, err)
}

func TestTransactionRequestBlock2(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb := NewTransactionBuilder()
	_, err = txb.Finalize()
	assert.Error(t, err)

	txb = NewTransactionBuilder()
	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)
	bal := balance.New(balance.ColorIOTA, 1)
	txb.AddBalanceToOutput(addr, bal)

	reqBlk := NewRequestBlock(addr, 0)
	txb.AddRequestBlock(reqBlk)
	bal = balance.New(balance.ColorNew, 1)
	txb.AddBalanceToOutput(addr, bal)

	tx, err := txb.Finalize()
	assert.NoError(t, err)

	origin, err := tx.ValidateBlocks(&addr)
	assert.NoError(t, err)
	assert.Equal(t, origin, false)
}

func TestTransactionMultiBlocks(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb := NewTransactionBuilder()
	_, err = txb.Finalize()
	assert.Error(t, err)

	txb = NewTransactionBuilder()
	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	txb.MustAddInputs(o1)

	color := randomColor()
	bal := balance.New(color, 1)
	txb.AddBalanceToOutput(addr, bal)

	txb.AddStateBlock(NewStateBlockParams{
		Color:      color,
		StateIndex: 42,
	})

	reqBlk := NewRequestBlock(addr, 0)
	txb.AddRequestBlock(reqBlk)
	bal = balance.New(balance.ColorNew, 1)
	txb.AddBalanceToOutput(addr, bal)

	tx, err := txb.Finalize()
	assert.NoError(t, err)

	origin, err := tx.ValidateBlocks(&addr)
	assert.NoError(t, err)
	assert.Equal(t, origin, false)
}

func TestDeterminism(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb1 := NewTransactionBuilder()
	txb2 := NewTransactionBuilder()

	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	o2 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	o3 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())

	color1 := randomColor()
	color2 := randomColor()
	color3 := randomColor()

	bal1 := balance.New(color1, 1)
	bal2 := balance.New(color2, 2)
	bal3 := balance.New(color3, 3)

	txb1.MustAddInputs(o1, o2, o3)
	txb1.AddBalanceToOutput(addr, bal1)
	txb1.AddBalanceToOutput(addr, bal2)
	txb1.AddBalanceToOutput(addr, bal3)

	txb1.AddStateBlock(NewStateBlockParams{
		Color:      color1,
		StateIndex: 42,
	})

	tx1, err := txb1.Finalize()
	assert.NoError(t, err)

	txb2.MustAddInputs(o2, o3, o1)
	txb2.AddBalanceToOutput(addr, bal2)
	txb2.AddBalanceToOutput(addr, bal3)
	txb2.AddBalanceToOutput(addr, bal1)

	txb2.AddStateBlock(NewStateBlockParams{
		Color:      color1,
		StateIndex: 42,
	})
	tx2, err := txb2.Finalize()
	assert.NoError(t, err)

	assert.True(t, bytes.Equal(tx1.Bytes(), tx2.Bytes()))
}

func TestCheckInputs(t *testing.T) {
	addr, err := address.FromBase58(testAddress)
	assert.NoError(t, err)

	txb1 := NewTransactionBuilder()
	txb2 := NewTransactionBuilder()

	o1 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())
	o2 := valuetransaction.NewOutputID(addr, valuetransaction.RandomID())

	err = txb1.AddInputs(o1, o2)
	assert.NoError(t, err)

	err = txb2.AddInputs(o1, o1)
	assert.Error(t, err)
}
