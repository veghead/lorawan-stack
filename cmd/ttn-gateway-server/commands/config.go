// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"github.com/TheThingsNetwork/ttn/cmd/internal/shared"
	shared_gatewayserver "github.com/TheThingsNetwork/ttn/cmd/internal/shared/gatewayserver"
	conf "github.com/TheThingsNetwork/ttn/pkg/config"
	"github.com/TheThingsNetwork/ttn/pkg/gatewayserver"
)

// Config of the gateway server to start.
type Config struct {
	conf.ServiceBase `name:",squash"`
	GS               gatewayserver.Config `name:"gs"`
}

// DefaultConfig of the gateway server.
var DefaultConfig = Config{
	ServiceBase: shared.DefaultServiceBase,
	GS:          shared_gatewayserver.DefaultGatewayServerConfig,
}
