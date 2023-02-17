// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676609525407/Quality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quality

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

type QualityParser struct {
	*antlr.BaseParser
}

var qualityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func qualityParserInit() {
	staticData := &qualityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "QUALITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "quality",
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

// QualityParserInit initializes any static state used to implement QualityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQualityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QualityParserInit() {
	staticData := &qualityParserStaticData
	staticData.once.Do(qualityParserInit)
}

// NewQualityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQualityParser(input antlr.TokenStream) *QualityParser {
	QualityParserInit()
	this := new(QualityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &qualityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Quality.g4"

	return this
}

// QualityParser tokens.
const (
	QualityParserEOF      = antlr.TokenEOF
	QualityParserQUALITY  = 1
	QualityParserOWNKEY   = 2
	QualityParserSPLITKEY = 3
	QualityParserWS       = 4
)

// QualityParser rules.
const (
	QualityParserRULE_expression = 0
	QualityParserRULE_quality    = 1
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
	p.RuleIndex = QualityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QualityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Quality() IQualityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(QualityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QualityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QualityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *QualityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QualityParserRULE_expression)

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
		p.Quality()
	}
	{
		p.SetState(5)
		p.Match(QualityParserEOF)
	}

	return localctx
}

// IQualityContext is an interface to support dynamic dispatch.
type IQualityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualityContext differentiates from other interfaces.
	IsQualityContext()
}

type QualityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualityContext() *QualityContext {
	var p = new(QualityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QualityParserRULE_quality
	return p
}

func (*QualityContext) IsQualityContext() {}

func NewQualityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualityContext {
	var p = new(QualityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QualityParserRULE_quality

	return p
}

func (s *QualityContext) GetParser() antlr.Parser { return s.parser }

func (s *QualityContext) QUALITY() antlr.TerminalNode {
	return s.GetToken(QualityParserQUALITY, 0)
}

func (s *QualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QualityListener); ok {
		listenerT.EnterQuality(s)
	}
}

func (s *QualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QualityListener); ok {
		listenerT.ExitQuality(s)
	}
}

func (p *QualityParser) Quality() (localctx IQualityContext) {
	this := p
	_ = this

	localctx = NewQualityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QualityParserRULE_quality)

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
		p.Match(QualityParserQUALITY)
	}

	return localctx
}
