//go:generate go run ./ynab_enum_gen Account Checking:checking Savings:savings Cash:cash CreditCard:creditCard LineOfCredit:lineOfCredit OtherAsset:otherAsset OtherLiability:otherLiability Mortgage:mortgage AutoLoan:autoLoan StudentLoan:studentLoan PersonalLoan:personalLoan MedicalDebt:medicalDebt OtherDebt:otherDebt
//go:generate go run ./ynab_enum_gen Goal TB:TB TBD:TBD MF:MF NEED:NEED DEBT:DEBT
//go:generate go run ./ynab_enum_gen Cleared Cleared:cleared Uncleared:uncleared Reconciled:reconciled
//go:generate go run ./ynab_enum_gen Flag Red:red Orange:orange Yellow:yellow Green:green Blue:blue Purple:purple

package ynab
