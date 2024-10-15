/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// TokenAuthorizationResponseBuilder contains the data and logic needed to build 'token_authorization_response' objects.
//
//
type TokenAuthorizationResponseBuilder struct {
	bitmap_ uint32
	account *AccountBuilder
}

// NewTokenAuthorizationResponse creates a new builder of 'token_authorization_response' objects.
func NewTokenAuthorizationResponse() *TokenAuthorizationResponseBuilder {
	return &TokenAuthorizationResponseBuilder{}
}

// Account sets the value of the 'account' attribute to the given value.
//
//
func (b *TokenAuthorizationResponseBuilder) Account(value *AccountBuilder) *TokenAuthorizationResponseBuilder {
	b.account = value
	if value != nil {
		b.bitmap_ |= 1
	} else {
		b.bitmap_ &^= 1
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *TokenAuthorizationResponseBuilder) Copy(object *TokenAuthorizationResponse) *TokenAuthorizationResponseBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	if object.account != nil {
		b.account = NewAccount().Copy(object.account)
	} else {
		b.account = nil
	}
	return b
}

// Build creates a 'token_authorization_response' object using the configuration stored in the builder.
func (b *TokenAuthorizationResponseBuilder) Build() (object *TokenAuthorizationResponse, err error) {
	object = new(TokenAuthorizationResponse)
	object.bitmap_ = b.bitmap_
	if b.account != nil {
		object.account, err = b.account.Build()
		if err != nil {
			return
		}
	}
	return
}
