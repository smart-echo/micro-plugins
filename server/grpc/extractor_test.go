package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHandler struct{}

type TestRequest struct{}

type TestResponse struct{}

func (t *TestHandler) Test(ctx context.Context, req *TestRequest, rsp *TestResponse) error {
	return nil
}

type TestRequest2 struct {
	Int64Arg        int64                         `json:"Int64Arg,omitempty"`
	StringArg       string                        `json:"StringArg,omitempty"`
	StringSliceArg  []string                      `json:"StringSliceArg,omitempty"`
	PointerSliceArg []*TestRequest                `json:"PointerSliceArg,omitempty"`
	MapArg          map[string]bool               `json:"MapArg,omitempty"`
	MapPointerArg   map[*TestRequest]TestResponse `json:"MapPointerArg,omitempty"`
}
type TestResponse2 struct {
	Int64Arg        int64                         `json:"Int64Arg,omitempty"`
	StringArg       string                        `json:"StringArg,omitempty"`
	StringSliceArg  []string                      `json:"StringSliceArg,omitempty"`
	PointerSliceArg []*TestRequest                `json:"PointerSliceArg,omitempty"`
	MapArg          map[string]bool               `json:"MapArg,omitempty"`
	MapPointerArg   map[*TestRequest]TestResponse `json:"MapPointerArg,omitempty"`
}

func (t *TestHandler) Test2(ctx context.Context, req *TestRequest2, rsp *TestResponse2) error {
	return nil
}

func TestExtractEndpoint(t *testing.T) {
	handler := &TestHandler{}
	typ := reflect.TypeOf(handler)

	type param struct {
		name  string
		value string
	}
	tcs := []struct {
		name    string
		reqName string
		reqType string
		reqArgs []param
		rspName string
		rspType string
		rspArgs []param
	}{
		{
			name:    "Test",
			reqName: "TestRequest",
			reqType: "TestRequest",
			rspName: "TestResponse",
			rspType: "TestResponse",
		},
		{
			name:    "Test2",
			reqName: "TestRequest2",
			reqType: "TestRequest2",
			rspName: "TestResponse2",
			rspType: "TestResponse2",
			reqArgs: []param{
				{
					name:  "Int64Arg",
					value: "int64",
				},
				{
					name:  "StringArg",
					value: "string",
				},
				{
					name:  "StringSliceArg",
					value: "[]string",
				},
				{
					name:  "PointerSliceArg",
					value: "[]TestRequest",
				},
				{
					name:  "MapArg",
					value: "map[string]bool",
				},
				{
					name:  "MapPointerArg",
					value: "map[TestRequest]TestResponse",
				},
			},
			rspArgs: []param{
				{
					name:  "Int64Arg",
					value: "int64",
				},
				{
					name:  "StringArg",
					value: "string",
				},
				{
					name:  "StringSliceArg",
					value: "[]string",
				},
				{
					name:  "PointerSliceArg",
					value: "[]TestRequest",
				},
				{
					name:  "MapArg",
					value: "map[string]bool",
				},
				{
					name:  "MapPointerArg",
					value: "map[TestRequest]TestResponse",
				},
			},
		},
	}

	for _, tc := range tcs {
		m, ok := typ.MethodByName(tc.name)
		assert.True(t, ok)
		e := extractEndpoint(m)
		assert.Equal(t, tc.name, e.Name)
		assert.NotNil(t, e.Request)
		assert.NotNil(t, e.Response)
		assert.Equal(t, tc.reqName, e.Request.Name)
		assert.Equal(t, tc.reqType, e.Request.Type)
		assert.Equal(t, tc.rspName, e.Response.Name)
		assert.Equal(t, tc.rspType, e.Response.Type)
		for i, v := range tc.reqArgs {
			assert.Equal(t, v.name, e.Request.Values[i].Name)
			assert.Equal(t, v.value, e.Request.Values[i].Type)
		}
		for i, v := range tc.rspArgs {
			assert.Equal(t, v.name, e.Response.Values[i].Name)
			assert.Equal(t, v.value, e.Response.Values[i].Type)
		}
	}

}
