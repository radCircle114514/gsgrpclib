package client

import (
	"context"
	"log"
	"reflect"

	workbuddy_rpc "github.com/iznilul/gsgrpclib/proto/workbuddy"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func InitWorkBuddyRpcClientConn() (*grpc.ClientConn, error) {
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

	conn, err := grpc.NewClient(Config.ServerConfig.RpcWorkBuddyHost, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func InvokeWorkBuddyRPCMethod(ctx context.Context, methodName string, ao *workbuddy_rpc.RequestAO) (*workbuddy_rpc.ResponseVO, error) {
	conn, err := InitWorkBuddyRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := workbuddy_rpc.NewWorkBuddyRPCClient(conn)

	ctx, cancel := SetTimeout(ctx)
	defer cancel()

	method := reflect.ValueOf(client).MethodByName(methodName)
	if !method.IsValid() {
		return nil, errors.New("method not found")
	}

	args := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(ao)}
	result := method.Call(args)

	if result[0].IsValid() {
		vo, ok := result[0].Interface().(*workbuddy_rpc.ResponseVO)
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
