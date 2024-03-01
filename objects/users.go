package objects

import "github.com/edgedb/edgedb-go"

type User struct {
	Id              edgedb.OptionalUUID `edgedb:"id" json:"id"`
	Name            edgedb.OptionalStr  `edgedb:"name" json:"name"`
	Email           edgedb.OptionalStr  `edgedb:"email" json:"email"`
	Company         edgedb.OptionalStr  `edgedb:"company" json:"company"`
	Bio             edgedb.OptionalStr  `edgedb:"bio" json:"bio"`
	Location        edgedb.OptionalStr  `edgedb:"location" json:"location"`
	TwitterUsername edgedb.OptionalStr  `edgedb:"twitter_username" json:"twitter_username"`
}

type CreateUser struct {
	Id edgedb.OptionalUUID `edgedb:"id" json:"id"`
}
