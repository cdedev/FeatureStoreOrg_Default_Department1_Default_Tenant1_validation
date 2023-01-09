// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238774423/Normal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Normal

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

type NormalParser struct {
	*antlr.BaseParser
}

var normalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func normalParserInit() {
	staticData := &normalParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NORMAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "normal",
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

// NormalParserInit initializes any static state used to implement NormalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNormalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NormalParserInit() {
	staticData := &normalParserStaticData
	staticData.once.Do(normalParserInit)
}

// NewNormalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNormalParser(input antlr.TokenStream) *NormalParser {
	NormalParserInit()
	this := new(NormalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &normalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Normal.g4"

	return this
}

// NormalParser tokens.
const (
	NormalParserEOF      = antlr.TokenEOF
	NormalParserNORMAL   = 1
	NormalParserOWNKEY   = 2
	NormalParserSPLITKEY = 3
	NormalParserWS       = 4
)

// NormalParser rules.
const (
	NormalParserRULE_expression = 0
	NormalParserRULE_normal     = 1
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
	p.RuleIndex = NormalParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NormalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Normal() INormalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INormalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INormalContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NormalParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NormalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NormalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NormalParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NormalParserRULE_expression)

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
		p.Normal()
	}
	{
		p.SetState(5)
		p.Match(NormalParserEOF)
	}

	return localctx
}

// INormalContext is an interface to support dynamic dispatch.
type INormalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNormalContext differentiates from other interfaces.
	IsNormalContext()
}

type NormalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormalContext() *NormalContext {
	var p = new(NormalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NormalParserRULE_normal
	return p
}

func (*NormalContext) IsNormalContext() {}

func NewNormalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalContext {
	var p = new(NormalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NormalParserRULE_normal

	return p
}

func (s *NormalContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalContext) NORMAL() antlr.TerminalNode {
	return s.GetToken(NormalParserNORMAL, 0)
}

func (s *NormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NormalListener); ok {
		listenerT.EnterNormal(s)
	}
}

func (s *NormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NormalListener); ok {
		listenerT.ExitNormal(s)
	}
}

func (p *NormalParser) Normal() (localctx INormalContext) {
	this := p
	_ = this

	localctx = NewNormalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NormalParserRULE_normal)

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
		p.Match(NormalParserNORMAL)
	}

	return localctx
}
