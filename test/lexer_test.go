package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

func Test_lexer(t *testing.T) {

	lexer_rule := `
rule "test" salience 1
begin
规则管理
end
`
	dataContext := context.NewDataContext()
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(lexer_rule)
	assert.Error(t, e1)
	// if e1 != nil {
	// 	panic(e1)
	// }

	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	assert.Error(t, e)

}
