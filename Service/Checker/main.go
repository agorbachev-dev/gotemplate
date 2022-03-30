package Checker

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceID string
	Status    string
}
