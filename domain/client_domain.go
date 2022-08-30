package domain

import "context"

type Client struct {
	ID               int64       `json:"id"`
	Client_name      string      `json:"client_name"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	Post_code        Postal_code `json:"post_code"`
	Physical_address string      `json:"physical_address"`
	Currency         Currency    `json:"currency"`
}


type ClientUseCase interface {
	Store(context.Context,*Client) error
}

type ClientRepository interface {
	Store(ctx context.Context, clnt *Client) error
}