package account

type Account struct {
	FullName []string
	Note     string
}

type Tree struct {
	Account
	Children map[string]Tree
}
