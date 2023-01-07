package database

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// VectorizedCloudDatabase is a struct representing a connection to a vectorized cloud database
type VectorizedCloudDatabase struct {
	client *bigquery.Client
}

// NewVectorizedCloudDatabase returns a new VectorizedCloudDatabase instance
func NewVectorizedCloudDatabase(projectID string) (*VectorizedCloudDatabase, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &VectorizedCloudDatabase{client: client}, nil
}

// Query executes a SQL query on the vectorized cloud database
func (db *VectorizedCloudDatabase) Query(query string) (*bigquery.RowIterator, error) {
	ctx := context.Background()
	q := db.client.Query(query)
	return q.Read(ctx)
}

// InsertRows inserts rows of data into a table in the vectorized cloud database
func (db *VectorizedCloudDatabase) InsertRows(table string, rows [][]bigquery.Value) error {
	ctx := context.Background()
	u := db.client.Dataset(db.client.Dataset(table).DatasetID).Table(table).Uploader()
	for _, row := range rows {
		if err := u.Put(ctx, row); err != nil {
			return err
		}
	}
	return u.Close()
}

// ListTables lists the tables in the vectorized cloud database
func (db *VectorizedCloudDatabase) ListTables() ([]*bigquery.Table, error) {
	ctx := context.Background()
	it := db.client.Datasets(ctx).Tables(ctx)
	var tables []*bigquery.Table
	for {
		table, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

