package grpc

import (
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/contextcomponent"
	"github.com/Alp4ka/classifier-aaS/pkg/api"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	uuid "github.com/google/uuid"
	"io"
	"sync"
)

func (s *Server) Process(src api.GWManagerService_ProcessServer) (err error) {
	ctx := src.Context()
	defer func() {
		if err != nil {
			mlogger.L(ctx).Error("Process failed", field.Error(err))
		}
	}()

	var (
		sessionIDOnce sync.Once
		sessionID     uuid.UUID
	)
	for {
		req, err := src.Recv()
		if err == io.EOF {
			mlogger.L(ctx).Info("Client closed connection")
			break
		} else if err != nil {
			return fmt.Errorf("failed to receive client request: %w", err)
		}
		ctx = field.WithContextFields(ctx, field.String("id", req.GetSessionId()), field.String("req", req.GetRequestData()))
		mlogger.L(ctx).Info("Handle client request")

		// Session ID.
		sessID, err := uuid.Parse(req.GetSessionId())
		if err != nil {
			return fmt.Errorf("unable to parse session uuid: %w", err)
		}
		sessionIDOnce.Do(func() { sessionID = sessID })
		if sessionID != sessID {
			return fmt.Errorf("session id mismatch")
		}

		// Getting session.
		sess, err := s.contextService.GetSession(ctx, &contextcomponent.GetSessionParams{SessionID: sessionID})
		if err != nil {
			return fmt.Errorf("failed to acquire session: %w", err)
		}

		// Handle request.

		ret := &api.ProcessResponse{}
		*ret.ResponseData = sess.Model.ID.String()
		err = src.Send(ret)
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}
	}

	return nil
}