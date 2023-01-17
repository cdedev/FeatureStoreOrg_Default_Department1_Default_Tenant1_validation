// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928304634/Notes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Notes

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

type NotesParser struct {
	*antlr.BaseParser
}

var notesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func notesParserInit() {
	staticData := &notesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NOTES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "notes",
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

// NotesParserInit initializes any static state used to implement NotesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNotesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NotesParserInit() {
	staticData := &notesParserStaticData
	staticData.once.Do(notesParserInit)
}

// NewNotesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNotesParser(input antlr.TokenStream) *NotesParser {
	NotesParserInit()
	this := new(NotesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &notesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Notes.g4"

	return this
}

// NotesParser tokens.
const (
	NotesParserEOF      = antlr.TokenEOF
	NotesParserNOTES    = 1
	NotesParserOWNKEY   = 2
	NotesParserSPLITKEY = 3
	NotesParserWS       = 4
)

// NotesParser rules.
const (
	NotesParserRULE_expression = 0
	NotesParserRULE_notes      = 1
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
	p.RuleIndex = NotesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NotesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Notes() INotesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NotesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NotesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NotesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NotesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NotesParserRULE_expression)

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
		p.Notes()
	}
	{
		p.SetState(5)
		p.Match(NotesParserEOF)
	}

	return localctx
}

// INotesContext is an interface to support dynamic dispatch.
type INotesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotesContext differentiates from other interfaces.
	IsNotesContext()
}

type NotesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotesContext() *NotesContext {
	var p = new(NotesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NotesParserRULE_notes
	return p
}

func (*NotesContext) IsNotesContext() {}

func NewNotesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotesContext {
	var p = new(NotesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NotesParserRULE_notes

	return p
}

func (s *NotesContext) GetParser() antlr.Parser { return s.parser }

func (s *NotesContext) NOTES() antlr.TerminalNode {
	return s.GetToken(NotesParserNOTES, 0)
}

func (s *NotesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NotesListener); ok {
		listenerT.EnterNotes(s)
	}
}

func (s *NotesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NotesListener); ok {
		listenerT.ExitNotes(s)
	}
}

func (p *NotesParser) Notes() (localctx INotesContext) {
	this := p
	_ = this

	localctx = NewNotesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NotesParserRULE_notes)

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
		p.Match(NotesParserNOTES)
	}

	return localctx
}
