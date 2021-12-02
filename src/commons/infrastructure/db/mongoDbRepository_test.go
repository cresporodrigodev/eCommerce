package db

import (
	"testing"
)

func TestMongoDbRepository_Insert(t *testing.T) {
	type collection struct {
		UserAlias string `bson:"user_alias"`
		UserName  string `bson:"user_name"`
		UserId    string `bson:"user_id"`
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
					UserAlias: "cresporodrigodev",
					UserName:  "rodrigo crespo",
					UserId:    "41589231",
				}},
			},
			wantErr: false,
		},
		{
			name: "insert many",
			fields: fields{
				host: "localhost",
				port: 27017,
			},
			args: args{
				database:   "ecommerce_testing",
				collection: "ecommerce_test",
				documents: []interface{}{collection{
					UserAlias: "cresporodrigo",
					UserName:  "crespo rodrigo",
					UserId:    "5782098",
				}, collection{
					UserAlias: "rodrigocrespo",
					UserName:  "devrc",
					UserId:    "3627190",
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
