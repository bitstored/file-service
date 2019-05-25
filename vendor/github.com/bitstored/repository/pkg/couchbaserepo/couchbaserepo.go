package couchbaserepo

import (
	"context"

	"github.com/couchbase/gocb"
)

// Repository is a couchbase implementation to of the Repository pattern.
type Repository struct {
	bucket *gocb.Bucket
	TTL    uint32
}

// NewRepository returnes a new Repository.
func NewRepository(b *gocb.Bucket) *Repository {
	return &Repository{
		bucket: b,
	}
}

// Create creates a new document. It returns a DocumentExistsError if the key already exists.
func (r *Repository) Create(ctx context.Context, key string, doc interface{}) (uint, error) {
	gocbCas, err := r.bucket.Insert(key, doc, r.TTL)
	return uint(gocbCas), err
}

// Update replaces a existing document.
func (r *Repository) Update(ctx context.Context, key string, cas uint, doc interface{}) (uint, error) {
	gocbCas, err := r.bucket.Replace(key, doc, gocb.Cas(cas), r.TTL)
	return uint(gocbCas), err
}

// Read reads the document from the key.
func (r *Repository) Read(ctx context.Context, key string, doc interface{}) (uint, error) {
	gocbCas, err := r.bucket.Get(key, &doc)
	return uint(gocbCas), err
}

// Delete deletes the document with the given key.
func (r *Repository) Delete(ctx context.Context, key string, cas uint) (uint, error) {
	cas1, err := r.bucket.Remove(key, gocb.Cas(cas))
	return uint(cas1), err
}
