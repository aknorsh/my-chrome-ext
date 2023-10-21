package example

import (
	"testing"

	pb "github.com/aknorsh/awesome-server/protofiles"
)

func TestHandleGetUserProfileDocumentV4(t *testing.T) {
	type args struct {
		req *pb.GetUserProfileDocumentV4Request
	}
	tests := []struct {
		name string
		args args
		resp *pb.GetUserProfileDocumentV4Response
	}{
		{
			name: "Success",
			args: args{req: ...},
			resp: &pb.GetUserProfileDocumentV4Response{
				...
			},
		}
		{
			name: "Failure",
			args: args{req: ...},
			resp: &pb.GetUserProfileDocumentV4Response{
				...
			},
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := HandleGetUserProfileDocumentV4(tt.args.req)
			if err != nil {
				t.Errorf("HandleGetUserProfileDocumentV4() error = %v", err)
				return
			}
			if resp != tt.resp {
				t.Errorf("HandleGetUserProfileDocumentV4() got = %v, want %v", resp, tt.resp)
			}
		})
	}
}