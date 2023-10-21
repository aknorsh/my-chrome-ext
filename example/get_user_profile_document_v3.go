package example

import (
	"context"
	"errors"

	pb "github.com/aknorsh/awesome-server/protofiles"
)

func HandleGetUserProfileDocumentV3(ctx context.Context, req *pb.GetUserProfileDocumentV3Request) (*pb.GetUserProfileDocumentV3Response, error){
	uid := req.UserId

	name, err := GetNameFromUID(uid)
	if err != nil {
		return nil, errors.New("GetNameFromUID: %w", err)
	}

	desc, err := GetDescFromUID(uid)
	if err != nil {
		return nil, errors.New("GetDescFromUID: %w", err)
	}

	resp := BuildResponse(ctx, name, desc, req.IsXXX, req.HasYYY)

	return resp, nil
}