package app

//
// import "static_collector/internal/gather/network"
//
// type networkGenerator struct {
//	network []network.Stats
//	err     error
//}
//
// func (gen *networkGenerator) Get() {
//	gen.network, gen.err = network.Get()
// }
//
// func (gen *networkGenerator) Error() error {
//	return gen.err
// }
//
// func (gen *networkGenerator) Print(out chan<- Value) {
//	for _, network := range gen.network {
//		out <- Value{
//			"network." + network.Protocol + " " +
//	network.RecvQ + " " + network.SendQ + " " + network.Local + " " + network.Foreign, network.State, "bytes"}
//	}
// }
