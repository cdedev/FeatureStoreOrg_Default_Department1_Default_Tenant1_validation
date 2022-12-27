// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116307224/MaternityBenefitsScored.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaternityBenefitsScored

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

type MaternityBenefitsScoredParser struct {
	*antlr.BaseParser
}

var maternitybenefitsscoredParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func maternitybenefitsscoredParserInit() {
	staticData := &maternitybenefitsscoredParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MATERNITYBENEFITSSCORED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "maternitybenefitsscored",
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

// MaternityBenefitsScoredParserInit initializes any static state used to implement MaternityBenefitsScoredParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMaternityBenefitsScoredParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MaternityBenefitsScoredParserInit() {
	staticData := &maternitybenefitsscoredParserStaticData
	staticData.once.Do(maternitybenefitsscoredParserInit)
}

// NewMaternityBenefitsScoredParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMaternityBenefitsScoredParser(input antlr.TokenStream) *MaternityBenefitsScoredParser {
	MaternityBenefitsScoredParserInit()
	this := new(MaternityBenefitsScoredParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &maternitybenefitsscoredParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MaternityBenefitsScored.g4"

	return this
}

// MaternityBenefitsScoredParser tokens.
const (
	MaternityBenefitsScoredParserEOF                     = antlr.TokenEOF
	MaternityBenefitsScoredParserMATERNITYBENEFITSSCORED = 1
	MaternityBenefitsScoredParserOWNKEY                  = 2
	MaternityBenefitsScoredParserSPLITKEY                = 3
	MaternityBenefitsScoredParserWS                      = 4
)

// MaternityBenefitsScoredParser rules.
const (
	MaternityBenefitsScoredParserRULE_expression              = 0
	MaternityBenefitsScoredParserRULE_maternitybenefitsscored = 1
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
	p.RuleIndex = MaternityBenefitsScoredParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaternityBenefitsScoredParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Maternitybenefitsscored() IMaternitybenefitsscoredContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaternitybenefitsscoredContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaternitybenefitsscoredContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MaternityBenefitsScoredParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaternityBenefitsScoredListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaternityBenefitsScoredListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MaternityBenefitsScoredParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MaternityBenefitsScoredParserRULE_expression)

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
		p.Maternitybenefitsscored()
	}
	{
		p.SetState(5)
		p.Match(MaternityBenefitsScoredParserEOF)
	}

	return localctx
}

// IMaternitybenefitsscoredContext is an interface to support dynamic dispatch.
type IMaternitybenefitsscoredContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaternitybenefitsscoredContext differentiates from other interfaces.
	IsMaternitybenefitsscoredContext()
}

type MaternitybenefitsscoredContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaternitybenefitsscoredContext() *MaternitybenefitsscoredContext {
	var p = new(MaternitybenefitsscoredContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MaternityBenefitsScoredParserRULE_maternitybenefitsscored
	return p
}

func (*MaternitybenefitsscoredContext) IsMaternitybenefitsscoredContext() {}

func NewMaternitybenefitsscoredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaternitybenefitsscoredContext {
	var p = new(MaternitybenefitsscoredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaternityBenefitsScoredParserRULE_maternitybenefitsscored

	return p
}

func (s *MaternitybenefitsscoredContext) GetParser() antlr.Parser { return s.parser }

func (s *MaternitybenefitsscoredContext) MATERNITYBENEFITSSCORED() antlr.TerminalNode {
	return s.GetToken(MaternityBenefitsScoredParserMATERNITYBENEFITSSCORED, 0)
}

func (s *MaternitybenefitsscoredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaternitybenefitsscoredContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaternitybenefitsscoredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaternityBenefitsScoredListener); ok {
		listenerT.EnterMaternitybenefitsscored(s)
	}
}

func (s *MaternitybenefitsscoredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaternityBenefitsScoredListener); ok {
		listenerT.ExitMaternitybenefitsscored(s)
	}
}

func (p *MaternityBenefitsScoredParser) Maternitybenefitsscored() (localctx IMaternitybenefitsscoredContext) {
	this := p
	_ = this

	localctx = NewMaternitybenefitsscoredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MaternityBenefitsScoredParserRULE_maternitybenefitsscored)

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
		p.Match(MaternityBenefitsScoredParserMATERNITYBENEFITSSCORED)
	}

	return localctx
}
