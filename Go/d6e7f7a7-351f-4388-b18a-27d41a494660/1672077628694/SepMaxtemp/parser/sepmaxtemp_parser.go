// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077628694/SepMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SepMaxtemp

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

type SepMaxtempParser struct {
	*antlr.BaseParser
}

var sepmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sepmaxtempParserInit() {
	staticData := &sepmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SEPMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sepmaxtemp",
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

// SepMaxtempParserInit initializes any static state used to implement SepMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSepMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SepMaxtempParserInit() {
	staticData := &sepmaxtempParserStaticData
	staticData.once.Do(sepmaxtempParserInit)
}

// NewSepMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSepMaxtempParser(input antlr.TokenStream) *SepMaxtempParser {
	SepMaxtempParserInit()
	this := new(SepMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sepmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SepMaxtemp.g4"

	return this
}

// SepMaxtempParser tokens.
const (
	SepMaxtempParserEOF        = antlr.TokenEOF
	SepMaxtempParserSEPMAXTEMP = 1
	SepMaxtempParserOWNKEY     = 2
	SepMaxtempParserSPLITKEY   = 3
	SepMaxtempParserWS         = 4
)

// SepMaxtempParser rules.
const (
	SepMaxtempParserRULE_expression = 0
	SepMaxtempParserRULE_sepmaxtemp = 1
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
	p.RuleIndex = SepMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SepMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sepmaxtemp() ISepmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SepMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SepMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SepMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SepMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SepMaxtempParserRULE_expression)

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
		p.Sepmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(SepMaxtempParserEOF)
	}

	return localctx
}

// ISepmaxtempContext is an interface to support dynamic dispatch.
type ISepmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSepmaxtempContext differentiates from other interfaces.
	IsSepmaxtempContext()
}

type SepmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySepmaxtempContext() *SepmaxtempContext {
	var p = new(SepmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SepMaxtempParserRULE_sepmaxtemp
	return p
}

func (*SepmaxtempContext) IsSepmaxtempContext() {}

func NewSepmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SepmaxtempContext {
	var p = new(SepmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SepMaxtempParserRULE_sepmaxtemp

	return p
}

func (s *SepmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *SepmaxtempContext) SEPMAXTEMP() antlr.TerminalNode {
	return s.GetToken(SepMaxtempParserSEPMAXTEMP, 0)
}

func (s *SepmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SepmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SepmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SepMaxtempListener); ok {
		listenerT.EnterSepmaxtemp(s)
	}
}

func (s *SepmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SepMaxtempListener); ok {
		listenerT.ExitSepmaxtemp(s)
	}
}

func (p *SepMaxtempParser) Sepmaxtemp() (localctx ISepmaxtempContext) {
	this := p
	_ = this

	localctx = NewSepmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SepMaxtempParserRULE_sepmaxtemp)

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
		p.Match(SepMaxtempParserSEPMAXTEMP)
	}

	return localctx
}
