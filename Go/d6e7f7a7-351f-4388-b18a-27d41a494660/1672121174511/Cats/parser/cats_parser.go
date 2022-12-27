// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121174511/Cats.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cats

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CatsParser struct {
	*antlr.BaseParser
}

var catsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func catsParserInit() {
	staticData := &catsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CATS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cats",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 10, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 7, 0, 4, 1, 0, 0, 0, 2, 7, 1, 0, 0, 0, 4, 5, 3, 2,
		1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 5, 1, 0, 0, 8, 3, 1, 0,
		0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CatsParserInit initializes any static state used to implement CatsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCatsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CatsParserInit() {
	staticData := &catsParserStaticData
	staticData.once.Do(catsParserInit)
}

// NewCatsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCatsParser(input antlr.TokenStream) *CatsParser {
	CatsParserInit()
	this := new(CatsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &catsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cats.g4"

	return this
}

// CatsParser tokens.
const (
	CatsParserEOF      = antlr.TokenEOF
	CatsParserCATS     = 1
	CatsParserOWNKEY   = 2
	CatsParserSPLITKEY = 3
	CatsParserWS       = 4
)

// CatsParser rules.
const (
	CatsParserRULE_expression = 0
	CatsParserRULE_cats       = 1
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CatsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CatsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cats() ICatsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CatsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CatsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CatsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CatsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CatsParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Cats()
	}
	{
		p.SetState(5)
		p.Match(CatsParserEOF)
	}

	return localctx
}

// ICatsContext is an interface to support dynamic dispatch.
type ICatsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCatsContext differentiates from other interfaces.
	IsCatsContext()
}

type CatsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatsContext() *CatsContext {
	var p = new(CatsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CatsParserRULE_cats
	return p
}

func (*CatsContext) IsCatsContext() {}

func NewCatsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatsContext {
	var p = new(CatsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CatsParserRULE_cats

	return p
}

func (s *CatsContext) GetParser() antlr.Parser { return s.parser }

func (s *CatsContext) CATS() antlr.TerminalNode {
	return s.GetToken(CatsParserCATS, 0)
}

func (s *CatsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CatsListener); ok {
		listenerT.EnterCats(s)
	}
}

func (s *CatsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CatsListener); ok {
		listenerT.ExitCats(s)
	}
}

func (p *CatsParser) Cats() (localctx ICatsContext) {
	this := p
	_ = this

	localctx = NewCatsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CatsParserRULE_cats)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)
		p.Match(CatsParserCATS)
	}

	return localctx
}
