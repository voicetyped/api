package transcriberv1

import (
	"context"
	apic "github.com/antinvestor/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math"
)

const ctxKeyService = apic.CtxServiceKey("transcriberClientKey")

func defaultTranscriberClientOptions() []apic.ClientOption {
	return []apic.ClientOption{
		apic.WithEndpoint("transcribe.api.antinvestor.com:443"),
		apic.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		apic.WithGRPCDialOption(grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func ToContext(ctx context.Context, client *TranscribeClient) context.Context {
	return context.WithValue(ctx, ctxKeyService, client)
}

func FromContext(ctx context.Context) *TranscribeClient {
	client, ok := ctx.Value(ctxKeyService).(*TranscribeClient)
	if !ok {
		return nil
	}

	return client
}

// TranscribeClient is a client for interacting with the transcribe service API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type TranscribeClient struct {
	// gRPC connection to the service.
	clientConn *grpc.ClientConn

	// The gRPC API client.
	client TranscriberServiceClient

	// The x-ant-* metadata to be sent with each request.
	xMetadata metadata.MD
}

// InstantiateNotificationClient creates a new notification client.
//
// The service that an application uses to send and access received messages
func InstantiateNotificationClient(clientConnection *grpc.ClientConn, transcribeServiceCli TranscriberServiceClient) *TranscribeClient {
	c := &TranscribeClient{
		clientConn: clientConnection,
		client:     transcribeServiceCli,
	}

	c.setClientInfo()

	return c
}

// NewNotificationClient creates a new notification client.
//
// The service that an application uses to send and access received messages
func NewNotificationClient(ctx context.Context, opts ...apic.ClientOption) (*TranscribeClient, error) {
	clientOpts := defaultTranscriberClientOptions()

	connPool, err := apic.DialConnection(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}

	transcribeServiceCli := NewTranscriberServiceClient(connPool)
	return InstantiateNotificationClient(connPool, transcribeServiceCli), nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (nc *TranscribeClient) Close() error {
	return nc.clientConn.Close()
}

// Service creates a new notification service for use to invoke.
func (nc *TranscribeClient) Service() TranscriberServiceClient {

	if nc.client != nil {
		return nc.client
	}

	return NewTranscriberServiceClient(nc.clientConn)
}

// setClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (nc *TranscribeClient) setClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", apic.VersionGo()}, keyval...)
	kv = append(kv, "grpc", grpc.Version)
	nc.xMetadata = metadata.Pairs("x-ai-api-client", apic.XAntHeader(kv...))
}

func (nc *TranscribeClient) Transcribe(ctx context.Context) (TranscriberService_TranscribeClient, error) {
	return nc.Service().Transcribe(ctx)
}
