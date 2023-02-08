// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834395460/Sector.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sector

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

type SectorParser struct {
	*antlr.BaseParser
}

var sectorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sectorParserInit() {
	staticData := &sectorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SECTOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sector",
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

// SectorParserInit initializes any static state used to implement SectorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSectorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SectorParserInit() {
	staticData := &sectorParserStaticData
	staticData.once.Do(sectorParserInit)
}

// NewSectorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSectorParser(input antlr.TokenStream) *SectorParser {
	SectorParserInit()
	this := new(SectorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sectorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sector.g4"

	return this
}

// SectorParser tokens.
const (
	SectorParserEOF      = antlr.TokenEOF
	SectorParserSECTOR   = 1
	SectorParserOWNKEY   = 2
	SectorParserSPLITKEY = 3
	SectorParserWS       = 4
)

// SectorParser rules.
const (
	SectorParserRULE_expression = 0
	SectorParserRULE_sector     = 1
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
	p.RuleIndex = SectorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SectorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sector() ISectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SectorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SectorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SectorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SectorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SectorParserRULE_expression)

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
		p.Sector()
	}
	{
		p.SetState(5)
		p.Match(SectorParserEOF)
	}

	return localctx
}

// ISectorContext is an interface to support dynamic dispatch.
type ISectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectorContext differentiates from other interfaces.
	IsSectorContext()
}

type SectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectorContext() *SectorContext {
	var p = new(SectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SectorParserRULE_sector
	return p
}

func (*SectorContext) IsSectorContext() {}

func NewSectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectorContext {
	var p = new(SectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SectorParserRULE_sector

	return p
}

func (s *SectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SectorContext) SECTOR() antlr.TerminalNode {
	return s.GetToken(SectorParserSECTOR, 0)
}

func (s *SectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SectorListener); ok {
		listenerT.EnterSector(s)
	}
}

func (s *SectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SectorListener); ok {
		listenerT.ExitSector(s)
	}
}

func (p *SectorParser) Sector() (localctx ISectorContext) {
	this := p
	_ = this

	localctx = NewSectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SectorParserRULE_sector)

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
		p.Match(SectorParserSECTOR)
	}

	return localctx
}
