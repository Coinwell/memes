package lightning

import (
	"encoding/json"
	"fmt"

	"github.com/lightningnetwork/lnd/lnrpc"
)

// stakwork payments
// "pay me" (stakwork sends a one-time token)
// payer=stakwork,amount=3000000
// zion pubkey
// zion creates an invoice, "confirm"
// msg sent to stakwork (with invoice, and secret in keysend)
// stakwork

// invites:

// defining deeplinks into zion that do stuff

// "postMessage" action set
// or hovered over button set

const LND_KEYSEND_KEY = 5482373484
const ZION_CUSTOM_RECORD_KEY = 133773310

func receiveInvoice(i *lnrpc.Invoice) {
	if !i.IsKeysend {
		return
	}
	if len(i.Htlcs) < 1 {
		return
	}
	recs := i.Htlcs[0].CustomRecords
	buf := recs[ZION_CUSTOM_RECORD_KEY]

	var payload map[string]interface{}
	err := json.Unmarshal(buf, &payload)
	if err != nil {
		fmt.Println(err)
	}

	value := i.Value

	fmt.Printf("payload: %+v, value: %d\n", payload, value)
}
