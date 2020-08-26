package filecoin

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/shopspring/decimal"
	"testing"
)

// 发送FileCoin
func TestClient_MpoolPush(t *testing.T) {
	c := testClient()

	from, _ := address.NewFromString("t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	to, _ := address.NewFromString("t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi")

	msg := &Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      FromFil(decimal.NewFromFloat(1)),
		GasLimit:   0,
		GasFeeCap:  decimal.Zero,
		GasPremium: decimal.Zero,
		Method:     0,
		Params:     nil,
	}

	msg, err := c.GasEstimateMessageGas(context.Background(), msg, &MessageSendSpec{MaxFee: "1000000000000"}, nil)
	if err != nil {
		t.Error(err)
	}

	actor, err := c.StateGetActor(context.Background(), msg.From, nil)
	if err != nil {
		t.Error(err)
	}

	msg.Nonce = actor.Nonce
	t.Log(msg)

	sm, err := c.WalletSignMessage(context.Background(), msg.From, msg)
	if err != nil {
		t.Error(err)
	}

	id, err := c.MpoolPush(context.Background(), sm)
	if err != nil {
		t.Error(err)
	}

	t.Log(id.Version())
	t.Log(id)
}