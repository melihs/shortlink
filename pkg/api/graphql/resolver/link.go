// Code generated by protoc-gen-gotemplate. DO NOT EDIT.
// source: pkg/link/link.proto

package resolver

import "github.com/batazor/shortlink/pkg/link"

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

type LinkFilterInput struct {
	Url      *StringFilterInput
	Hash     *StringFilterInput
	Describe *StringFilterInput
}
