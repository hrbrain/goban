package series

type AggregationMethod string

const (
	Mean  = AggregationMethod("Mean")
	Count = AggregationMethod("Count")
	Sum   = AggregationMethod("Sum")
	None  = AggregationMethod("")
)
