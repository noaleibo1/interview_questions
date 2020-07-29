package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type testBtreeSuite struct {
	suite.Suite
}

func TestBTree(t *testing.T) {
	suite.Run(t, &testBtreeSuite{})
}

func (suite *testBtreeSuite) TestSingleNodeIsValid() {
	expectedResult := fmt.Sprintf("%v", turnStringArrayToStringOfByteArrays([]string{singleNodeTree.Value}))

	serializedTree := Serialize(singleNodeTree)
	suite.Require().Equal(expectedResult, serializedTree)
}

func (suite *testBtreeSuite) TestSeralizeComplexTreeIsValid() {
	serializedTree := Serialize(complexTree)
	suite.Require().Equal(complexTreeSerialized, serializedTree)
}

func (suite *testBtreeSuite) TestDeseralizeComplexTreeIsValid() {
	deserialize := Deserialize(complexTreeSerialized)
	suite.Require().True(areTwoTreesEqual(complexTree, deserialize))
}

func (suite *testBtreeSuite) test() {

}

func (suite *testBtreeSuite) TestNotATree() {
	graphTreeNode.Left = graphTree

	suite.Require().Panics(func() {
		Serialize(graphTree)
	})
}

func (suite *testBtreeSuite) TestDeserializeInputIsValid(){
	suite.Require().Panics(func() {
		Deserialize(InvalidDeserializeInput)
	})
}