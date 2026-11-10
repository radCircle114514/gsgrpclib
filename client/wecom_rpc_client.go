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

// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.booking.
package client

import (
	"context"
	"github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"reflect"
)

//func customInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	// Add a custom value to the context
//	ctx = metadata.AppendToOutgoingContext(ctx, "from", "kpi")
//	// Perform any other logic before invoking the RPC method
//	//global.Info.Info("Added custom value to context before calling method %s", method)
//	// Call the RPC method with the updated context
//	return invoker(ctx, method, req, reply, cc, opts...)
//}

func InitWecomRpcClientConn() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// WeDrive file uploads carry base64 payloads; allow messages up to 32MB
	// instead of the gRPC default 4MB limit.
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(32<<20), grpc.MaxCallSendMsgSize(32<<20)))
	if Config.ServerConfig.Tls {
		creds, err := credentials.NewClientTLSFromFile(Config.ServerConfig.CaFile, Config.ServerConfig.ServerNameOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	//opts = append(opts, grpc.WithUnaryInterceptor(customInterceptor))

	conn, err := grpc.NewClient(Config.ServerConfig.RpcWecomHost, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func InvokeWecomRPCMethod(ctx context.Context, methodName string, ao *wecom_rpc.RequestAO) (*wecom_rpc.ResponseVO, error) {
	conn, err := InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := wecom_rpc.NewWecomRPCClient(conn)

	ctx, cancel := SetTimeout(ctx)
	defer cancel()

	method := reflect.ValueOf(client).MethodByName(methodName)
	if !method.IsValid() {
		return nil, errors.New("method not found")
	}

	args := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(ao)}
	result := method.Call(args)

	if result[0].IsValid() {
		vo, ok := result[0].Interface().(*wecom_rpc.ResponseVO)
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
