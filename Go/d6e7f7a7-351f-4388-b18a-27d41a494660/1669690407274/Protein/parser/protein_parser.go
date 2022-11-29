// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690407274/Protein.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Protein

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

type ProteinParser struct {
	*antlr.BaseParser
}

var proteinParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func proteinParserInit() {
	staticData := &proteinParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PROTEIN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "protein",
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

// ProteinParserInit initializes any static state used to implement ProteinParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProteinParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProteinParserInit() {
	staticData := &proteinParserStaticData
	staticData.once.Do(proteinParserInit)
}

// NewProteinParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProteinParser(input antlr.TokenStream) *ProteinParser {
	ProteinParserInit()
	this := new(ProteinParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &proteinParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Protein.g4"

	return this
}

// ProteinParser tokens.
const (
	ProteinParserEOF      = antlr.TokenEOF
	ProteinParserPROTEIN  = 1
	ProteinParserOWNKEY   = 2
	ProteinParserSPLITKEY = 3
	ProteinParserWS       = 4
)

// ProteinParser rules.
const (
	ProteinParserRULE_expression = 0
	ProteinParserRULE_protein    = 1
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
	p.RuleIndex = ProteinParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProteinParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Protein() IProteinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProteinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProteinContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProteinParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProteinListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProteinListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ProteinParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProteinParserRULE_expression)

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
		p.Protein()
	}
	{
		p.SetState(5)
		p.Match(ProteinParserEOF)
	}

	return localctx
}

// IProteinContext is an interface to support dynamic dispatch.
type IProteinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProteinContext differentiates from other interfaces.
	IsProteinContext()
}

type ProteinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProteinContext() *ProteinContext {
	var p = new(ProteinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProteinParserRULE_protein
	return p
}

func (*ProteinContext) IsProteinContext() {}

func NewProteinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProteinContext {
	var p = new(ProteinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProteinParserRULE_protein

	return p
}

func (s *ProteinContext) GetParser() antlr.Parser { return s.parser }

func (s *ProteinContext) PROTEIN() antlr.TerminalNode {
	return s.GetToken(ProteinParserPROTEIN, 0)
}

func (s *ProteinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProteinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProteinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProteinListener); ok {
		listenerT.EnterProtein(s)
	}
}

func (s *ProteinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProteinListener); ok {
		listenerT.ExitProtein(s)
	}
}

func (p *ProteinParser) Protein() (localctx IProteinContext) {
	this := p
	_ = this

	localctx = NewProteinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProteinParserRULE_protein)

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
		p.Match(ProteinParserPROTEIN)
	}

	return localctx
}
