// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671866796744/CloneSize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CloneSize

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

type CloneSizeParser struct {
	*antlr.BaseParser
}

var clonesizeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clonesizeParserInit() {
	staticData := &clonesizeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLONESIZE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "clonesize",
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

// CloneSizeParserInit initializes any static state used to implement CloneSizeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCloneSizeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CloneSizeParserInit() {
	staticData := &clonesizeParserStaticData
	staticData.once.Do(clonesizeParserInit)
}

// NewCloneSizeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCloneSizeParser(input antlr.TokenStream) *CloneSizeParser {
	CloneSizeParserInit()
	this := new(CloneSizeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clonesizeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CloneSize.g4"

	return this
}

// CloneSizeParser tokens.
const (
	CloneSizeParserEOF       = antlr.TokenEOF
	CloneSizeParserCLONESIZE = 1
	CloneSizeParserOWNKEY    = 2
	CloneSizeParserSPLITKEY  = 3
	CloneSizeParserWS        = 4
)

// CloneSizeParser rules.
const (
	CloneSizeParserRULE_expression = 0
	CloneSizeParserRULE_clonesize  = 1
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
	p.RuleIndex = CloneSizeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CloneSizeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Clonesize() IClonesizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClonesizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClonesizeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CloneSizeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloneSizeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloneSizeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CloneSizeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CloneSizeParserRULE_expression)

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
		p.Clonesize()
	}
	{
		p.SetState(5)
		p.Match(CloneSizeParserEOF)
	}

	return localctx
}

// IClonesizeContext is an interface to support dynamic dispatch.
type IClonesizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClonesizeContext differentiates from other interfaces.
	IsClonesizeContext()
}

type ClonesizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClonesizeContext() *ClonesizeContext {
	var p = new(ClonesizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CloneSizeParserRULE_clonesize
	return p
}

func (*ClonesizeContext) IsClonesizeContext() {}

func NewClonesizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClonesizeContext {
	var p = new(ClonesizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CloneSizeParserRULE_clonesize

	return p
}

func (s *ClonesizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClonesizeContext) CLONESIZE() antlr.TerminalNode {
	return s.GetToken(CloneSizeParserCLONESIZE, 0)
}

func (s *ClonesizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClonesizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClonesizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloneSizeListener); ok {
		listenerT.EnterClonesize(s)
	}
}

func (s *ClonesizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloneSizeListener); ok {
		listenerT.ExitClonesize(s)
	}
}

func (p *CloneSizeParser) Clonesize() (localctx IClonesizeContext) {
	this := p
	_ = this

	localctx = NewClonesizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CloneSizeParserRULE_clonesize)

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
		p.Match(CloneSizeParserCLONESIZE)
	}

	return localctx
}
