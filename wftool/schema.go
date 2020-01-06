package wftool

import ()

// right now pretty much just writing out the CWL spec in Go types

// will this tool just marshal without enforcing/validating the cwl?
// e.g., if scatter, then scattermethod - will we perform that check here?
// or does this tool assume your cwl is error-free
// probably this tool should have some kind of validation function
// this tool should answer, to some degree, the question - "will this cwl run?"
// "will mariner even attempt to run this workflow?"

// WorkflowJSON is the JSON representation of a CWL workflow
type WorkflowJSON struct {
	Graph WorkflowGraph
}

// WorkflowGraph contains all the CWLObjects of the workflow
type WorkflowGraph []CWLObject

// as much as possible, don't repeat fields among structs
// make basic, atomic structs, and embed into other structs as needed

// CWLObject represents a workflow, expressiontool, commandlinetool, ...
// TODO
type CWLObject interface {
	// some methods
}

// ObjectMeta ..
type ObjectMeta struct {
	CoreMeta
	RequirementsAndHints
	Class      string
	CWLVersion string
}

// RequirementsAndHints ..
// NOTE: possibly need to make types for all the different requirements
// though I'm not sure that'll be necessary at all
// since the requirements themselves don't get touched/modifed during unmarshalling
type RequirementsAndHints struct {
	Requirements []interface{}
	Hints        []interface{} // some schema!
}

// CoreMeta ..
type CoreMeta struct {
	ID    string
	Label string
	Doc   string
}

// Workflow ..
type Workflow struct {
	ObjectMeta
	Inputs  []InputParameter
	Outputs []WorkflowOutputParameter
	Steps   []WorkflowStep
}

// WorkflowStep ..
type WorkflowStep struct {
	CoreMeta
	RequirementsAndHints
	In            []WorkflowStepInput
	Out           []WorkflowStepOutput // could be []string or []map[string]string
	Run           interface{}
	Scatter       []string
	ScatterMethod string
}

// WorkflowStepInput ..
type WorkflowStepInput struct {
	ID        string
	Source    []string
	LinkMerge string // 'merge_nested' or 'merge_flattened'
	Default   interface{}
	ValueFrom string
}

// WorkflowStepOutput ..
type WorkflowStepOutput struct {
	ID string
}

// InputParameter ..
type InputParameter struct {
	CoreMeta
	FileParameterFields
	InputBinding CommandLineBinding
	Default      interface{}
}

// CommandLineBinding ..
type CommandLineBinding struct {
	LoadContents  bool
	Position      int
	Prefix        string
	Separate      bool
	ItemSeparator string
	ValueFrom     string
	ShellQuote    bool
}

// WorkflowOutputParameter ..
type WorkflowOutputParameter struct {
	CoreMeta
	FileParameterFields
	OutputBinding CommandOutputBinding
	OutputSource  []string
	LinkMerge     string
}

// FileParameterFields ..
type FileParameterFields struct {
	SecondaryFiles []string
	Streamable     bool
	Format         []string
	Type           []interface{} // TODO
	// NOTE: handling 'Type' requires some thought - several possibilities here
	// see: https://www.commonwl.org/v1.0/Workflow.html#InputParameter
}

// CommandOutputBinding ..
type CommandOutputBinding struct {
	Glob         []string
	LoadContents bool
	OutputEval   string
}

// CommandLineTool ..
type CommandLineTool struct {
	ObjectMeta
	Inputs             []CommandInputParameter
	Outputs            []CommandOutputParameter
	BaseCommand        []string
	Arguments          []interface{} // an argument is one of 'expression' | 'string' | 'commandlinebinding'
	Stdin              string
	Stderr             string
	Stdout             string
	SuccessCodes       []int
	TemporaryFailCodes []int
	PermanentFailCodes []int
}

// CommandInputParameter ..
type CommandInputParameter struct {
	CoreMeta
	FileParameterFields
	InputBinding CommandLineBinding
	Default      interface{}
}

// CommandOutputParameter ..
type CommandOutputParameter struct {
	CoreMeta
	FileParameterFields
	OutputBinding CommandOutputBinding
}

// ExpressionTool ..
type ExpressionTool struct {
	ObjectMeta
	Inputs     []InputParameter
	Outputs    []ExpressionToolOutputParameter
	Expression string
}

// ExpressionToolOutputParameter ..
type ExpressionToolOutputParameter struct {
	CoreMeta
	FileParameterFields
	OutputBinding CommandOutputBinding
}
