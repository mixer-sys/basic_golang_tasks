### **Уровень 3: Обработка ошибок и тестирование**

1. **Валидатор пароля**
    
    Напишите функцию **`ValidatePassword(pass string) error`**, которая проверяет:
    
    - Длина >= 8 символов.
    - Есть минимум одна цифра и заглавная буква.
        
        Возвращает **`nil`**, если пароль валиден, или ошибку с описанием проблемы.
        
    
    **Пример ошибки**:
    
    go
    
    ```
    err := ValidatePassword("qwerty")
    fmt.Println(err) // "password must be at least 8 characters long"
    ```