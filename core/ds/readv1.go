package ds

import "github.com/astromechza/coinagelib/core/account"

type GetAccountTreeOptions struct {
	Root         []string
	MaximumDepth int
}

type CanGetAccount interface {
	GetAccount(fullName []string) (*account.Account, error)
}

type CanGetAccountTree interface {
	GetAccountTree(*GetAccountTreeOptions) (*account.Tree, error)
}

type CanGetAccountList interface {
	GetAccountList(*GetAccountTreeOptions) ([]*account.Account, error)
}
