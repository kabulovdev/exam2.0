package repo

import (
	pb "reating/genproto/reating"
)

type ReatingInfoI interface {
	Create(*pb.ReatingInfo) (*pb.ReatingInfo, error)
	GetReating(*pb.Id) (*pb.ReatingInfo, error)
	Update(*pb.ReatingInfo) (*pb.ReatingInfo, error)
	Delet(*pb.Id) (*pb.EmptyReating, error)
}
