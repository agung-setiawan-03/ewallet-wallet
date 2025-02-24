package external

import (
	"context"
	"ewallet-wallet/constants"
	"ewallet-wallet/external/proto/tokenvalidation"
	"ewallet-wallet/internal/models"
	"fmt"

	"google.golang.org/grpc"
	"github.com/pkg/errors"
)

type External struct {

}

func (*External) ValidateToken(ctx context.Context, token string) (models.TokenData, error) {
	var (
		resp models.TokenData
	)

	conn, err := grpc.Dial("localhost:7000", grpc.WithInsecure())
	if err != nil {
		return resp, errors.Wrap(err, "Failed to dial ums grpc")
	}
	defer conn.Close()

	client := tokenvalidation.NewTokenValidationClient(conn)

	req := &tokenvalidation.TokenRequest{
		Token: token,
	}
	response, err := client.ValidateToken(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "Failed to validate token")
	}
	if response.Message != constants.SuccessMessage {
		return resp, fmt.Errorf("Got response error from ums: %s", response.Message)
	}

	resp.UserID = response.Data.UserId
	resp.Username = response.Data.Username
	resp.FullName = response.Data.FullName

	return resp, nil
}
