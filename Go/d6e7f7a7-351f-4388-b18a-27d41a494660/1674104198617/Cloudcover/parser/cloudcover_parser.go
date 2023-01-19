// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674104198617/Cloudcover.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cloudcover

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

type CloudcoverParser struct {
	*antlr.BaseParser
}

var cloudcoverParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cloudcoverParserInit() {
	staticData := &cloudcoverParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLOUDCOVER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cloudcover",
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

// CloudcoverParserInit initializes any static state used to implement CloudcoverParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCloudcoverParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CloudcoverParserInit() {
	staticData := &cloudcoverParserStaticData
	staticData.once.Do(cloudcoverParserInit)
}

// NewCloudcoverParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCloudcoverParser(input antlr.TokenStream) *CloudcoverParser {
	CloudcoverParserInit()
	this := new(CloudcoverParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cloudcoverParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cloudcover.g4"

	return this
}

// CloudcoverParser tokens.
const (
	CloudcoverParserEOF        = antlr.TokenEOF
	CloudcoverParserCLOUDCOVER = 1
	CloudcoverParserOWNKEY     = 2
	CloudcoverParserSPLITKEY   = 3
	CloudcoverParserWS         = 4
)

// CloudcoverParser rules.
const (
	CloudcoverParserRULE_expression = 0
	CloudcoverParserRULE_cloudcover = 1
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
	p.RuleIndex = CloudcoverParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CloudcoverParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cloudcover() ICloudcoverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICloudcoverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICloudcoverContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CloudcoverParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloudcoverListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloudcoverListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CloudcoverParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CloudcoverParserRULE_expression)

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
		p.Cloudcover()
	}
	{
		p.SetState(5)
		p.Match(CloudcoverParserEOF)
	}

	return localctx
}

// ICloudcoverContext is an interface to support dynamic dispatch.
type ICloudcoverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCloudcoverContext differentiates from other interfaces.
	IsCloudcoverContext()
}

type CloudcoverContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloudcoverContext() *CloudcoverContext {
	var p = new(CloudcoverContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CloudcoverParserRULE_cloudcover
	return p
}

func (*CloudcoverContext) IsCloudcoverContext() {}

func NewCloudcoverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloudcoverContext {
	var p = new(CloudcoverContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CloudcoverParserRULE_cloudcover

	return p
}

func (s *CloudcoverContext) GetParser() antlr.Parser { return s.parser }

func (s *CloudcoverContext) CLOUDCOVER() antlr.TerminalNode {
	return s.GetToken(CloudcoverParserCLOUDCOVER, 0)
}

func (s *CloudcoverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloudcoverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CloudcoverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloudcoverListener); ok {
		listenerT.EnterCloudcover(s)
	}
}

func (s *CloudcoverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloudcoverListener); ok {
		listenerT.ExitCloudcover(s)
	}
}

func (p *CloudcoverParser) Cloudcover() (localctx ICloudcoverContext) {
	this := p
	_ = this

	localctx = NewCloudcoverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CloudcoverParserRULE_cloudcover)

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
		p.Match(CloudcoverParserCLOUDCOVER)
	}

	return localctx
}
