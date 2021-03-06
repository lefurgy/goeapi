//
// Copyright (c) 2015, Arista Networks, Inc.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright notice,
//   this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//
//   * Neither the name of Arista Networks nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL ARISTA NETWORKS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
// BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
// OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN
// IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package goeapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// EapiConnectionEntity is an interface representing the ability to execute a
// single json transaction, obtaining the Response for a given Request.
type EapiConnectionEntity interface {
	Execute(commands []interface{}, encoding string) (*JSONRPCResponse, error)
	Error() error
}

// EapiConnection represents the base object for implementing an EapiConnection
// type. This clase should not be instantiated directly
type EapiConnection struct {
	transport   string
	err         error
	socketError error
	url         string
	host        string
	port        int
	path        string
	auth        string
}

// Execute the list of commands on the destination node. In the case of
// EapiConnection, this serves as a base model and is not fully implemented.
//
// Args:
//  commands ([]interface): list of commands to execute on remote node
//  encoding (string): The encoding to send along with the request
//                      message to the destination node.  Valid values include
//                      'json' or 'text'.  This argument will influence the
//                      response encoding
// Returns:
//  pointer to JSONRPCResponse or error on failure
func (conn *EapiConnection) Execute(commands []interface{},
	encoding string) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No connection")
	}
	return &JSONRPCResponse{}, fmt.Errorf("Not Currently Implemented")
}

// Authentication Configures the user authentication for eAPI. This method
// configures the username and password combination to use for authenticating
// to eAPI.
//
// Args:
//  username (string): The username to use to authenticate the eAPI
//                      connection with
//  password (string): The password in clear text to use to authenticate
//                      the eAPI connection with
func (conn *EapiConnection) Authentication(username string, passwd string) {
	auth := username + ":" + passwd
	//data := []byte(auth)
	//str := base64.StdEncoding.EncodeToString(data)
	//str = strings.Replace(str,"\n","",-1)
	str := strings.Replace(auth, "\n", "", -1)
	conn.auth = str
}

// getURL helper method to prebuild url for http/s
func (conn *EapiConnection) getURL() string {
	if conn == nil {
		return ""
	}
	url := conn.transport + "://" + conn.auth + "@" + conn.host + "/command-api/"
	return url
}

// Error returns the current error for Connection
func (conn *EapiConnection) Error() error {
	if conn == nil {
		return nil
	}
	return conn.err
}

// SetError sets error for Connection
func (conn *EapiConnection) SetError(e error) {
	if conn == nil {
		return
	}
	conn.err = e
}

// ClearError clears any error for Connection
func (conn *EapiConnection) ClearError() {
	if conn == nil {
		return
	}
	conn.err = nil
}

// buildJSONRequest builds a JSON request given a list of commands, encoding
// type of either json or text, and request id. The command list input is made
// up of a list of interface{} types. This is so associative entries and list
// entries both can be used. Returns []byte of the built JSON request.
// Successful call returns err == nil.
func buildJSONRequest(commands []interface{},
	encoding string, reqid int) ([]byte, error) {
	p := Parameters{1, commands, encoding}

	req := Request{"2.0", "runCmds", p, reqid}
	data, err := json.Marshal(req)
	//debugJSON(data)
	return data, err
}

// SocketEapiConnection represents the EapiConnection for handling Socket
// level transactions
type SocketEapiConnection struct {
	EapiConnection
}

//
const defaultUnixSocket = "/var/run/command-api.sock"

// NewSocketEapiConnection initializes a SocketEapiConnection.
//
// Args:
//  transport (string): The transport to use to create the instance.
//  host (string): The IP addres or DNS host name of the connection device.
//  username(string): The username to pass to the device to authenticate
//                    the eAPI connection.
//  password(string): The password to pass to the device to authenticate
//                    the eAPI connection. The default value is ''
//  port(int): The TCP port of the endpoint for the eAPI connection.
//
// Returns:
//  Newly created SocketEapiConnection
func NewSocketEapiConnection(transport string, host string, username string,
	password string, port int) EapiConnectionEntity {
	conn := EapiConnection{transport: transport, host: host, port: port}
	return &SocketEapiConnection{conn}
}

// send the eAPI request to the destination node
//
// This method is responsible for sending an eAPI request to the
// destination node and returning a response based on the JSONRPCResponse
// object.  eAPI responds to request messages with either a success
// message or failure message. On successful decode of the Response,
// a JSONRPCResponse type is returned. Otherwise err is returned.
func (conn *SocketEapiConnection) send(data []byte) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No Connection")
	}
	return &JSONRPCResponse{}, fmt.Errorf("Not Currently Implemented")
}

