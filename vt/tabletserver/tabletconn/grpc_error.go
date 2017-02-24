package tabletconn

import (
	"io"

	"google.golang.org/grpc"
	"gopkg.in/sqle/vitess-go.v1/vt/vterrors"

	vtrpcpb "gopkg.in/sqle/vitess-go.v1/vt/proto/vtrpc"
)

// ErrorFromGRPC converts a GRPC error to vtError for
// tabletserver calls.
func ErrorFromGRPC(err error) error {
	// io.EOF is end of stream. Don't treat it as an error.
	if err == nil || err == io.EOF {
		return nil
	}
	return vterrors.New(vtrpcpb.Code(grpc.Code(err)), "vttablet: "+err.Error())
}

// ErrorFromVTRPC converts a *vtrpcpb.RPCError to vtError for
// tabletserver calls.
func ErrorFromVTRPC(err *vtrpcpb.RPCError) error {
	if err == nil {
		return nil
	}
	code := err.Code
	if code == vtrpcpb.Code_OK {
		code = vterrors.LegacyErrorCodeToCode(err.LegacyCode)
	}
	return vterrors.New(code, "vttablet: "+err.Message)
}
