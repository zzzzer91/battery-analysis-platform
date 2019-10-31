package conv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSlice2FloatSlice(t *testing.T) {
	ast := assert.New(t)
	sl := []string{"0.1", "0.12", "3.332"}
	ret, err := StringSlice2FloatSlice(sl)
	ast.Nil(err)
	ast.Equal([]float64{0.1, 0.12, 3.332}, ret)
}
