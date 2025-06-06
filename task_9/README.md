### **Уровень 4: Интерфейсы и пакеты**

1. **Геометрические фигуры**
    
    Создайте интерфейс **`Shape`** с методом **`Area() float64`**. Реализуйте его для структур:
    
    - **`Circle`** (радиус).
    - **`Rectangle`** (ширина, высота).
    
    **Пример**:
    
    go
    
    ```
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}
    fmt.Println(circle.Area()) // 78.5398...
    fmt.Println(rectangle.Area()) // 24
    ```