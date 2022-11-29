// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669698121341/Sulphurdioxide.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sulphurdioxide

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

type SulphurdioxideParser struct {
	*antlr.BaseParser
}

var sulphurdioxideParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sulphurdioxideParserInit() {
	staticData := &sulphurdioxideParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SULPHURDIOXIDE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sulphurdioxide",
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

// SulphurdioxideParserInit initializes any static state used to implement SulphurdioxideParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSulphurdioxideParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SulphurdioxideParserInit() {
	staticData := &sulphurdioxideParserStaticData
	staticData.once.Do(sulphurdioxideParserInit)
}

// NewSulphurdioxideParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSulphurdioxideParser(input antlr.TokenStream) *SulphurdioxideParser {
	SulphurdioxideParserInit()
	this := new(SulphurdioxideParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sulphurdioxideParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sulphurdioxide.g4"

	return this
}

// SulphurdioxideParser tokens.
const (
	SulphurdioxideParserEOF            = antlr.TokenEOF
	SulphurdioxideParserSULPHURDIOXIDE = 1
	SulphurdioxideParserOWNKEY         = 2
	SulphurdioxideParserSPLITKEY       = 3
	SulphurdioxideParserWS             = 4
)

// SulphurdioxideParser rules.
const (
	SulphurdioxideParserRULE_expression     = 0
	SulphurdioxideParserRULE_sulphurdioxide = 1
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
	p.RuleIndex = SulphurdioxideParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SulphurdioxideParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sulphurdioxide() ISulphurdioxideContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISulphurdioxideContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISulphurdioxideContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SulphurdioxideParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SulphurdioxideListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SulphurdioxideListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SulphurdioxideParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SulphurdioxideParserRULE_expression)

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
		p.Sulphurdioxide()
	}
	{
		p.SetState(5)
		p.Match(SulphurdioxideParserEOF)
	}

	return localctx
}

// ISulphurdioxideContext is an interface to support dynamic dispatch.
type ISulphurdioxideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSulphurdioxideContext differentiates from other interfaces.
	IsSulphurdioxideContext()
}

type SulphurdioxideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySulphurdioxideContext() *SulphurdioxideContext {
	var p = new(SulphurdioxideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SulphurdioxideParserRULE_sulphurdioxide
	return p
}

func (*SulphurdioxideContext) IsSulphurdioxideContext() {}

func NewSulphurdioxideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SulphurdioxideContext {
	var p = new(SulphurdioxideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SulphurdioxideParserRULE_sulphurdioxide

	return p
}

func (s *SulphurdioxideContext) GetParser() antlr.Parser { return s.parser }

func (s *SulphurdioxideContext) SULPHURDIOXIDE() antlr.TerminalNode {
	return s.GetToken(SulphurdioxideParserSULPHURDIOXIDE, 0)
}

func (s *SulphurdioxideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SulphurdioxideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SulphurdioxideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SulphurdioxideListener); ok {
		listenerT.EnterSulphurdioxide(s)
	}
}

func (s *SulphurdioxideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SulphurdioxideListener); ok {
		listenerT.ExitSulphurdioxide(s)
	}
}

func (p *SulphurdioxideParser) Sulphurdioxide() (localctx ISulphurdioxideContext) {
	this := p
	_ = this

	localctx = NewSulphurdioxideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SulphurdioxideParserRULE_sulphurdioxide)

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
		p.Match(SulphurdioxideParserSULPHURDIOXIDE)
	}

	return localctx
}
