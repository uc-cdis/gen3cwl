package cwlgotest

import (
	"reflect"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_tmap_tool(t *testing.T) {
	f := load("tmap-tool.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")

	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("python:2-slim")

	Expect(t, root.Inputs[0].ID).ToBe("reads")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].ID).ToBe("stages")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("array")
	Expect(t, root.Inputs[1].Types[0].Items[0].Type).ToBe("#Stage")
	Expect(t, root.Inputs[1].Binding.Position).ToBe(1)
	Expect(t, root.Inputs[2].ID).ToBe("#args.py")
	Expect(t, root.Inputs[2].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[2].Default.Kind).ToBe(reflect.Map)
	Expect(t, root.Inputs[2].Binding.Position).ToBe(-1)

	Expect(t, root.Outputs[0].ID).ToBe("sam")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"output.sam"})
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("null")
	Expect(t, root.Outputs[0].Types[1].Type).ToBe("File")
	Expect(t, root.Outputs[1].ID).ToBe("args")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("string[]")

	Expect(t, root.Requirements[0].Class).ToBe("SchemaDefRequirement")
	// Expect(t, root.Requirements[0].Types[0].Name).ToBe("Map1")
	Expect(t, root.Requirements[0].Types[0].Type).ToBe("record")
	Expect(t, root.Requirements[0].Types[0].Fields[0].Name).ToBe("algo")
	Expect(t, root.Requirements[0].Types[0].Fields[0].Types[0].Type).ToBe("enum")
	// Expect(t, root.Requirements[0].Types[0].Fields[0].Types[0].Name).ToBe("JustMap1")
	Expect(t, root.Requirements[0].Types[0].Fields[0].Types[0].Symbols[0]).ToBe("map1")
	Expect(t, root.Requirements[0].Types[0].Fields[0].Binding.Position).ToBe(0)
	Expect(t, root.Requirements[0].Types[0].Fields[1].Name).ToBe("maxSeqLen")
	Expect(t, root.Requirements[0].Types[0].Fields[1].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[0].Fields[1].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[0].Fields[1].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[0].Fields[1].Binding.Prefix).ToBe("--max-seq-length")
	Expect(t, root.Requirements[0].Types[0].Fields[2].Name).ToBe("minSeqLen")
	Expect(t, root.Requirements[0].Types[0].Fields[2].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[0].Fields[2].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[0].Fields[2].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[0].Fields[2].Binding.Prefix).ToBe("--min-seq-length")
	Expect(t, root.Requirements[0].Types[0].Fields[3].Name).ToBe("seedLength")
	Expect(t, root.Requirements[0].Types[0].Fields[3].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[0].Fields[3].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[0].Fields[3].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[0].Fields[3].Binding.Prefix).ToBe("--seed-length")

	// Expect(t, root.Requirements[0].Types[1].Name).ToBe("Map2")
	Expect(t, root.Requirements[0].Types[1].Type).ToBe("record")
	Expect(t, root.Requirements[0].Types[1].Fields[0].Name).ToBe("algo")
	Expect(t, root.Requirements[0].Types[1].Fields[0].Types[0].Type).ToBe("enum")
	// Expect(t, root.Requirements[0].Types[1].Fields[0].Types[0].Name).ToBe("JustMap2")
	Expect(t, root.Requirements[0].Types[1].Fields[0].Types[0].Symbols[0]).ToBe("map2")
	Expect(t, root.Requirements[0].Types[1].Fields[0].Binding.Position).ToBe(0)
	Expect(t, root.Requirements[0].Types[1].Fields[1].Name).ToBe("maxSeqLen")
	Expect(t, root.Requirements[0].Types[1].Fields[1].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[1].Fields[1].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[1].Fields[1].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[1].Fields[1].Binding.Prefix).ToBe("--max-seq-length")
	Expect(t, root.Requirements[0].Types[1].Fields[2].Name).ToBe("minSeqLen")
	Expect(t, root.Requirements[0].Types[1].Fields[2].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[1].Fields[2].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[1].Fields[2].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[1].Fields[2].Binding.Prefix).ToBe("--min-seq-length")
	Expect(t, root.Requirements[0].Types[1].Fields[3].Name).ToBe("maxSeedHits")
	Expect(t, root.Requirements[0].Types[1].Fields[3].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[1].Fields[3].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[1].Fields[3].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[1].Fields[3].Binding.Prefix).ToBe("--max-seed-hits")

	// Expect(t, root.Requirements[0].Types[2].Name).ToBe("Map3")
	Expect(t, root.Requirements[0].Types[2].Type).ToBe("record")
	Expect(t, root.Requirements[0].Types[2].Fields[0].Name).ToBe("algo")
	Expect(t, root.Requirements[0].Types[2].Fields[0].Types[0].Type).ToBe("enum")
	// Expect(t, root.Requirements[0].Types[2].Fields[0].Types[0].Name).ToBe("JustMap3")
	Expect(t, root.Requirements[0].Types[2].Fields[0].Types[0].Symbols[0]).ToBe("map3")
	Expect(t, root.Requirements[0].Types[2].Fields[0].Binding.Position).ToBe(0)
	Expect(t, root.Requirements[0].Types[2].Fields[1].Name).ToBe("maxSeqLen")
	Expect(t, root.Requirements[0].Types[2].Fields[1].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[2].Fields[1].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[2].Fields[1].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[2].Fields[1].Binding.Prefix).ToBe("--max-seq-length")
	Expect(t, root.Requirements[0].Types[2].Fields[2].Name).ToBe("minSeqLen")
	Expect(t, root.Requirements[0].Types[2].Fields[2].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[2].Fields[2].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[2].Fields[2].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[2].Fields[2].Binding.Prefix).ToBe("--min-seq-length")
	Expect(t, root.Requirements[0].Types[2].Fields[3].Name).ToBe("fwdSearch")
	Expect(t, root.Requirements[0].Types[2].Fields[3].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[2].Fields[3].Types[1].Type).ToBe("boolean")
	Expect(t, root.Requirements[0].Types[2].Fields[3].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[2].Fields[3].Binding.Prefix).ToBe("--fwd-search")

	// Expect(t, root.Requirements[0].Types[3].Name).ToBe("Map4")
	Expect(t, root.Requirements[0].Types[3].Type).ToBe("record")
	Expect(t, root.Requirements[0].Types[3].Fields[0].Name).ToBe("algo")
	Expect(t, root.Requirements[0].Types[3].Fields[0].Types[0].Type).ToBe("enum")
	// Expect(t, root.Requirements[0].Types[3].Fields[0].Types[0].Name).ToBe("JustMap4")
	Expect(t, root.Requirements[0].Types[3].Fields[0].Types[0].Symbols[0]).ToBe("map4")
	Expect(t, root.Requirements[0].Types[3].Fields[0].Binding.Position).ToBe(0)
	Expect(t, root.Requirements[0].Types[3].Fields[1].Name).ToBe("maxSeqLen")
	Expect(t, root.Requirements[0].Types[3].Fields[1].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[3].Fields[1].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[3].Fields[1].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[3].Fields[1].Binding.Prefix).ToBe("--max-seq-length")
	Expect(t, root.Requirements[0].Types[3].Fields[2].Name).ToBe("minSeqLen")
	Expect(t, root.Requirements[0].Types[3].Fields[2].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[3].Fields[2].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[3].Fields[2].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[3].Fields[2].Binding.Prefix).ToBe("--min-seq-length")
	Expect(t, root.Requirements[0].Types[3].Fields[3].Name).ToBe("seedStep")
	Expect(t, root.Requirements[0].Types[3].Fields[3].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[3].Fields[3].Types[1].Type).ToBe("int")
	Expect(t, root.Requirements[0].Types[3].Fields[3].Binding.Position).ToBe(2)
	Expect(t, root.Requirements[0].Types[3].Fields[3].Binding.Prefix).ToBe("--seed-step")

	// Expect(t, root.Requirements[0].Types[4].Name).ToBe("Stage")
	Expect(t, root.Requirements[0].Types[4].Type).ToBe("record")
	Expect(t, root.Requirements[0].Types[4].Fields[0].Name).ToBe("stageId")
	Expect(t, root.Requirements[0].Types[4].Fields[0].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[4].Fields[0].Binding.Position).ToBe(0)
	Expect(t, root.Requirements[0].Types[4].Fields[0].Binding.Prefix).ToBe("stage")
	Expect(t, root.Requirements[0].Types[4].Fields[0].Binding.Separate).ToBe(false)
	Expect(t, root.Requirements[0].Types[4].Fields[1].Name).ToBe("stageOption1")
	Expect(t, root.Requirements[0].Types[4].Fields[1].Types[0].Type).ToBe("null")
	Expect(t, root.Requirements[0].Types[4].Fields[1].Types[1].Type).ToBe("boolean")
	Expect(t, root.Requirements[0].Types[4].Fields[1].Binding.Position).ToBe(1)
	Expect(t, root.Requirements[0].Types[4].Fields[1].Binding.Prefix).ToBe("-n")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Name).ToBe("algos")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Types[0].Type).ToBe("array")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Types[0].Items[0].Type).ToBe("#Map1")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Types[0].Items[1].Type).ToBe("#Map2")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Types[0].Items[2].Type).ToBe("#Map3")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Types[0].Items[3].Type).ToBe("#Map4")
	Expect(t, root.Requirements[0].Types[4].Fields[2].Binding.Position).ToBe(2)
}
