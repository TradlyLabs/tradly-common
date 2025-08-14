package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatArray_Value(t *testing.T) {
	// Test empty array
	emptyArr := FlatArray[string]{}
	val, err := emptyArr.Value()
	assert.NoError(t, err)
	assert.Equal(t, "{}", val)

	// Test array with single string element
	singleStrArr := FlatArray[string]{"test"}
	val, err = singleStrArr.Value()
	assert.NoError(t, err)
	assert.Equal(t, "{'test'}", val)

	// Test array with multiple string elements
	multiStrArr := FlatArray[string]{"test1", "test2", "test's3"}
	val, err = multiStrArr.Value()
	assert.NoError(t, err)
	assert.Equal(t, "{'test1','test2','test''s3'}", val)

	// Test array with integer elements
	intArr := FlatArray[int]{1, 2, 3}
	val, err = intArr.Value()
	assert.NoError(t, err)
	assert.Equal(t, "{1,2,3}", val)

	// Test array with float elements
	floatArr := FlatArray[float64]{1.1, 2.2, 3.3}
	val, err = floatArr.Value()
	assert.NoError(t, err)
	assert.Equal(t, "{1.1,2.2,3.3}", val)
}

func TestFlatArray_Scan(t *testing.T) {
	// Test scanning from PostgreSQL array format
	var strArr FlatArray[string]
	err := strArr.Scan("{'test1','test2','test''s3'}")
	assert.NoError(t, err)
	assert.Equal(t, FlatArray[string]{"test1", "test2", "test's3"}, strArr)

	// Test scanning from JSON array format
	var intArr FlatArray[int]
	err = intArr.Scan("[1,2,3]")
	assert.NoError(t, err)
	assert.Equal(t, FlatArray[int]{1, 2, 3}, intArr)

	// Test scanning from single value
	var singleArr FlatArray[float64]
	err = singleArr.Scan("42.5")
	assert.NoError(t, err)
	assert.Equal(t, FlatArray[float64]{42.5}, singleArr)

	// Test scanning nil
	var nilArr FlatArray[string]
	err = nilArr.Scan(nil)
	assert.NoError(t, err)
	assert.Nil(t, nilArr)
}