package govector

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestEverything(t *testing.T) {
	x := IntToVector([]int{1, 2, 3, 4, 6, 5})
	w := Float64ToVector([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 4.0})

	d_x, err := x.Diff()
	assert.Equal(t, nil, err, "Error calculating vector differences")

	d_w, err := w.Diff()
	assert.Equal(t, nil, err, "Error calculating vector differences")

	empirical, err := d_w.Ecdf()
	assert.Equal(t, nil, err, "Error creating CDF function")

	_, err = empirical(2.4)
	assert.Equal(t, nil, err, "Error calculating CDF percentile")

	_, err = d_x.WeightedMean(d_w)
	assert.Equal(t, nil, err, "Error calculating weighted mean")

	_, err = x.Quantiles(Float64ToVector([]float64{0.05, 0.95}))
	assert.Equal(t, nil, err, "Error calculating quantiles")

	_, err = x.Cumsum()
	assert.Equal(t, nil, err, "Error calculating cumulative sum")
}