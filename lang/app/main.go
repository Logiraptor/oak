package main

type JobApplication struct {
	CompanyName string
	Salary      int
	Accepted    bool
}

func main() {
	fields := structToFields(JobApplication{})

	printCreateTable(fields)
}
