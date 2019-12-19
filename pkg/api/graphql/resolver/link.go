// Code generated by protoc-gen-gotemplate. DO NOT EDIT.
// source: pkg/link/link.proto

package resolver

import (
	"github.com/batazor/shortlink/pkg/link"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type LinkResolver struct {
	*link.Link
}

func (r *LinkResolver) Url() string {
	return r.Link.Url
}

func (r *LinkResolver) Hash() string {
	return r.Link.Hash
}

func (r *LinkResolver) Describe() string {
	return r.Link.Describe
}

func (r *LinkResolver) Created_at() *timestamp.Timestamp {
	return r.Link.CreatedAt
}

func (r *LinkResolver) Updated_at() *timestamp.Timestamp {
	return r.Link.UpdatedAt
}

type LinkFilterInput struct {
	Url        *StringFilterInput
	Hash       *StringFilterInput
	Describe   *StringFilterInput
	Created_at *StringFilterInput
	Updated_at *StringFilterInput
}
