package r2g2

import "encoding/json"

const (
	getFunctionDetailCMD = "pdfj"
)

type Function struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`
	Addr uint64 `json:"addr"`
	Ops  []*Op  `json:"ops"`
}

type Op struct {
	Offset   uint64   `json:"offset"`
	Ptr      uint64   `json:"ptr"`
	ESIL     string   `json:"esil"`
	RefPtr   bool     `json:"refptr"`
	FcnAddr  uint64   `json:"fcn_addr"`
	FcnLast  uint64   `json:"fcn_last"`
	Size     uint64   `json:"size"`
	OpCode   string   `json:"opcode"`
	Disasm   string   `json:"disasm"`
	Bytes    string   `json:"bytes"`
	Family   string   `json:"family"`
	Type     string   `json:"type"`
	TypeNum  uint64   `json:"type_num"`
	Type2Num uint64   `json:"type2_num"`
	Flags    []string `json:"flags"`
	Jump     uint64   `json:"jump"`
	Fail     uint64   `json:"fail"`
	XRefs    []*XRef  `json:"xrefs"`
}

type XRef struct {
	Addr uint64 `json:"addr"`
	Type string `jsoin:"type"`
}

func (c *Client) GetCurrentFunction() (*Function, error) {
	data, err := c.Run(getFunctionDetailCMD)
	if err != nil {
		return nil, err
	}
	var function Function
	err = json.Unmarshal(data, &function)
	return &function, err
}
