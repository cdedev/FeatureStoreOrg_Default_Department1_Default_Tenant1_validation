// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637572205/Documents.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Documents

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

type DocumentsParser struct {
	*antlr.BaseParser
}

var documentsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func documentsParserInit() {
	staticData := &documentsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DOCUMENTS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "documents",
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

// DocumentsParserInit initializes any static state used to implement DocumentsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDocumentsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DocumentsParserInit() {
	staticData := &documentsParserStaticData
	staticData.once.Do(documentsParserInit)
}

// NewDocumentsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDocumentsParser(input antlr.TokenStream) *DocumentsParser {
	DocumentsParserInit()
	this := new(DocumentsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &documentsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Documents.g4"

	return this
}

// DocumentsParser tokens.
const (
	DocumentsParserEOF       = antlr.TokenEOF
	DocumentsParserDOCUMENTS = 1
	DocumentsParserOWNKEY    = 2
	DocumentsParserSPLITKEY  = 3
	DocumentsParserWS        = 4
)

// DocumentsParser rules.
const (
	DocumentsParserRULE_expression = 0
	DocumentsParserRULE_documents  = 1
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
	p.RuleIndex = DocumentsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DocumentsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Documents() IDocumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocumentsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DocumentsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocumentsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocumentsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DocumentsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DocumentsParserRULE_expression)

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
		p.Documents()
	}
	{
		p.SetState(5)
		p.Match(DocumentsParserEOF)
	}

	return localctx
}

// IDocumentsContext is an interface to support dynamic dispatch.
type IDocumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentsContext differentiates from other interfaces.
	IsDocumentsContext()
}

type DocumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentsContext() *DocumentsContext {
	var p = new(DocumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DocumentsParserRULE_documents
	return p
}

func (*DocumentsContext) IsDocumentsContext() {}

func NewDocumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentsContext {
	var p = new(DocumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DocumentsParserRULE_documents

	return p
}

func (s *DocumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentsContext) DOCUMENTS() antlr.TerminalNode {
	return s.GetToken(DocumentsParserDOCUMENTS, 0)
}

func (s *DocumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocumentsListener); ok {
		listenerT.EnterDocuments(s)
	}
}

func (s *DocumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocumentsListener); ok {
		listenerT.ExitDocuments(s)
	}
}

func (p *DocumentsParser) Documents() (localctx IDocumentsContext) {
	this := p
	_ = this

	localctx = NewDocumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DocumentsParserRULE_documents)

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
		p.Match(DocumentsParserDOCUMENTS)
	}

	return localctx
}
