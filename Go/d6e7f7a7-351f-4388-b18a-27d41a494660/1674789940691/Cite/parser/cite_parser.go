// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789940691/Cite.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cite

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

type CiteParser struct {
	*antlr.BaseParser
}

var citeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func citeParserInit() {
	staticData := &citeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CITE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cite",
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

// CiteParserInit initializes any static state used to implement CiteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCiteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CiteParserInit() {
	staticData := &citeParserStaticData
	staticData.once.Do(citeParserInit)
}

// NewCiteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCiteParser(input antlr.TokenStream) *CiteParser {
	CiteParserInit()
	this := new(CiteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &citeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cite.g4"

	return this
}

// CiteParser tokens.
const (
	CiteParserEOF      = antlr.TokenEOF
	CiteParserCITE     = 1
	CiteParserOWNKEY   = 2
	CiteParserSPLITKEY = 3
	CiteParserWS       = 4
)

// CiteParser rules.
const (
	CiteParserRULE_expression = 0
	CiteParserRULE_cite       = 1
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
	p.RuleIndex = CiteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CiteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cite() ICiteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICiteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICiteContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CiteParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CiteListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CiteListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CiteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CiteParserRULE_expression)

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
		p.Cite()
	}
	{
		p.SetState(5)
		p.Match(CiteParserEOF)
	}

	return localctx
}

// ICiteContext is an interface to support dynamic dispatch.
type ICiteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCiteContext differentiates from other interfaces.
	IsCiteContext()
}

type CiteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCiteContext() *CiteContext {
	var p = new(CiteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CiteParserRULE_cite
	return p
}

func (*CiteContext) IsCiteContext() {}

func NewCiteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CiteContext {
	var p = new(CiteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CiteParserRULE_cite

	return p
}

func (s *CiteContext) GetParser() antlr.Parser { return s.parser }

func (s *CiteContext) CITE() antlr.TerminalNode {
	return s.GetToken(CiteParserCITE, 0)
}

func (s *CiteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CiteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CiteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CiteListener); ok {
		listenerT.EnterCite(s)
	}
}

func (s *CiteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CiteListener); ok {
		listenerT.ExitCite(s)
	}
}

func (p *CiteParser) Cite() (localctx ICiteContext) {
	this := p
	_ = this

	localctx = NewCiteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CiteParserRULE_cite)

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
		p.Match(CiteParserCITE)
	}

	return localctx
}
