package dataframe

// Records are converted object from dataframe by rows
type Records []Record

func NewRecords(records []Record) Records {
	return records
}

func (r Records) Append(record Record) Records {
	return append(r, record)
}
