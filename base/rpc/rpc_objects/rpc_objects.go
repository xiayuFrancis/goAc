package rpc_objects

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, replu *int) error{
	*replu = args.N * args.M
	return nil
}