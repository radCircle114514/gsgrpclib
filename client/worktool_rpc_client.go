/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"context"
	"log"
	"reflect"

	worktool_rpc "github.com/iznilul/gsgrpclib/proto/worktool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func InitWorkToolRpcClientConn() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if Config.ServerConfig.Tls {
		creds, err := credentials.NewClientTLSFromFile(Config.ServerConfig.CaFile, Config.ServerConfig.ServerNameOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(Config.ServerConfig.RpcWorkToolHost, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func InvokeWorkToolRPCMethod(ctx context.Context, methodName string, ao *worktool_rpc.RequestAO) (*worktool_rpc.ResponseVO, error) {
	conn, err := InitWorkToolRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := worktool_rpc.NewWorkToolRPCClient(conn)

	ctx, cancel := SetTimeout(ctx)
	defer cancel()

	method := reflect.ValueOf(client).MethodByName(methodName)
	if !method.IsValid() {
		return nil, errors.New("method not found")
	}

	args := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(ao)}
	result := method.Call(args)

	if result[0].IsValid() {
		vo, ok := result[0].Interface().(*worktool_rpc.ResponseVO)
		if !ok {
			return nil, errors.New("invalid response")
		}
		if result[1].IsValid() {
			err, _ := result[1].Interface().(error)
			if err != nil {
				return nil, errors.Wrap(err, "error")
			}
		}
		return vo, nil
	}
	return nil, errors.New("invalid response")
}
