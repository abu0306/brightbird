package repo

import (
	"context"
	"time"

	"github.com/hunjixin/brightbird/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListParams struct {
	JobId primitive.ObjectID
}

type ITaskRepo interface {
	List(context.Context, ListParams) ([]*types.Task, error)
	UpdateVersion(ctx context.Context, id primitive.ObjectID, versionMap map[string]string) error
	Get(context.Context, primitive.ObjectID) (*types.Task, error)
	Save(context.Context, types.Task) (primitive.ObjectID, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

var _ ITaskRepo = (*TaskRepo)(nil)

func NewTaskRepo(db *mongo.Database) *TaskRepo {
	return &TaskRepo{taskCol: db.Collection("tasks")}
}

type TaskRepo struct {
	taskCol *mongo.Collection
}

func (j *TaskRepo) List(ctx context.Context, params ListParams) ([]*types.Task, error) {
	filter := bson.D{}
	if params.JobId.IsZero() {
		filter = append(filter, bson.E{Key: "jobId", Value: params.JobId})
	}
	cur, err := j.taskCol.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var tf []*types.Task
	err = cur.All(ctx, &tf)
	if err != nil {
		return nil, err
	}
	return tf, nil
}

func (j *TaskRepo) Get(ctx context.Context, id primitive.ObjectID) (*types.Task, error) {
	tf := &types.Task{}
	err := j.taskCol.FindOne(ctx, bson.D{{"_id", id}}).Decode(tf)
	if err != nil {
		return nil, err
	}
	return tf, nil
}

func (j *TaskRepo) Save(ctx context.Context, task types.Task) (primitive.ObjectID, error) {
	if task.ID.IsZero() {
		task.ID = primitive.NewObjectID()
	}

	count, err := j.taskCol.CountDocuments(ctx, bson.D{{"_id", task.ID}})
	if err != nil {
		return primitive.ObjectID{}, err
	}
	if count == 0 {
		task.BaseTime.CreateTime = time.Now().Unix()
		task.BaseTime.ModifiedTime = time.Now().Unix()
	} else {
		task.BaseTime.ModifiedTime = time.Now().Unix()
	}

	_, err = j.taskCol.InsertOne(ctx, task)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return task.ID, nil
}

func (j *TaskRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := j.taskCol.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}

func (j *TaskRepo) UpdateVersion(ctx context.Context, id primitive.ObjectID, versionMap map[string]string) error {
	update := bson.M{
		"$set": bson.M{
			"versions": versionMap,
		},
	}
	_, err := j.taskCol.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	return nil
}
