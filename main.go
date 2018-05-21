package main

import (
"github.com/alfg/blockchain"
"fmt"
"net/http"
	"encoding/json"
)


type Block struct {
	Hash    string
	Height int
	Time    int
}

type Transaction struct {
	Raw    string
	Block Block
}

type Transactions struct {
	Transactions    []Transaction
}


var hash = make(map[string]Transaction)


func main() {
	http.HandleFunc("/", getTransactions)
	http.ListenAndServe(":3000", nil)

}

type Profile struct {
	Name    string
	Hobbies []string
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	address:=r.FormValue("address")

	if address==""{
		http.Error(w, "Address is not set", 400)
		return
	}

	fmt.Print("Requested address=", address)

	c, _ := blockchain.New()
	resp, e := c.GetAddress(address)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	var Transactions Transactions
	for i := range resp.Txs {
		var trx Transaction
		//WHERE IS RAW?
		trx.Raw="????"

		trx.Block.Time=resp.Txs[i].Time
		trx.Block.Hash=resp.Txs[i].Hash
		trx.Block.Height=resp.Txs[i].BlockHeight

		//Append to Transactions
		Transactions.Transactions=append(Transactions.Transactions, trx)
	}

	js, err := json.Marshal(Transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
