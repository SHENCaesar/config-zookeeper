// Copyright 2023 CloudWeGo Authors
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

package server

import (
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/config-zookeeper/utils"
	"github.com/kitex-contrib/config-zookeeper/zookeeper"
)

const (
	limiterConfigName = "limit"
)

// ZookeeperServerSuite  zookeeper server config suite, configure limiter config dynamically from zookeeper .
type ZookeeperServerSuite struct {
	zookeeperClient zookeeper.Client
	service         string
	opts            utils.Options
}

// NewSuite service is the destination service.
func NewSuite(service string, cli zookeeper.Client, opts ...utils.Option) *ZookeeperServerSuite {
	su := &ZookeeperServerSuite{
		service:         service,
		zookeeperClient: cli,
	}
	for _, opt := range opts {
		opt.Apply(&su.opts)
	}
	return su
}

// Options return a list server.Option
func (s *ZookeeperServerSuite) Options() []server.Option {
	opts := make([]server.Option, 0, 2)
	opts = append(opts, WithLimiter(s.service, s.zookeeperClient, s.opts))
	return opts
}
