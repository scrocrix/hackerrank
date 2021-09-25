package medium_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type balancedBrackets struct {
	suite.Suite
}

func (this *balancedBrackets) TestIsBalanced() {
	this.Run("Validate different aspects of algorithm", func() {
		result := medium.IsBalanced("{[()]}")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("{[(])}")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("{{[[(())]]}}")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("[)")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("()")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("{")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("aa")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("{(([])[])[]}")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("{(([])[])[]}[]")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("{(([])[])[]]}")
		require.Equal(this.T(), "NO", result)
	})

	this.Run("Test case 5", func() {
		result := medium.IsBalanced("{}")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("}([[{)[]))]{){}[")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("{]]{()}{])")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("(){}")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("{}{()}{{}}")
		require.Equal(this.T(), "YES", result)
	})

	this.Run("Test case 13", func() {
		result := medium.IsBalanced("{}{}{}[([]{}{})]{}()()[[[]][()[]]]()[][]{}({}{()}([]{})()){[]}")
		require.Equal(this.T(), "YES", result)

		result = medium.IsBalanced("(([([]{{}({{[[()[][]]]}({[][()]{}}{}(){}){}([])}{()()})})]{()([])}")
		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("[]))([{}()()])([]){}[{()()}][{}[[]()]({{([{}{}[]{()}[]])[(){([])}")
		require.Equal(this.T(), "NO", result)
	})

	this.Run("Test case 9", func() {
		result := medium.IsBalanced("[]]{{}})]][[]({{[{}]}})({{}({{[]{()}([][{[()]}]{})}()})}{{}}{})]()(){}}(()({()}[([](){[]()}[])])[])[])][{[{[]}]{}([])}]()(()))}){([{}])}[(([]){[]{}})]{}({}{})}){}({{}([][](){{[][{()([[{}()]]{()}{{}{[()]}})[()[]{}](){[{}()[]]")

		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("[{{}}{[{}][]()}[]](())[[][]][]()}}[({}([[{([]){}}]()([()(){}]){([()]())}](()))(()))]]{}()[][{[{}(([]){([()]{()()}([{}][[[]{[[(({([([]){()[]}]){(())}[]}))][((([]{})[{}[[()]({({[()[]]{}(()[{}[][[{}][][]({()}[{([])}][])]][]{})")

		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("([]) }){}{((){})}}){[]}[]()(()(()))(()[{{}}]){}({{{((()([](()[][]{}){({})}{(([{({{}})}]))})))}}})]]))]]}]]))})]}]}})}))})}]}}])")

		require.Equal(this.T(), "NO", result)

		result = medium.IsBalanced("[([{{}}]{[[][][([[]]){[]}{}]]}[]{{}}{})[[]]]{{}}(()[[[[[(){}[]]({}{[]})[][[][]]]]{}]{[{}]{[{[][](()({{()}}){([]({({{[]}([([()]{()[[([({{{[]{(){}}[][]({{[([])()](())([{[]([()]{})}]){}([]){()}{}[]([[()]])}()})[{}]}()}(())}){{}()}[]]{{}})]][[]({{[{}]}})({{}({{[]{()}([][{[()]}]{})}()})}{{}}{})]()(){}}(()({()}[([](){[]()}[])])[])[])][{[{[]}]{}([])}]()(()))}){([{}])}[(([]){[]{}})]{}({}{})}){}({{}([][](){{[][{()([[{}()]]{()}{{}{[()]}})[()[]{}](){[{}()[]][{{}}{[{}][]()}[]](())[[][]][]()}}[({}([[{([]){}}]()([()(){}]){([()]())}](()))(()))]]{}()[][{[{}(([]){([()]{()()}([{}][[[]{[[(({([([]){()[]}]){(())}[]}))][((([]{})[{}[[()]({({[()[]]{}(()[{}[][[{}][][]({()}[{([])}][])]][]{})([])}){}{((){})}}){[]}[]()(()(()))(()[{{}}]){}({{{((()([](()[][]{}){({})}{(([{({{}})}]))})))}}})]]))]]}]]))})]}]}})}))})}]}}])")

		require.Equal(this.T(), "YES", result)
	})
}

func TestBalancedBrackets(t *testing.T) {
	suite.Run(t, new(balancedBrackets))
}