// Execute the list of commands on the destination node
//
// This method takes a list of commands and sends them to the
// destination node, returning the results. It is assumed that the
// list of commands (type []interface{}) has been properly built and
// enable mode passwd is set if needed. On success, a reference
// to JSONRPCResponse is returned...otherwise err is set.
//
// Args:
//  commands ([]interface): list of commands to execute on remote node
//  encoding (string): The encoding to send along with the request
//                      message to the destination node.  Valid values include
//                      'json' or 'text'.  This argument will influence the
//                      response encoding
// Returns:
//  pointer to JSONRPCResponse or error on failure
func (conn *SocketEapiConnection) Execute(commands []interface{},
	encoding string) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No connection")
	}
	conn.ClearError()
	data, err := buildJSONRequest(commands, encoding, os.Getpid())
	if err != nil {
		conn.SetError(err)
		return &JSONRPCResponse{}, err
	}
	return conn.send(data)
}

// HTTPLocalEapiConnection is an EapiConnection suited for local HTTP connection
type HTTPLocalEapiConnection struct {
	EapiConnection
}

// UseDefaultPortNum recommends the underlying api to use default value for
// Port Number.
const UseDefaultPortNum = -1

// DefaultHTTPLocalPort uses 8080
const DefaultHTTPLocalPort = 8080

// NewHTTPLocalEapiConnection initializes a HTTPLocalEapiConnection.
//
// Args:
//  transport (string): The transport to use to create the instance.
//  host (string): The IP addres or DNS host name of the connection device.
//  username(string): The username to pass to the device to authenticate
//                    the eAPI connection.
//  password(string): The password to pass to the device to authenticate
//                    the eAPI connection. The default value is ''
//  port(int): The TCP port of the endpoint for the eAPI connection.
//
// Returns:
//  Newly created SocketEapiConnection
func NewHTTPLocalEapiConnection(transport string, host string, username string,
	password string, port int) EapiConnectionEntity {
	conn := EapiConnection{transport: transport, host: host, port: port}
	return &HTTPLocalEapiConnection{conn}
}

// send the eAPI request to the destination node
//
// This method is responsible for sending an eAPI request to the
// destination node and returning a response based on the JSONRPCResponse
// object.  eAPI responds to request messages with either a success
// message or failure message. On successful decode of the Response,
// a JSONRPCResponse type is returned. Otherwise err is returned.
//
// Args:
//  data ([]byte): data to be sent
// Returns:
//  ptr to JSONRPCResponse on success. Otherwise error will be returned.
func (conn *HTTPLocalEapiConnection) send(data []byte) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No Connection")
	}
	return &JSONRPCResponse{}, fmt.Errorf("Not Currently Implemented")
}

// Execute the list of commands
//
// This method takes a list of commands and sends them to the
// destination node, returning the results. It is assumed that the
// list of commands (type []interface{}) has been properly built and
// enable mode passwd is set if needed. On success, a reference
// to JSONRPCResponse is returned...otherwise err is set.
//
// Args:
//  commands ([]interface): list of commands to execute on remote node
//  encoding (string): The encoding to send along with the request
//                      message to the destination node.  Valid values include
//                      'json' or 'text'.  This argument will influence the
//                      response encoding
// Returns:
//  pointer to JSONRPCResponse or error on failure
func (conn *HTTPLocalEapiConnection) Execute(commands []interface{},
	encoding string) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No connection")
	}
	conn.ClearError()
	data, err := buildJSONRequest(commands, encoding, os.Getpid())
	if err != nil {
		conn.SetError(err)
		return &JSONRPCResponse{}, err
	}
	return conn.send(data)
}

// HTTPEapiConnection is an EapiConnection suited for HTTP connection
type HTTPEapiConnection struct {
	EapiConnection
}

const defaultHTTPPort = 80

// NewHTTPEapiConnection initializes a HttpEapiConnection.
//
// Args:
//  transport (string): The transport to use to create the instance.
//  host (string): The IP addres or DNS host name of the connection device.
//  username(string): The username to pass to the device to authenticate
//                    the eAPI connection.
//  password(string): The password to pass to the device to authenticate
//                    the eAPI connection. The default value is ''
//  port(int): The TCP port of the endpoint for the eAPI connection.
//
// Returns:
//  Newly created HTTPEapiConnection
func NewHTTPEapiConnection(transport string, host string, username string,
	password string, port int) EapiConnectionEntity {
	port = defaultHTTPPort
	conn := EapiConnection{transport: transport, host: host, port: port}
	conn.Authentication(username, password)
	return &HTTPEapiConnection{conn}
}

// send the eAPI request to the destination node
//
// This method is responsible for sending an eAPI request to the
// destination node and returning a response based on the JSONRPCResponse
// object.  eAPI responds to request messages with either a success
// message or failure message. On successful decode of the Response,
// a JSONRPCResponse type is returned. Otherwise err is returned.
//
// Args:
//  data ([]byte): data to be sent
// Returns:
//  ptr to JSONRPCResponse on success. Otherwise error will be returned.
func (conn *HTTPEapiConnection) send(data []byte) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No Connection")
	}

	client := &http.Client{}
	url := conn.getURL()
	resp, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		conn.SetError(err)
		return &JSONRPCResponse{}, err
	}
	defer resp.Body.Close()

	jsonRsp := decodeEapiResponse(resp)

	// check for errors in the JSON response
	if jsonRsp.Error != nil {
		err := fmt.Errorf("JSON Error(%d): %s", jsonRsp.Error.Code,
			jsonRsp.Error.Message)
		conn.SetError(err)
		return jsonRsp, err
	}
	return jsonRsp, nil
}

