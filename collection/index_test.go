package collection

import (
	"testing"

	"github.com/beaconsoftwarellc/gadget/v2/generator"
	assert1 "github.com/stretchr/testify/assert"
)

type Testable struct {
	A string
	F int
}

func GetF(t Testable) []int {
	return []int{t.F}
}

func TestNewIndex(t *testing.T) {
	assert := assert1.New(t)
	actual := NewIndex(GetF)
	assert.NotNil(actual)
	assert.Equal(0, actual.Len())

	actual = NewIndex(GetF, Testable{F: 1})
	assert.NotNil(actual)
	assert.Equal(1, actual.Len())
}

func TestIndex_Add(t *testing.T) {
	assert := assert1.New(t)
	actual := NewIndex(GetF)
	expected := Testable{A: "0", F: generator.Int()}
	expected1 := Testable{A: "1", F: generator.Int()}

	actual.Add(expected, expected1)
	assert.Equal(2, actual.Len())
	assert.ElementsMatch([]Testable{expected}, actual.Get(expected.F))
	assert.ElementsMatch([]Testable{expected1}, actual.Get(expected1.F))

	expected2 := Testable{A: "2", F: expected.F}
	actual.Add(expected2)
	assert.Equal(2, actual.Len())
	assert.ElementsMatch([]Testable{expected, expected2}, actual.Get(expected.F))
	assert.ElementsMatch([]Testable{expected1}, actual.Get(expected1.F))
}

func TestIndex_Len(t *testing.T) {
	assert := assert1.New(t)
	actual := NewIndex(GetF)
	assert.Equal(0, actual.Len())
	testable := Testable{F: generator.Int()}
	actual.Add(testable)
	assert.Equal(1, actual.Len())
	actual.Add(testable)
	assert.Equal(1, actual.Len())
	actual.Add(Testable{F: generator.Int()})
	assert.Equal(2, actual.Len())
	actual.Remove(testable)
	assert.Equal(1, actual.Len())
}

func TestIndex_Get_Empty(t *testing.T) {
	assert := assert1.New(t)
	idx := NewIndex(GetF)
	actual := idx.Get(generator.Int())
	assert.Empty(actual)
	idx.Add(Testable{F: generator.Int()})
	idx.Add(Testable{F: generator.Int()})
	actual = idx.Get(generator.Int())
	assert.Empty(actual)
}

func TestIndex_Remove(t *testing.T) {
	assert := assert1.New(t)
	actual := NewIndex(GetF)
	testable := Testable{A: "0", F: generator.Int()}
	assert.NotPanics(func() { actual.Remove(testable) })

	testable1 := Testable{A: "1", F: generator.Int()}
	testable2 := Testable{A: "2", F: testable.F}
	actual.Add(testable)
	actual.Add(testable1)
	actual.Add(testable2)
	actual.Remove(testable)
	assert.Equal(2, actual.Len())
	actual.Remove(testable1)
	assert.Equal(1, actual.Len())
	actual.Remove(testable2)
	assert.Equal(0, actual.Len())
}
