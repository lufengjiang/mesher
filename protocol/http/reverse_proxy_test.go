/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package http_test

import (
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/go-chassis/mesher/cmd"
	"github.com/go-chassis/mesher/protocol/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetLocalServiceAddress(t *testing.T) {
	cmd.Configs = &cmd.ConfigFromCmd{
		LocalServicePorts: "",
	}
	err := cmd.Configs.GeneratePortsMap()

	inv := invocation.New(nil)
	inv.Protocol = "rest"
	err = http.SetLocalServiceAddress(inv, "8080")
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1:8080", inv.Endpoint)

	err = http.SetLocalServiceAddress(inv, "")
	assert.Error(t, err)

	cmd.Configs.LocalServicePorts = "rest:80"
	err = cmd.Configs.GeneratePortsMap()
	t.Log(cmd.Configs.PortsMap)
	assert.NoError(t, err)
	err = http.SetLocalServiceAddress(inv, "")
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1:80", inv.Endpoint)
}
