package lndclient

import (
	context "context"
	"github.com/lightningnetwork/lnd/lnrpc/neutrinorpc"
	grpc "google.golang.org/grpc"
	"time"
)

type NeutrinoKitClient interface {
	//
	//Status returns the status of the light client neutrino instance,
	//along with height and hash of the best block, and a list of connected
	//peers.
	Status(ctx context.Context, in *neutrinorpc.StatusRequest, opts ...grpc.CallOption) (*neutrinorpc.StatusResponse, error)
	//
	//AddPeer adds a new peer that has already been connected to the server.
	AddPeer(ctx context.Context, in *neutrinorpc.AddPeerRequest, opts ...grpc.CallOption) (*neutrinorpc.AddPeerResponse, error)
	//
	//DisconnectPeer disconnects a peer by target address. Both outbound and
	//inbound nodes will be searched for the target node. An error message will
	//be returned if the peer was not found.
	DisconnectPeer(ctx context.Context, in *neutrinorpc.DisconnectPeerRequest, opts ...grpc.CallOption) (*neutrinorpc.DisconnectPeerResponse, error)
	//
	//IsBanned returns true if the peer is banned, otherwise false.
	IsBanned(ctx context.Context, in *neutrinorpc.IsBannedRequest, opts ...grpc.CallOption) (*neutrinorpc.IsBannedResponse, error)
	//
	//GetBlockHeader returns a block header with a particular block hash.
	GetBlockHeader(ctx context.Context, in *neutrinorpc.GetBlockHeaderRequest, opts ...grpc.CallOption) (*neutrinorpc.GetBlockHeaderResponse, error)
	//
	//GetBlock returns a block with a particular block hash.
	GetBlock(ctx context.Context, in *neutrinorpc.GetBlockRequest, opts ...grpc.CallOption) (*neutrinorpc.GetBlockResponse, error)
	//
	//GetCFilter returns a compact filter from a block.
	GetCFilter(ctx context.Context, in *neutrinorpc.GetCFilterRequest, opts ...grpc.CallOption) (*neutrinorpc.GetCFilterResponse, error)
	//
	//GetBlockHash returns the header hash of a block at a given height.
	GetBlockHash(ctx context.Context, in *neutrinorpc.GetBlockHashRequest, opts ...grpc.CallOption) (*neutrinorpc.GetBlockHashResponse, error)
}

type NeutrinoClient struct {
	client      NeutrinoKitClient
	readonlyMac serializedMacaroon
	timeout     time.Duration
}

func newNeutrinoClient(conn grpc.ClientConnInterface,
	readonlyMac serializedMacaroon, timeout time.Duration) *NeutrinoClient {

	return &NeutrinoClient{
		client:      neutrinorpc.NewNeutrinoKitClient(conn),
		readonlyMac: readonlyMac,
		timeout:     timeout,
	}
}

type neutrinoKitClient struct {
	cc grpc.ClientConnInterface
}

func NewNeutrinoKitClient(cc grpc.ClientConnInterface) NeutrinoKitClient {
	return &neutrinoKitClient{cc}
}

func (c *neutrinoKitClient) Status(ctx context.Context, in *neutrinorpc.StatusRequest, opts ...grpc.CallOption) (*neutrinorpc.StatusResponse, error) {
	out := new(neutrinorpc.StatusResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) AddPeer(ctx context.Context, in *neutrinorpc.AddPeerRequest, opts ...grpc.CallOption) (*neutrinorpc.AddPeerResponse, error) {
	out := new(neutrinorpc.AddPeerResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) DisconnectPeer(ctx context.Context, in *neutrinorpc.DisconnectPeerRequest, opts ...grpc.CallOption) (*neutrinorpc.DisconnectPeerResponse, error) {
	out := new(neutrinorpc.DisconnectPeerResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/DisconnectPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) IsBanned(ctx context.Context, in *neutrinorpc.IsBannedRequest, opts ...grpc.CallOption) (*neutrinorpc.IsBannedResponse, error) {
	out := new(neutrinorpc.IsBannedResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/IsBanned", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) GetBlockHeader(ctx context.Context, in *neutrinorpc.GetBlockHeaderRequest, opts ...grpc.CallOption) (*neutrinorpc.GetBlockHeaderResponse, error) {
	out := new(neutrinorpc.GetBlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/GetBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) GetBlock(ctx context.Context, in *neutrinorpc.GetBlockRequest, opts ...grpc.CallOption) (*neutrinorpc.GetBlockResponse, error) {
	out := new(neutrinorpc.GetBlockResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) GetCFilter(ctx context.Context, in *neutrinorpc.GetCFilterRequest, opts ...grpc.CallOption) (*neutrinorpc.GetCFilterResponse, error) {
	out := new(neutrinorpc.GetCFilterResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/GetCFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *neutrinoKitClient) GetBlockHash(ctx context.Context, in *neutrinorpc.GetBlockHashRequest, opts ...grpc.CallOption) (*neutrinorpc.GetBlockHashResponse, error) {
	out := new(neutrinorpc.GetBlockHashResponse)
	err := c.cc.Invoke(ctx, "/neutrinorpc.NeutrinoKit/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
