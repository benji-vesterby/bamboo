package bamboo

import (
	"context"
	"math/rand"
	"testing"
)

var intData = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
}

var floatData = []float64{
	1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1, 11.1, 12.1, 13.1, 14.1, 15.1, 16.1, 17.1, 18.1, 19.1, 20.1,
	21.1, 22.1, 23.1, 24.1, 25.1, 26.1, 27.1, 28.1, 29.1, 30.1, 31.1, 32.1, 33.1, 34.1, 35.1, 36.1, 37.1, 38.1, 39.1, 40.1,
	41.1, 42.1, 43.1, 44.1, 45.1, 46.1, 47.1, 48.1, 49.1, 50.1, 51.1, 52.1, 53.1, 54.1, 55.1, 56.1, 57.1, 58.1, 59.1, 60.1,
	61.1, 62.1, 63.1, 64.1, 65.1, 66.1, 67.1, 68.1, 69.1, 70.1, 71.1, 72.1, 73.1, 74.1, 75.1, 76.1, 77.1, 78.1, 79.1, 80.1,
	81.1, 82.1, 83.1, 84.1, 85.1, 86.1, 87.1, 88.1, 89.1, 90.1, 91.1, 92.1, 93.1, 94.1, 95.1, 96.1, 97.1, 98.1, 99.1, 100.1,
}

const DATASETS = 10
const DATACOUNT = 500

func TestSeries_SetData_Int(t *testing.T) {

	var data = make([][]int, DATASETS)

	for i := 0; i < DATASETS; i++ {
		data[i] = make([]int, DATACOUNT)

		for j := 0; j < DATACOUNT; j++ {
			data[i][j] = rand.Int()
		}
	}

	for _, dset := range data {
		var series = Series{}

		var err error
		if err = series.SetData(dset); err == nil {
			for index := range dset {
				if value, ok := series.Get(index).(int); ok {
					if value != dset[index] {
						t.Errorf("Got [%v], Expected [%v]\n", value, dset[index])
					}
				}
			}
		} else {
			t.Error(err.Error())
		}
	}
}

func TestSeries_SetData_Float(t *testing.T) {

	var data = make([][]float64, DATASETS)

	for i := 0; i < DATASETS; i++ {
		data[i] = make([]float64, DATACOUNT)

		for j := 0; j < DATACOUNT; j++ {
			data[i][j] = rand.Float64()
		}
	}

	for _, dset := range data {
		var series = Series{}

		var err error
		if err = series.SetData(dset); err == nil {
			for index := range dset {
				if value, ok := series.Get(index).(float64); ok {
					if value != dset[index] {
						t.Errorf("Got [%v], Expected [%v]\n", value, dset[index])
					}
				}
			}
		} else {
			t.Error(err.Error())
		}
	}
}

func TestSeries_Lambda(t *testing.T) {
	tables := []struct {
		dataIn  []int
		dataOut []int
		lambda  func(ctx context.Context, column interface{}) (columnOut interface{}, override bool)
	}{
		{
			[]int{1, 2, 3},
			[]int{2, 4, 6},
			func(ctx context.Context, columnIn interface{}) (columnOut interface{}, override bool) {
				if val, ok := columnIn.(int); ok {
					columnOut = val * 2
				}

				return columnOut, true
			},
		},
	}

	for _, table := range tables {
		var series = Series{}

		var err error
		if err = series.SetData(table.dataIn); err == nil {
			if _, err = series.Lambda(context.Background(), table.lambda); err == nil {

				for index := range table.dataOut {

					if value, ok := series.Get(index).(int); ok {
						if value != table.dataOut[index] {
							t.Errorf("Got [%v], Expected [%v]\n", value, table.dataOut[index])
						}
					}
				}

			} else {
				t.Error(err.Error())
			}
		} else {
			t.Error(err.Error())
		}
	}
}

func BenchmarkSeries_Lambda_IntAddition(b *testing.B) {

	var series = Series{}
	if err := series.SetData(intData); err == nil {

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			if _, err := series.Lambda(context.Background(), intAdditionLambda); err != nil {
				b.Error(err.Error())
			}
		}

	} else {
		b.Error(err.Error())
	}
}

func BenchmarkSeries_Lambda_IntMultiplier(b *testing.B) {
	var series = Series{}
	if err := series.SetData(intData); err == nil {

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			if _, err := series.Lambda(context.Background(), intMultiplicationLambda); err != nil {
				b.Error(err.Error())
			}
		}

	} else {
		b.Error(err.Error())
	}
}

func BenchmarkSeries_Lambda_IntDivision(b *testing.B) {

	var series = Series{}
	if err := series.SetData(intData); err == nil {

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			if _, err := series.Lambda(context.Background(), intDivisionLambda); err != nil {
				b.Error(err.Error())
			}
		}

	} else {
		b.Error(err.Error())
	}
}

func intAdditionLambda(ctx context.Context, column interface{}) (column_out interface{}, override bool) {

	if val, ok := column.(int); ok {
		column_out = val + 2
	}

	return column_out, true
}

func intMultiplicationLambda(ctx context.Context, column interface{}) (column_out interface{}, override bool) {

	if val, ok := column.(int); ok {
		column_out = val * 2
	}

	return column_out, true
}

func intDivisionLambda(ctx context.Context, column interface{}) (column_out interface{}, override bool) {

	if val, ok := column.(int); ok {
		column_out = val / 2
	}

	return column_out, true
}
