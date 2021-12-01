package template

import (
	"fmt"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_removeUnspecifiedValueWhenNil(t *testing.T) {
	var fixture = gendoc.Enum{
		Values: nil,
	}

	removeUnspecifiedValue(&fixture)
	require.Len(t, fixture.Values, 0, "size should be 0")
}

func Test_removeUnspecifiedValueWhenEmpty(t *testing.T) {
	var values = make([]*gendoc.EnumValue, 0)
	var fixture = gendoc.Enum{
		Values: values,
	}

	removeUnspecifiedValue(&fixture)
	require.Len(t, fixture.Values, 0, "size should be 0")
}

func Test_removeUnspecifiedValueDoesntContainAny(t *testing.T) {
	var values = make([]*gendoc.EnumValue, 0)

	for i := 0; i < 5; i++ {
		values = append(values, &gendoc.EnumValue{
			Name: fmt.Sprintf("TEST_%d", i),
		})
	}

	var fixture = gendoc.Enum{
		Values: values,
	}

	removeUnspecifiedValue(&fixture)
	require.Len(t, fixture.Values, 5, "size should be 5")
}

func Test_removeUnspecifiedValueWhenContainsOnlyOneUnspecified(t *testing.T) {
	var values = make([]*gendoc.EnumValue, 0)

	values = append(values, &gendoc.EnumValue{
		Name: "TEST_UNSPECIFIED",
	})

	var fixture = gendoc.Enum{
		Values: values,
	}

	removeUnspecifiedValue(&fixture)
	require.Len(t, fixture.Values, 0, "size should be 0")
}

func Test_removeUnspecifiedValueWhenContainsOnlyOneUnknown(t *testing.T) {
	var values = make([]*gendoc.EnumValue, 0)

	values = append(values, &gendoc.EnumValue{
		Name: "TEST_UNKNOWN",
	})

	var fixture = gendoc.Enum{
		Values: values,
	}

	removeUnspecifiedValue(&fixture)
	require.Len(t, fixture.Values, 0, "size should be 0")
}

func Test_removeUnspecifiedValueWhenContainsMany(t *testing.T) {
	var values = make([]*gendoc.EnumValue, 0)

	values = append(values, &gendoc.EnumValue{
		Name: "TEST_UNSPECIFIED",
	})

	for i := 0; i < 5; i++ {
		values = append(values, &gendoc.EnumValue{
			Name: fmt.Sprintf("TEST_%d", i),
		})
	}

	var fixture = gendoc.Enum{
		Values: values,
	}

	removeUnspecifiedValue(&fixture)
	require.Len(t, fixture.Values, 5, "size should be 5")
}
