package main

import (
	"github.com/ghaoo/wxwork"
	"os"
	"strconv"
)

type payload struct {
	client *wxwork.Agent
}

func New() *payload {
	return &payload{
		client: newAgent(),
	}
}

func newAgent() *wxwork.Agent {
	corpid := os.Getenv("WORKWX_CORP_ID")
	secret := os.Getenv("WORKWX_SECRET")
	agentid, err := strconv.Atoi(os.Getenv("WORKWX_AGENT_ID"))
	if err != nil {
		panic(err)
	}
	a := wxwork.NewAgent(corpid, agentid)
	a = a.WithSecret(secret)

	token := os.Getenv("WORKWX_RECV_TOKEN")
	encodingAESKey := os.Getenv("WORKWX_RECV_AES_KEY")
	a.SetMsgCrypt(token, encodingAESKey)

	return a
}
