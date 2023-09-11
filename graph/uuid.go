package graph

import (
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"

	"io"
	"strconv"
)

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	if str, ok := v.(string); ok {
		return uuid.Parse(str)
	}
	return uuid.Nil, errors.New("input must be an RFC-4122 formatted string")
}

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(id.String()))
	})
}
