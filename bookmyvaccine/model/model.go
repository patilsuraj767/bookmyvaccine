package model

type Session struct {
	Date                     string
	Available_capacity       int
	Min_age_limit            int
	Vaccine                  string
	Slots                    []string
	Available_capacity_dose1 int
	Available_capacity_dose2 int
}

type Center struct {
	Center_id int
	Name      string
	Address   string
	Sessions  []Session
}

type Data struct {
	Centers []Center
}
