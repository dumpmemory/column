// +build ignore

package column

import (
	"testing"

	"github.com/kelindar/bitmap"
	"github.com/stretchr/testify/assert"
)

func TestOfnumbers(t *testing.T) {
	c := makenumbers().(*columnnumber)

	{ // Set the value at index
		c.Update(9, number(99))
		assert.Equal(t, 10, len(c.data))
		assert.True(t, c.Contains(9))
	}

	{ // Get the values
		v, ok := c.Value(9)
		assert.Equal(t, number(99), v)
		assert.True(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(99), f)
		assert.True(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(99), i)
		assert.True(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(99), u)
		assert.True(t, ok)
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Intersect(&other)
		assert.Equal(t, uint64(0b1000000000), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Difference(&other)
		assert.Equal(t, uint64(0xfffffffffffffdff), other[0])
	}

	{
		other := bitmap.Bitmap{0xffffffffffffffff}
		c.Union(&other)
		assert.Equal(t, uint64(0xffffffffffffffff), other[0])
	}

	{ // Remove the value
		c.DeleteMany(&bitmap.Bitmap{0b1000000000})
		v, ok := c.Value(9)
		assert.Equal(t, number(0), v)
		assert.False(t, ok)

		f, ok := c.Float64(9)
		assert.Equal(t, float64(0), f)
		assert.False(t, ok)

		i, ok := c.Int64(9)
		assert.Equal(t, int64(0), i)
		assert.False(t, ok)

		u, ok := c.Uint64(9)
		assert.Equal(t, uint64(0), u)
		assert.False(t, ok)
	}

	{ // Update several items at once
		c.UpdateMany([]Update{
			{Kind: UpdatePut, Index: 1, Value: number(2)},
			{Kind: UpdatePut, Index: 2, Value: number(3)},
			{Kind: UpdateAdd, Index: 1, Value: number(2)},
		})
		assert.True(t, c.Contains(1))
		assert.True(t, c.Contains(2))
		v, _ := c.Int64(1)
		assert.Equal(t, int64(4), v)
	}

}
