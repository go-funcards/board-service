package board

import (
	"github.com/go-funcards/board-service/proto/v1"
	"github.com/go-funcards/slice"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Member struct {
	MemberID string   `json:"member_id" bson:"member_id,omitempty"`
	Roles    []string `json:"roles" bson:"roles,omitempty"`
	Delete   bool     `json:"-" bson:"-"`
}

type Board struct {
	BoardID   string    `json:"board_id" bson:"_id,omitempty"`
	OwnerID   string    `json:"owner_id" bson:"owner_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Metadata  string    `json:"metadata" bson:"metadata,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	Members   []Member  `json:"members" bson:"members,omitempty"`
}

type Filter struct {
	BoardIDs  []string `json:"board_ids,omitempty"`
	OwnerIDs  []string `json:"owner_ids,omitempty"`
	MemberIDs []string `json:"member_ids,omitempty"`
}

func (b Board) toProto() *v1.BoardsResponse_Board {
	return &v1.BoardsResponse_Board{
		BoardId:   b.BoardID,
		OwnerId:   b.OwnerID,
		Name:      b.Name,
		Metadata:  b.Metadata,
		CreatedAt: timestamppb.New(b.CreatedAt),
		Members: slice.Map(b.Members, func(m Member) *v1.BoardsResponse_Board_Member {
			return m.toProto()
		}),
	}
}

func (m Member) toProto() *v1.BoardsResponse_Board_Member {
	return &v1.BoardsResponse_Board_Member{
		MemberId: m.MemberID,
		Roles:    m.Roles,
	}
}

func CreateBoard(in *v1.CreateBoardRequest) Board {
	return Board{
		BoardID:   in.GetBoardId(),
		OwnerID:   in.GetOwnerId(),
		Name:      in.GetName(),
		Metadata:  in.GetMetadata(),
		CreatedAt: time.Now().UTC(),
		Members: slice.Map(in.GetMembers(), func(item *v1.CreateBoardRequest_Member) Member {
			return Member{
				MemberID: item.GetMemberId(),
				Roles:    item.GetRoles(),
			}
		}),
	}
}

func UpdateBoard(in *v1.UpdateBoardRequest) Board {
	return Board{
		BoardID:  in.GetBoardId(),
		Name:     in.GetName(),
		Metadata: in.GetMetadata(),
		Members: slice.Map(in.GetMembers(), func(item *v1.UpdateBoardRequest_Member) Member {
			return Member{
				MemberID: item.GetMemberId(),
				Roles:    item.GetRoles(),
				Delete:   item.GetDelete(),
			}
		}),
	}
}

func CreateFilter(in *v1.BoardsRequest) Filter {
	return Filter{
		BoardIDs:  in.GetBoardIds(),
		OwnerIDs:  in.GetOwnerIds(),
		MemberIDs: in.GetMemberIds(),
	}
}
