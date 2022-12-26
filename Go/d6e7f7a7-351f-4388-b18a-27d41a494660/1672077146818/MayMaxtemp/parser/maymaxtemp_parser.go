// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077146818/MayMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MayMaxtemp

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

type MayMaxtempParser struct {
	*antlr.BaseParser
}

var maymaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func maymaxtempParserInit() {
	staticData := &maymaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MAYMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "maymaxtemp",
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

// MayMaxtempParserInit initializes any static state used to implement MayMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMayMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MayMaxtempParserInit() {
	staticData := &maymaxtempParserStaticData
	staticData.once.Do(maymaxtempParserInit)
}

// NewMayMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMayMaxtempParser(input antlr.TokenStream) *MayMaxtempParser {
	MayMaxtempParserInit()
	this := new(MayMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &maymaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MayMaxtemp.g4"

	return this
}

// MayMaxtempParser tokens.
const (
	MayMaxtempParserEOF        = antlr.TokenEOF
	MayMaxtempParserMAYMAXTEMP = 1
	MayMaxtempParserOWNKEY     = 2
	MayMaxtempParserSPLITKEY   = 3
	MayMaxtempParserWS         = 4
)

// MayMaxtempParser rules.
const (
	MayMaxtempParserRULE_expression = 0
	MayMaxtempParserRULE_maymaxtemp = 1
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
	p.RuleIndex = MayMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MayMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Maymaxtemp() IMaymaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaymaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaymaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MayMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MayMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MayMaxtempParserRULE_expression)

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
		p.Maymaxtemp()
	}
	{
		p.SetState(5)
		p.Match(MayMaxtempParserEOF)
	}

	return localctx
}

// IMaymaxtempContext is an interface to support dynamic dispatch.
type IMaymaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaymaxtempContext differentiates from other interfaces.
	IsMaymaxtempContext()
}

type MaymaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaymaxtempContext() *MaymaxtempContext {
	var p = new(MaymaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MayMaxtempParserRULE_maymaxtemp
	return p
}

func (*MaymaxtempContext) IsMaymaxtempContext() {}

func NewMaymaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaymaxtempContext {
	var p = new(MaymaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MayMaxtempParserRULE_maymaxtemp

	return p
}

func (s *MaymaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *MaymaxtempContext) MAYMAXTEMP() antlr.TerminalNode {
	return s.GetToken(MayMaxtempParserMAYMAXTEMP, 0)
}

func (s *MaymaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaymaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaymaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMaxtempListener); ok {
		listenerT.EnterMaymaxtemp(s)
	}
}

func (s *MaymaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMaxtempListener); ok {
		listenerT.ExitMaymaxtemp(s)
	}
}

func (p *MayMaxtempParser) Maymaxtemp() (localctx IMaymaxtempContext) {
	this := p
	_ = this

	localctx = NewMaymaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MayMaxtempParserRULE_maymaxtemp)

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
		p.Match(MayMaxtempParserMAYMAXTEMP)
	}

	return localctx
}
