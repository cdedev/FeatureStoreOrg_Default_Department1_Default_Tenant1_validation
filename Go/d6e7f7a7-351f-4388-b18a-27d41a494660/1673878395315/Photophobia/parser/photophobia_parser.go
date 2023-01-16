// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878395315/Photophobia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Photophobia

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

type PhotophobiaParser struct {
	*antlr.BaseParser
}

var photophobiaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func photophobiaParserInit() {
	staticData := &photophobiaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PHOTOPHOBIA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "photophobia",
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

// PhotophobiaParserInit initializes any static state used to implement PhotophobiaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPhotophobiaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PhotophobiaParserInit() {
	staticData := &photophobiaParserStaticData
	staticData.once.Do(photophobiaParserInit)
}

// NewPhotophobiaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPhotophobiaParser(input antlr.TokenStream) *PhotophobiaParser {
	PhotophobiaParserInit()
	this := new(PhotophobiaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &photophobiaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Photophobia.g4"

	return this
}

// PhotophobiaParser tokens.
const (
	PhotophobiaParserEOF         = antlr.TokenEOF
	PhotophobiaParserPHOTOPHOBIA = 1
	PhotophobiaParserOWNKEY      = 2
	PhotophobiaParserSPLITKEY    = 3
	PhotophobiaParserWS          = 4
)

// PhotophobiaParser rules.
const (
	PhotophobiaParserRULE_expression  = 0
	PhotophobiaParserRULE_photophobia = 1
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
	p.RuleIndex = PhotophobiaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhotophobiaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Photophobia() IPhotophobiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhotophobiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhotophobiaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PhotophobiaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhotophobiaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhotophobiaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PhotophobiaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PhotophobiaParserRULE_expression)

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
		p.Photophobia()
	}
	{
		p.SetState(5)
		p.Match(PhotophobiaParserEOF)
	}

	return localctx
}

// IPhotophobiaContext is an interface to support dynamic dispatch.
type IPhotophobiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhotophobiaContext differentiates from other interfaces.
	IsPhotophobiaContext()
}

type PhotophobiaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhotophobiaContext() *PhotophobiaContext {
	var p = new(PhotophobiaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PhotophobiaParserRULE_photophobia
	return p
}

func (*PhotophobiaContext) IsPhotophobiaContext() {}

func NewPhotophobiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhotophobiaContext {
	var p = new(PhotophobiaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhotophobiaParserRULE_photophobia

	return p
}

func (s *PhotophobiaContext) GetParser() antlr.Parser { return s.parser }

func (s *PhotophobiaContext) PHOTOPHOBIA() antlr.TerminalNode {
	return s.GetToken(PhotophobiaParserPHOTOPHOBIA, 0)
}

func (s *PhotophobiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhotophobiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhotophobiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhotophobiaListener); ok {
		listenerT.EnterPhotophobia(s)
	}
}

func (s *PhotophobiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhotophobiaListener); ok {
		listenerT.ExitPhotophobia(s)
	}
}

func (p *PhotophobiaParser) Photophobia() (localctx IPhotophobiaContext) {
	this := p
	_ = this

	localctx = NewPhotophobiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PhotophobiaParserRULE_photophobia)

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
		p.Match(PhotophobiaParserPHOTOPHOBIA)
	}

	return localctx
}
