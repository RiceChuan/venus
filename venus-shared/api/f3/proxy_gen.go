// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package f3

import (
	"context"

	address "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-f3/certs"
)

type F3Struct struct {
	Internal struct {
		F3GetCertificate       func(ctx context.Context, instance uint64) (*certs.FinalityCertificate, error) `perm:"read"`
		F3GetLatestCertificate func(ctx context.Context) (*certs.FinalityCertificate, error)                  `perm:"read"`
		F3Participate          func(ctx context.Context, minerID address.Address) (<-chan string, error)      `perm:"admin"`
	}
}

func (s *F3Struct) F3GetCertificate(p0 context.Context, p1 uint64) (*certs.FinalityCertificate, error) {
	return s.Internal.F3GetCertificate(p0, p1)
}
func (s *F3Struct) F3GetLatestCertificate(p0 context.Context) (*certs.FinalityCertificate, error) {
	return s.Internal.F3GetLatestCertificate(p0)
}
func (s *F3Struct) F3Participate(p0 context.Context, p1 address.Address) (<-chan string, error) {
	return s.Internal.F3Participate(p0, p1)
}
