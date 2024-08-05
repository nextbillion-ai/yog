package yog

// MatrixData pod matrix response
type MatrixData struct {
	Rows []MatrixRow `json:"rows"`
}

type MatrixRow struct {
	Elements []MatrixElement `json:"elements"`
}

type MatrixElement struct {
	Duration Value `json:"duration"`
	Distance Value `json:"distance"`
}

type Value struct {
	Value uint32 `json:"value"`
}

/*
*	file format:
*
*	{
*			"index": {
*					"1": {
*							"origin": [1, 2, 3, 4],
*							"destination": [1, 2, 3, 4]
*					},
*					"{chunk_index}": {
*							"origin {pg origin_index}": [1, 2, 3, 4],
*							"destination {pg destination_index}": [1, 2, 3, 4]
*					}
*			}
*	}
*
**/

type FileState int

const (
	FSucceeded FileState = iota
	FFailed
)

// TaskMeta SDK 需要的一些 meta 信息
type TaskMeta struct {
	Index      map[string]IndexItem `json:"index"`
	MatrixInfo MatrixInfo           `json:"matrix_info,omitempty"`
}

type IndexItem struct {
	Status      FileState `json:"status,omitempty"`
	Origin      []int     `json:"origin"`
	Destination []int     `json:"destination"`
	Offset      []int     `json:"offset,omitempty"`
}

type MatrixInfo struct {
	OriginCount      int `json:"origin_count,omitempty"`
	DestinationCount int `json:"destination_count,omitempty"`
}
