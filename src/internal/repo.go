package internal

import (
	"context"

	"github.com/ssibrahimbas/claim-auth.go/src/entity"
	ssiMongo "github.com/ssibrahimbas/claim-auth.go/src/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	c   *mongo.Collection
	ctx context.Context
	db  *ssiMongo.MongoDB
}

type RepoParams struct {
	C   *mongo.Collection
	Ctx context.Context
	Db  *ssiMongo.MongoDB
}

func NewRepo(p *RepoParams) *Repo {
	return &Repo{
		c:   p.C,
		ctx: p.Ctx,
		db:  p.Db,
	}
}

func (r *Repo) CreateUser(u *entity.User) *entity.User {
	res, err := r.c.InsertOne(r.ctx, u)
	if err != nil {
		panic(err)
	}
	u.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return u
}

func (r *Repo) GetUserByEmail(e string) (*entity.User, bool) {
	res := r.c.FindOne(r.ctx, bson.M{"email": e})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false
		}
		panic(err)
	}
	u := &entity.User{}
	err := res.Decode(&u)
	if err != nil {
		panic(err)
	}
	return u, true
}

func (r *Repo) AddAdminRole(uId string) {
	id, err := r.db.TransformId(uId)
	if err != nil {
		panic(err)
	}
	_, err = r.c.UpdateOne(r.ctx, bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"roles": []string{"admin"},
		},
	})
	if err != nil {
		panic(err)
	}
}
