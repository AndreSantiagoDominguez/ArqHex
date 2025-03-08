package domain

type IPurchase interface {
	Shopping(purchase *Purchase)(uint, error);
}