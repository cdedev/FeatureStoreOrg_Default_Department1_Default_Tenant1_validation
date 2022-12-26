// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077184093/MayMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MayMintemp

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

type MayMintempParser struct {
	*antlr.BaseParser
}

var maymintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func maymintempParserInit() {
	staticData := &maymintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MAYMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "maymintemp",
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

// MayMintempParserInit initializes any static state used to implement MayMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMayMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MayMintempParserInit() {
	staticData := &maymintempParserStaticData
	staticData.once.Do(maymintempParserInit)
}

// NewMayMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMayMintempParser(input antlr.TokenStream) *MayMintempParser {
	MayMintempParserInit()
	this := new(MayMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &maymintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MayMintemp.g4"

	return this
}

// MayMintempParser tokens.
const (
	MayMintempParserEOF        = antlr.TokenEOF
	MayMintempParserMAYMINTEMP = 1
	MayMintempParserOWNKEY     = 2
	MayMintempParserSPLITKEY   = 3
	MayMintempParserWS         = 4
)

// MayMintempParser rules.
const (
	MayMintempParserRULE_expression = 0
	MayMintempParserRULE_maymintemp = 1
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
	p.RuleIndex = MayMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MayMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Maymintemp() IMaymintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaymintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaymintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MayMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MayMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MayMintempParserRULE_expression)

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
		p.Maymintemp()
	}
	{
		p.SetState(5)
		p.Match(MayMintempParserEOF)
	}

	return localctx
}

// IMaymintempContext is an interface to support dynamic dispatch.
type IMaymintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaymintempContext differentiates from other interfaces.
	IsMaymintempContext()
}

type MaymintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaymintempContext() *MaymintempContext {
	var p = new(MaymintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MayMintempParserRULE_maymintemp
	return p
}

func (*MaymintempContext) IsMaymintempContext() {}

func NewMaymintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaymintempContext {
	var p = new(MaymintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MayMintempParserRULE_maymintemp

	return p
}

func (s *MaymintempContext) GetParser() antlr.Parser { return s.parser }

func (s *MaymintempContext) MAYMINTEMP() antlr.TerminalNode {
	return s.GetToken(MayMintempParserMAYMINTEMP, 0)
}

func (s *MaymintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaymintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaymintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMintempListener); ok {
		listenerT.EnterMaymintemp(s)
	}
}

func (s *MaymintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MayMintempListener); ok {
		listenerT.ExitMaymintemp(s)
	}
}

func (p *MayMintempParser) Maymintemp() (localctx IMaymintempContext) {
	this := p
	_ = this

	localctx = NewMaymintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MayMintempParserRULE_maymintemp)

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
		p.Match(MayMintempParserMAYMINTEMP)
	}

	return localctx
}
