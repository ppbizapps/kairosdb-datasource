package kairos

type Request struct {
	StartAbsolute string         `json:"start_absolute"`
	EndAbsolute   string         `json:"end_absolute"`
	Metrics       []*MetricQuery `json:"metrics"`
}

type MetricQuery struct {
	Name        string        `json:"name"`
	Aggregators []*Aggregator `json:"aggregators,omitempty"`
	GroupBy     []*Grouper    `json:"group_by,omitempty"`
}

type Aggregator struct {
	Name           string    `json:"name"`
	AlignSampling  bool      `json:"align_sampling"`
	AlignStartTime bool      `json:"align_start_time"`
	AlignEndTime   bool      `json:"align_end_time"`
	Sampling       *Sampling `json:"sampling"`
}

type Sampling struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Response struct {
	Queries []*QueryResponse `json:"queries"`
}

type QueryResponse struct {
	Results []*QueryResult `json:"results"`
}

type QueryResult struct {
	Name string `json:"name"`
	//Tags    map[string]string `json:"tags,omitempty"`
	Values []*DataPoint `json:"values"`
}

//TODO support group by time and value
type Grouper struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type DataPoint [2]float64
