// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675922679725/Curtosis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Curtosis

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

type CurtosisParser struct {
	*antlr.BaseParser
}

var curtosisParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func curtosisParserInit() {
	staticData := &curtosisParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CURTOSIS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "curtosis",
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

// CurtosisParserInit initializes any static state used to implement CurtosisParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCurtosisParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CurtosisParserInit() {
	staticData := &curtosisParserStaticData
	staticData.once.Do(curtosisParserInit)
}

// NewCurtosisParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCurtosisParser(input antlr.TokenStream) *CurtosisParser {
	CurtosisParserInit()
	this := new(CurtosisParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &curtosisParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Curtosis.g4"

	return this
}

// CurtosisParser tokens.
const (
	CurtosisParserEOF      = antlr.TokenEOF
	CurtosisParserCURTOSIS = 1
	CurtosisParserOWNKEY   = 2
	CurtosisParserSPLITKEY = 3
	CurtosisParserWS       = 4
)

// CurtosisParser rules.
const (
	CurtosisParserRULE_expression = 0
	CurtosisParserRULE_curtosis   = 1
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
	p.RuleIndex = CurtosisParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CurtosisParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Curtosis() ICurtosisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurtosisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurtosisContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CurtosisParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurtosisListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurtosisListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CurtosisParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CurtosisParserRULE_expression)

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
		p.Curtosis()
	}
	{
		p.SetState(5)
		p.Match(CurtosisParserEOF)
	}

	return localctx
}

// ICurtosisContext is an interface to support dynamic dispatch.
type ICurtosisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurtosisContext differentiates from other interfaces.
	IsCurtosisContext()
}

type CurtosisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurtosisContext() *CurtosisContext {
	var p = new(CurtosisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CurtosisParserRULE_curtosis
	return p
}

func (*CurtosisContext) IsCurtosisContext() {}

func NewCurtosisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurtosisContext {
	var p = new(CurtosisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CurtosisParserRULE_curtosis

	return p
}

func (s *CurtosisContext) GetParser() antlr.Parser { return s.parser }

func (s *CurtosisContext) CURTOSIS() antlr.TerminalNode {
	return s.GetToken(CurtosisParserCURTOSIS, 0)
}

func (s *CurtosisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurtosisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurtosisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurtosisListener); ok {
		listenerT.EnterCurtosis(s)
	}
}

func (s *CurtosisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurtosisListener); ok {
		listenerT.ExitCurtosis(s)
	}
}

func (p *CurtosisParser) Curtosis() (localctx ICurtosisContext) {
	this := p
	_ = this

	localctx = NewCurtosisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CurtosisParserRULE_curtosis)

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
		p.Match(CurtosisParserCURTOSIS)
	}

	return localctx
}
