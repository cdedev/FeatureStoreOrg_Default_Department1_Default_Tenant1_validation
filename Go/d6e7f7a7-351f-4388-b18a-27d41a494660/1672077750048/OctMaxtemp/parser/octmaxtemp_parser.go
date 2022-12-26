// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077750048/OctMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OctMaxtemp

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

type OctMaxtempParser struct {
	*antlr.BaseParser
}

var octmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func octmaxtempParserInit() {
	staticData := &octmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OCTMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "octmaxtemp",
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

// OctMaxtempParserInit initializes any static state used to implement OctMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOctMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OctMaxtempParserInit() {
	staticData := &octmaxtempParserStaticData
	staticData.once.Do(octmaxtempParserInit)
}

// NewOctMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOctMaxtempParser(input antlr.TokenStream) *OctMaxtempParser {
	OctMaxtempParserInit()
	this := new(OctMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &octmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "OctMaxtemp.g4"

	return this
}

// OctMaxtempParser tokens.
const (
	OctMaxtempParserEOF        = antlr.TokenEOF
	OctMaxtempParserOCTMAXTEMP = 1
	OctMaxtempParserOWNKEY     = 2
	OctMaxtempParserSPLITKEY   = 3
	OctMaxtempParserWS         = 4
)

// OctMaxtempParser rules.
const (
	OctMaxtempParserRULE_expression = 0
	OctMaxtempParserRULE_octmaxtemp = 1
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
	p.RuleIndex = OctMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OctMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Octmaxtemp() IOctmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOctmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOctmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OctMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OctMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OctMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OctMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OctMaxtempParserRULE_expression)

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
		p.Octmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(OctMaxtempParserEOF)
	}

	return localctx
}

// IOctmaxtempContext is an interface to support dynamic dispatch.
type IOctmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOctmaxtempContext differentiates from other interfaces.
	IsOctmaxtempContext()
}

type OctmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOctmaxtempContext() *OctmaxtempContext {
	var p = new(OctmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OctMaxtempParserRULE_octmaxtemp
	return p
}

func (*OctmaxtempContext) IsOctmaxtempContext() {}

func NewOctmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OctmaxtempContext {
	var p = new(OctmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OctMaxtempParserRULE_octmaxtemp

	return p
}

func (s *OctmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *OctmaxtempContext) OCTMAXTEMP() antlr.TerminalNode {
	return s.GetToken(OctMaxtempParserOCTMAXTEMP, 0)
}

func (s *OctmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OctmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OctMaxtempListener); ok {
		listenerT.EnterOctmaxtemp(s)
	}
}

func (s *OctmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OctMaxtempListener); ok {
		listenerT.ExitOctmaxtemp(s)
	}
}

func (p *OctMaxtempParser) Octmaxtemp() (localctx IOctmaxtempContext) {
	this := p
	_ = this

	localctx = NewOctmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OctMaxtempParserRULE_octmaxtemp)

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
		p.Match(OctMaxtempParserOCTMAXTEMP)
	}

	return localctx
}
