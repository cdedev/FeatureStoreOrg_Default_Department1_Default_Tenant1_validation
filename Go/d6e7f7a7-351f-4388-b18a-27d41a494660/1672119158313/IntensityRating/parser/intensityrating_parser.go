// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119158313/IntensityRating.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntensityRating

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

type IntensityRatingParser struct {
	*antlr.BaseParser
}

var intensityratingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func intensityratingParserInit() {
	staticData := &intensityratingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INTENSITYRATING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "intensityrating",
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

// IntensityRatingParserInit initializes any static state used to implement IntensityRatingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIntensityRatingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IntensityRatingParserInit() {
	staticData := &intensityratingParserStaticData
	staticData.once.Do(intensityratingParserInit)
}

// NewIntensityRatingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIntensityRatingParser(input antlr.TokenStream) *IntensityRatingParser {
	IntensityRatingParserInit()
	this := new(IntensityRatingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &intensityratingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "IntensityRating.g4"

	return this
}

// IntensityRatingParser tokens.
const (
	IntensityRatingParserEOF             = antlr.TokenEOF
	IntensityRatingParserINTENSITYRATING = 1
	IntensityRatingParserOWNKEY          = 2
	IntensityRatingParserSPLITKEY        = 3
	IntensityRatingParserWS              = 4
)

// IntensityRatingParser rules.
const (
	IntensityRatingParserRULE_expression      = 0
	IntensityRatingParserRULE_intensityrating = 1
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
	p.RuleIndex = IntensityRatingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntensityRatingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Intensityrating() IIntensityratingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntensityratingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntensityratingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IntensityRatingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntensityRatingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntensityRatingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IntensityRatingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IntensityRatingParserRULE_expression)

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
		p.Intensityrating()
	}
	{
		p.SetState(5)
		p.Match(IntensityRatingParserEOF)
	}

	return localctx
}

// IIntensityratingContext is an interface to support dynamic dispatch.
type IIntensityratingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntensityratingContext differentiates from other interfaces.
	IsIntensityratingContext()
}

type IntensityratingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntensityratingContext() *IntensityratingContext {
	var p = new(IntensityratingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IntensityRatingParserRULE_intensityrating
	return p
}

func (*IntensityratingContext) IsIntensityratingContext() {}

func NewIntensityratingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntensityratingContext {
	var p = new(IntensityratingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntensityRatingParserRULE_intensityrating

	return p
}

func (s *IntensityratingContext) GetParser() antlr.Parser { return s.parser }

func (s *IntensityratingContext) INTENSITYRATING() antlr.TerminalNode {
	return s.GetToken(IntensityRatingParserINTENSITYRATING, 0)
}

func (s *IntensityratingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntensityratingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntensityratingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntensityRatingListener); ok {
		listenerT.EnterIntensityrating(s)
	}
}

func (s *IntensityratingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntensityRatingListener); ok {
		listenerT.ExitIntensityrating(s)
	}
}

func (p *IntensityRatingParser) Intensityrating() (localctx IIntensityratingContext) {
	this := p
	_ = this

	localctx = NewIntensityratingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IntensityRatingParserRULE_intensityrating)

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
		p.Match(IntensityRatingParserINTENSITYRATING)
	}

	return localctx
}
