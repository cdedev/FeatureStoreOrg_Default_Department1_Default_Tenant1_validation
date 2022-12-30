// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377742445/NCores.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NCores

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

type NCoresParser struct {
	*antlr.BaseParser
}

var ncoresParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ncoresParserInit() {
	staticData := &ncoresParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NCORES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ncores",
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

// NCoresParserInit initializes any static state used to implement NCoresParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNCoresParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NCoresParserInit() {
	staticData := &ncoresParserStaticData
	staticData.once.Do(ncoresParserInit)
}

// NewNCoresParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNCoresParser(input antlr.TokenStream) *NCoresParser {
	NCoresParserInit()
	this := new(NCoresParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ncoresParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NCores.g4"

	return this
}

// NCoresParser tokens.
const (
	NCoresParserEOF      = antlr.TokenEOF
	NCoresParserNCORES   = 1
	NCoresParserOWNKEY   = 2
	NCoresParserSPLITKEY = 3
	NCoresParserWS       = 4
)

// NCoresParser rules.
const (
	NCoresParserRULE_expression = 0
	NCoresParserRULE_ncores     = 1
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
	p.RuleIndex = NCoresParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NCoresParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ncores() INcoresContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INcoresContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INcoresContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NCoresParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NCoresListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NCoresListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NCoresParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NCoresParserRULE_expression)

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
		p.Ncores()
	}
	{
		p.SetState(5)
		p.Match(NCoresParserEOF)
	}

	return localctx
}

// INcoresContext is an interface to support dynamic dispatch.
type INcoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNcoresContext differentiates from other interfaces.
	IsNcoresContext()
}

type NcoresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNcoresContext() *NcoresContext {
	var p = new(NcoresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NCoresParserRULE_ncores
	return p
}

func (*NcoresContext) IsNcoresContext() {}

func NewNcoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NcoresContext {
	var p = new(NcoresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NCoresParserRULE_ncores

	return p
}

func (s *NcoresContext) GetParser() antlr.Parser { return s.parser }

func (s *NcoresContext) NCORES() antlr.TerminalNode {
	return s.GetToken(NCoresParserNCORES, 0)
}

func (s *NcoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NcoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NcoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NCoresListener); ok {
		listenerT.EnterNcores(s)
	}
}

func (s *NcoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NCoresListener); ok {
		listenerT.ExitNcores(s)
	}
}

func (p *NCoresParser) Ncores() (localctx INcoresContext) {
	this := p
	_ = this

	localctx = NewNcoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NCoresParserRULE_ncores)

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
		p.Match(NCoresParserNCORES)
	}

	return localctx
}
