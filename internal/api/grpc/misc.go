package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

// Gets the key value from the provided context, or returns an empty string
func getSessionID(ctx context.Context) (uuid.UUID, error) {
	const key = "session-id"

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("unable to get metadata from context")
	}

	header, ok := md[key]
	if !ok || len(header) == 0 {
		return uuid.UUID{}, fmt.Errorf("unable to get session-id from metadata")
	}

	result, err := uuid.Parse(header[0])
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("session-id was not a UUID")
	}

	return result, nil
}
