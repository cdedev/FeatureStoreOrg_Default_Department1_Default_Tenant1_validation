// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072357370/Sociocultural.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sociocultural

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

type SocioculturalParser struct {
	*antlr.BaseParser
}

var socioculturalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func socioculturalParserInit() {
	staticData := &socioculturalParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SOCIOCULTURAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sociocultural",
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

// SocioculturalParserInit initializes any static state used to implement SocioculturalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSocioculturalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SocioculturalParserInit() {
	staticData := &socioculturalParserStaticData
	staticData.once.Do(socioculturalParserInit)
}

// NewSocioculturalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSocioculturalParser(input antlr.TokenStream) *SocioculturalParser {
	SocioculturalParserInit()
	this := new(SocioculturalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &socioculturalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sociocultural.g4"

	return this
}

// SocioculturalParser tokens.
const (
	SocioculturalParserEOF           = antlr.TokenEOF
	SocioculturalParserSOCIOCULTURAL = 1
	SocioculturalParserOWNKEY        = 2
	SocioculturalParserSPLITKEY      = 3
	SocioculturalParserWS            = 4
)

// SocioculturalParser rules.
const (
	SocioculturalParserRULE_expression    = 0
	SocioculturalParserRULE_sociocultural = 1
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
	p.RuleIndex = SocioculturalParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SocioculturalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sociocultural() ISocioculturalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISocioculturalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISocioculturalContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SocioculturalParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SocioculturalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SocioculturalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SocioculturalParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SocioculturalParserRULE_expression)

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
		p.Sociocultural()
	}
	{
		p.SetState(5)
		p.Match(SocioculturalParserEOF)
	}

	return localctx
}

// ISocioculturalContext is an interface to support dynamic dispatch.
type ISocioculturalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSocioculturalContext differentiates from other interfaces.
	IsSocioculturalContext()
}

type SocioculturalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySocioculturalContext() *SocioculturalContext {
	var p = new(SocioculturalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SocioculturalParserRULE_sociocultural
	return p
}

func (*SocioculturalContext) IsSocioculturalContext() {}

func NewSocioculturalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SocioculturalContext {
	var p = new(SocioculturalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SocioculturalParserRULE_sociocultural

	return p
}

func (s *SocioculturalContext) GetParser() antlr.Parser { return s.parser }

func (s *SocioculturalContext) SOCIOCULTURAL() antlr.TerminalNode {
	return s.GetToken(SocioculturalParserSOCIOCULTURAL, 0)
}

func (s *SocioculturalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SocioculturalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SocioculturalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SocioculturalListener); ok {
		listenerT.EnterSociocultural(s)
	}
}

func (s *SocioculturalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SocioculturalListener); ok {
		listenerT.ExitSociocultural(s)
	}
}

func (p *SocioculturalParser) Sociocultural() (localctx ISocioculturalContext) {
	this := p
	_ = this

	localctx = NewSocioculturalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SocioculturalParserRULE_sociocultural)

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
		p.Match(SocioculturalParserSOCIOCULTURAL)
	}

	return localctx
}
