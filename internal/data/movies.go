package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// Implement a UnmarshalJSON() method on the Runtime type so that it satisfies the
// json.Unmarshaler interface. IMPORTANT: Because UnmarshalJSON() needs to modify the
// receiver (our Runtime type), we must use a pointer receiver for this to work
// correctly.
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	runtime, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(runtime)
	return nil
}

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title" validate:"required,lte=500"`
	Year      int32     `json:"year,omitempty" validate:"required,ne=0,gte=1888,lteyear"`
	Runtime   Runtime   `json:"runtime,omitempty" validate:"required,gt=0"`
	Genres    []string  `json:"genres,omitempty" validate:"min=1,max=5,unique"`
	Version   int32     `json:"version,omitempty"`
}

// func (m Movie) MarshalJSON() ([]byte, error) {
// 	// Create a variable holding the custom runtime string, just like before.
// 	var runtime string

// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d mins", m.Runtime)
// 	}

// 	// Define a MovieAlias type which has the underlying type Movie. Due to the way that
// 	// Go handles type definitions (https://golang.org/ref/spec#Type_definitions) the
// 	// MovieAlias type will contain all the fields that our Movie struct has but,
// 	// importantly, none of the methods.
// 	type MovieAlias Movie

// 	// Embed the MovieAlias type inside the anonymous struct, along with a Runtime field
// 	// that has the type string and the necessary struct tags. It's important that we
// 	// embed the MovieAlias type here, rather than the Movie type directly, to avoid
// 	// inheriting the MarshalJSON() method of the Movie type (which would result in an
// 	// infinite loop during encoding).
// 	aux := struct {
// 		MovieAlias
// 		Runtime string `json:"runtime,omitempty"`
// 	}{
// 		MovieAlias: MovieAlias(m),
// 		Runtime:    runtime,
// 	}

// 	return json.Marshal(aux)
// }
