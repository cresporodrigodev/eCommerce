package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestMongoDbRepository_Close(t *testing.T) {
	type fields struct {
		host   string
		port   int64
		client *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MongoDbRepository{
				host:   tt.fields.host,
				port:   tt.fields.port,
				client: tt.fields.client,
			}
			if err := m.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDbRepository_Connect(t *testing.T) {
	type fields struct {
		host string
		port int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "local host 27017",
			fields:  fields{host: "localhost", port: 27017},
			wantErr: false,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMongoDbRepository(tt.fields.host, tt.fields.port)
			if err := m.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDbRepository_Insert(t *testing.T) {
	type collection struct {
		userAlias string `bson:"user_alias"`
		userName  string `bson:"user_name"`
		userId    string `bson:"user_id"`
	}
	type fields struct {
		host string
		port int64
	}
	type args struct {
		database   string
		collection string
		documents  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "insert one",
			fields: fields{
				host: "localhost",
				port: 27017,
			},
			args: args{
				database:   "ecommerce_testing",
				collection: "ecommerce_test",
				documents: []interface{}{collection{
					userAlias: "cresporodrigodev",
					userName:  "rodrigo crespo",
					userId:    "41589231",
				}},
			},
			wantErr: false,
		},
	}
	{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMongoDbRepository(tt.fields.host, tt.fields.port)
			if err := m.Insert(tt.args.database, tt.args.collection, tt.args.documents); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMongoDbRepository(t *testing.T) {
	type args struct {
		host string
		port int64
	}
	tests := []struct {
		name string
		args args
		want MongoDbRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMongoDbRepository(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMongoDbRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
