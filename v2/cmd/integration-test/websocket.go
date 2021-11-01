package main

import (
	"net"
	"strings"

	"github.com/gobwas/ws/wsutil"
	"github.com/projectdiscovery/nuclei/v2/internal/testutils"
)

var websocketTestCases = map[string]testutils.TestCase{
	"websocket/basic.yaml":    &websocketBasic{},
	"websocket/cswsh.yaml":    &websocketCswsh{},
	"websocket/no-cswsh.yaml": &websocketNoCswsh{},
}

type websocketBasic struct{}

// Execute executes a test case and returns an error if occurred
func (h *websocketBasic) Execute(filePath string) error {
	connHandler := func(conn net.Conn) {
		for {
			msg, op, _ := wsutil.ReadClientData(conn)
			if string(msg) != string("hello") {
				return
			}
			_ = wsutil.WriteServerMessage(conn, op, []byte("world"))
		}
	}
	originValidate := func(origin string) bool {
		return true
	}
	ts := testutils.NewWebsocketServer(connHandler, originValidate)
	defer ts.Close()

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, strings.ReplaceAll(ts.URL, "http", "ws"), debug)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return errIncorrectResultsCount(results)
	}
	return nil
}

type websocketCswsh struct{}

// Execute executes a test case and returns an error if occurred
func (h *websocketCswsh) Execute(filePath string) error {
	connHandler := func(conn net.Conn) {

	}
	originValidate := func(origin string) bool {
		return true
	}
	ts := testutils.NewWebsocketServer(connHandler, originValidate)
	defer ts.Close()

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, strings.ReplaceAll(ts.URL, "http", "ws"), debug)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return errIncorrectResultsCount(results)
	}
	return nil
}

type websocketNoCswsh struct{}

// Execute executes a test case and returns an error if occurred
func (h *websocketNoCswsh) Execute(filePath string) error {
	connHandler := func(conn net.Conn) {

	}
	originValidate := func(origin string) bool {
		return origin == "https://google.com"
	}
	ts := testutils.NewWebsocketServer(connHandler, originValidate)
	defer ts.Close()

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, strings.ReplaceAll(ts.URL, "http", "ws"), debug)
	if err != nil {
		return err
	}
	if len(results) != 0 {
		return errIncorrectResultsCount(results)
	}
	return nil
}
