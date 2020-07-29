package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type testHelpersSuite struct {
	suite.Suite
}

func TestHelpers(t *testing.T) {
	suite.Run(t, &testHelpersSuite{})
}

func (suite *testHelpersSuite) TestAreTwoTreesEqualShouldReturnTrue() {
	suite.Require().True(areTwoTreesEqual(complexTree, complexTree))
}

func (suite *testHelpersSuite) TestAreTwoTreesEqualShouldReturnFalse() {
	tree1 := &Btree{
		Value: "a",
		Left:  nil,
		Right: nil,
	}
	tree2 := &Btree{
		Value: "b",
		Left:  nil,
		Right: nil,
	}

	suite.Require().Falsef(areTwoTreesEqual(tree1, tree2), "The trees are different")
}

func (suite *testBtreeSuite) TestTurnStringArrayToStringOfByteArrays() {
	serializedTree := turnStringArrayToStringOfByteArrays(complexTreeSerializedHumanReadable)
	suite.Require().Equal(complexTreeSerializedManualyForTests,serializedTree)
}