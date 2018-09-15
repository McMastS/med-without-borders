package main

type MedicationType string

const (
	Aspirin        MedicationType = "C0001"
	Amiloride      MedicationType = "C0002"
	Amiodarone     MedicationType = "C0003"
	Bisoprolol     MedicationType = "C0004"
	Clopidogrel    MedicationType = "C0005"
	Digoxin        MedicationType = "C0006"
	Furosemide     MedicationType = "C0007"
	Losartan       MedicationType = "C0008"
	Methyldopa     MedicationType = "C0009"
	Nifedipine     MedicationType = "C0010"
	Spironolactone MedicationType = "C0011"
	Streptokinase  MedicationType = "C0012"
	Verapamil      MedicationType = "C0013"
)

// func MedicationTypeFromString(s string) MedicationType {
// 	switch s {
// 	case "C0001":
// 		return Aspirin
// 	case "C0002":
// 		return Amiloride
// 	case "C0003":
// 		return Amiodarone
// 	case "C0004":
// 		return Bisoprolol
// 	}
// }
