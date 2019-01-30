package ynab

type CurrencyFormat struct {
    ISOCode          string `json:"iso_code"`
    ExampleFormat    string `json:"example_format"`
    DecimalDigits    uint64 `json:"decimal_digits"`
    DecimalSeparator string `json:"decimal_separator"`
    GroupSeparator   string `json:"group_separator"`
    SymbolFirst      bool   `json:"symbol_first"`
    CurrencySymbol   string `json:"currency_symbol"`
    DisplaySymbol    bool   `json:"display_symbol"`
}
