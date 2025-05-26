1. **Калькулятор валют**
    
    Создайте структуру **`CurrencyConverter`** с полями:
    
    - **`Rate`** (курс доллара к рублю).
        
        Методы:
        
    - **`ConvertToUSD(rubles float64) float64`** — конвертирует рубли в доллары.
    - **`ConvertToRUB(dollars float64) float64`** — конвертирует доллары в рубли.
    
    **Пример использования**:
    
    go
    
    ```
    converter := CurrencyConverter{Rate: 75.5}
    fmt.Println(converter.ConvertToUSD(755)) // 10.0
    ```
    
