// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"fmt"

	"github.com/gobwas/glob"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ReplaceMatchArguments[K any] struct {
	Target      ottl.GetSetter[K] `ottlarg:"0"`
	Pattern     string            `ottlarg:"1"`
	Replacement string            `ottlarg:"2"`
}

func NewReplaceMatchFactory[K any]() ottl.Factory[K] {
	return ottl.NewFactory("replace_match", &ReplaceMatchArguments[K]{}, createReplaceMatchFunction[K])
}

func createReplaceMatchFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	args, ok := oArgs.(*ReplaceMatchArguments[K])

	if !ok {
		return nil, fmt.Errorf("ReplaceMatchFactory args must be of type *ReplaceMatchArguments[K]")
	}

	return replaceMatch(args.Target, args.Pattern, args.Replacement)
}

func replaceMatch[K any](target ottl.GetSetter[K], pattern string, replacement string) (ottl.ExprFunc[K], error) {
	glob, err := glob.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("the pattern supplied to replace_match is not a valid pattern: %w", err)
	}
	return func(ctx context.Context, tCtx K) (interface{}, error) {
		val, err := target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}
		if val == nil {
			return nil, nil
		}
		if valStr, ok := val.(string); ok {
			if glob.Match(valStr) {
				err = target.Set(ctx, tCtx, replacement)
				if err != nil {
					return nil, err
				}
			}
		}
		return nil, nil
	}, nil
}
