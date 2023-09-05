package tranco

import "time"

type ApiDate struct {
	time.Time
}

func (t *ApiDate) UnmarshalJSON(data []byte) error {
	d, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	t.Time = d
	return err
}

type ApiTimestamp struct {
	time.Time
}

func (t *ApiTimestamp) UnmarshalJSON(data []byte) error {
	d, err := time.Parse(`"2006-01-02T15:04:05.000000"`, string(data))
	if err != nil {
		return err
	}
	t.Time = d
	return err
}