// Execute the list of commands on the destination node
//
// This method takes a list of commands and sends them to the
// destination node, returning the results. It is assumed that the
// list of commands (type []interface{}) has been properly built and
// enable mode passwd is set if needed. On success, a reference
// to JSONRPCResponse is returned...otherwise err is set.
//
// Args:
//  commands ([]interface): list of commands to execute on remote node
//  encoding (string): The encoding to send along with the request
//                      message to the destination node.  Valid values include
//                      'json' or 'text'.  This argument will influence the
//                      response encoding
// Returns:
//  pointer to JSONRPCResponse or error on failure
func (conn *HTTPEapiConnection) Execute(commands []interface{},
	encoding string) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No connection")
	}
	conn.ClearError()
	data, err := buildJSONRequest(commands, encoding, os.Getpid())
	if err != nil {
		conn.SetError(err)
		return &JSONRPCResponse{}, err
	}
	return conn.send(data)
}

// HTTPSEapiConnection is an EapiConnection suited for HTTP connection
type HTTPSEapiConnection struct {
	EapiConnection
	path                string
	enforceVerification bool
}

// DefaultHTTPSPort default port used by https
const DefaultHTTPSPort = 443

// DefaultHTTPSPath command path
const DefaultHTTPSPath = "/command-api"

// NewHTTPSEapiConnection initializes an HttpsEapiConnection.
//
// Args:
//  transport (string): The transport to use to create the instance.
//  host (string): The IP addres or DNS host name of the connection device.
//  username(string): The username to pass to the device to authenticate
//                    the eAPI connection.
//  password(string): The password to pass to the device to authenticate
//                    the eAPI connection. The default value is ''
//  port(int): The TCP port of the endpoint for the eAPI connection.
//
// Returns:
//  Newly created HTTPSEapiConnection
func NewHTTPSEapiConnection(transport string, host string, username string,
	password string, port int) EapiConnectionEntity {
	port = DefaultHTTPSPort
	path := DefaultHTTPSPath

	conn := EapiConnection{transport: transport, host: host, port: port}

	conn.Authentication(username, password)
	return &HTTPSEapiConnection{path: path, EapiConnection: conn}
}

// send the eAPI request to the destination node
//
// This method is responsible for sending an eAPI request to the
// destination node and returning a response based on the JSONRPCResponse
// object.  eAPI responds to request messages with either a success
// message or failure message. On successful decode of the Response,
// a JSONRPCResponse type is returned. Otherwise err is returned.
//
// Args:
//  data ([]byte): data to be sent
// Returns:
//  ptr to JSONRPCResponse on success. Otherwise error will be returned.
func (conn *HTTPSEapiConnection) send(data []byte) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No Connection")
	}
	client := &http.Client{}
	url := conn.getURL()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}

	resp, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		conn.SetError(err)
		return &JSONRPCResponse{}, err
	}
	defer resp.Body.Close()

	jsonRsp := decodeEapiResponse(resp)

	// check for errors in the JSON response
	if jsonRsp.Error != nil {
		err := fmt.Errorf("JSON Error(%d): %s", jsonRsp.Error.Code,
			jsonRsp.Error.Message)
		conn.SetError(err)
		return jsonRsp, err
	}
	return jsonRsp, nil
}

// Execute the list of commands on the destination node
//
// This method takes a list of commands and sends them to the
// destination node, returning the results. It is assumed that the
// list of commands (type []interface{}) has been properly built and
// enable mode passwd is set if needed. On success, a reference
// to JSONRPCResponse is returned...otherwise err is set.
//
// Args:
//  commands ([]interface): list of commands to execute on remote node
//  encoding (string): The encoding to send along with the request
//                      message to the destination node.  Valid values include
//                      'json' or 'text'.  This argument will influence the
//                      response encoding
// Returns:
//  pointer to JSONRPCResponse or error on failure
func (conn *HTTPSEapiConnection) Execute(commands []interface{},
	encoding string) (*JSONRPCResponse, error) {
	if conn == nil {
		return &JSONRPCResponse{}, fmt.Errorf("No connection")
	}
	conn.ClearError()
	data, err := buildJSONRequest(commands, encoding, os.Getpid())
	if err != nil {
		conn.SetError(err)
		return &JSONRPCResponse{}, err
	}
	return conn.send(data)
}

// disableCertificateVerification disables https verification
func (conn *HTTPSEapiConnection) disableCertificateVerification() {
	conn.enforceVerification = false
}

// enableCertificateVerification enables https verification
func (conn *HTTPSEapiConnection) enableCertificateVerification() {
	conn.enforceVerification = true
}
