package public

// @program:     rpc_exercise
// @file:        service.go
// @author:      bowen
// @create:      2022-11-14 20:41
// @description:

type ServiceA struct {
}

type Args struct {
	X, Y int
}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}
