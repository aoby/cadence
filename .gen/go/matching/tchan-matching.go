// @generated Code generated by thrift-gen. Do not modify.

// Package matching is generated code used to make or handle TChannel calls using Thrift.
package matching

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"

	"code.uber.internal/devexp/minions/.gen/go/shared"
)

var _ = shared.GoUnusedProtection__

// Interfaces for the service and client for the services defined in the IDL.

// TChanMatchingService is the interface that defines the server handler and client interface.
type TChanMatchingService interface {
	AddActivityTask(ctx thrift.Context, addRequest *AddActivityTaskRequest) error
	AddDecisionTask(ctx thrift.Context, addRequest *AddDecisionTaskRequest) error
	PollForActivityTask(ctx thrift.Context, pollRequest *shared.PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse, error)
	PollForDecisionTask(ctx thrift.Context, pollRequest *shared.PollForDecisionTaskRequest) (*shared.PollForDecisionTaskResponse, error)
}

// Implementation of a client and service handler.

type tchanMatchingServiceClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanMatchingServiceInheritedClient(thriftService string, client thrift.TChanClient) *tchanMatchingServiceClient {
	return &tchanMatchingServiceClient{
		thriftService,
		client,
	}
}

// NewTChanMatchingServiceClient creates a client that can be used to make remote calls.
func NewTChanMatchingServiceClient(client thrift.TChanClient) TChanMatchingService {
	return NewTChanMatchingServiceInheritedClient("MatchingService", client)
}

func (c *tchanMatchingServiceClient) AddActivityTask(ctx thrift.Context, addRequest *AddActivityTaskRequest) error {
	var resp MatchingServiceAddActivityTaskResult
	args := MatchingServiceAddActivityTaskArgs{
		AddRequest: addRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "AddActivityTask", &args, &resp)
	if err == nil && !success {
		if e := resp.BadRequestError; e != nil {
			err = e
		}
		if e := resp.InternalServiceError; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanMatchingServiceClient) AddDecisionTask(ctx thrift.Context, addRequest *AddDecisionTaskRequest) error {
	var resp MatchingServiceAddDecisionTaskResult
	args := MatchingServiceAddDecisionTaskArgs{
		AddRequest: addRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "AddDecisionTask", &args, &resp)
	if err == nil && !success {
		if e := resp.BadRequestError; e != nil {
			err = e
		}
		if e := resp.InternalServiceError; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanMatchingServiceClient) PollForActivityTask(ctx thrift.Context, pollRequest *shared.PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse, error) {
	var resp MatchingServicePollForActivityTaskResult
	args := MatchingServicePollForActivityTaskArgs{
		PollRequest: pollRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "PollForActivityTask", &args, &resp)
	if err == nil && !success {
		if e := resp.BadRequestError; e != nil {
			err = e
		}
		if e := resp.InternalServiceError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanMatchingServiceClient) PollForDecisionTask(ctx thrift.Context, pollRequest *shared.PollForDecisionTaskRequest) (*shared.PollForDecisionTaskResponse, error) {
	var resp MatchingServicePollForDecisionTaskResult
	args := MatchingServicePollForDecisionTaskArgs{
		PollRequest: pollRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "PollForDecisionTask", &args, &resp)
	if err == nil && !success {
		if e := resp.BadRequestError; e != nil {
			err = e
		}
		if e := resp.InternalServiceError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

type tchanMatchingServiceServer struct {
	handler TChanMatchingService
}

// NewTChanMatchingServiceServer wraps a handler for TChanMatchingService so it can be
// registered with a thrift.Server.
func NewTChanMatchingServiceServer(handler TChanMatchingService) thrift.TChanServer {
	return &tchanMatchingServiceServer{
		handler,
	}
}

func (s *tchanMatchingServiceServer) Service() string {
	return "MatchingService"
}

func (s *tchanMatchingServiceServer) Methods() []string {
	return []string{
		"AddActivityTask",
		"AddDecisionTask",
		"PollForActivityTask",
		"PollForDecisionTask",
	}
}

func (s *tchanMatchingServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "AddActivityTask":
		return s.handleAddActivityTask(ctx, protocol)
	case "AddDecisionTask":
		return s.handleAddDecisionTask(ctx, protocol)
	case "PollForActivityTask":
		return s.handlePollForActivityTask(ctx, protocol)
	case "PollForDecisionTask":
		return s.handlePollForDecisionTask(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanMatchingServiceServer) handleAddActivityTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServiceAddActivityTaskArgs
	var res MatchingServiceAddActivityTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.AddActivityTask(ctx, req.AddRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanMatchingServiceServer) handleAddDecisionTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServiceAddDecisionTaskArgs
	var res MatchingServiceAddDecisionTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.AddDecisionTask(ctx, req.AddRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanMatchingServiceServer) handlePollForActivityTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServicePollForActivityTaskArgs
	var res MatchingServicePollForActivityTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.PollForActivityTask(ctx, req.PollRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanMatchingServiceServer) handlePollForDecisionTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServicePollForDecisionTaskArgs
	var res MatchingServicePollForDecisionTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.PollForDecisionTask(ctx, req.PollRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}