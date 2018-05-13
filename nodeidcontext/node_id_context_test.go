package nodeidcontext

import (
	"context"
	"testing"
)

func Test_NodeIDContext(t *testing.T) {
	ctx := context.Background()

	{
		_, ok := FromContext(ctx)
		if ok {
			t.Fatalf("expected %#v got %#v", false, true)
		}
	}

	{
		ctx = NewContext(ctx, "one")

		v, ok := FromContext(ctx)
		if !ok {
			t.Fatalf("expected %#v got %#v", true, false)
		}
		if v != "one" {
			t.Fatalf("expected %#v got %#v", "one", v)
		}
	}

	{
		ctx = NewContext(ctx, "two")

		v, ok := FromContext(ctx)
		if !ok {
			t.Fatalf("expected %#v got %#v", true, false)
		}
		if v != "two" {
			t.Fatalf("expected %#v got %#v", "two", v)
		}
	}
}
