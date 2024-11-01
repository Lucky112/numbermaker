## Numbermaker

Игра по потивам Beltmatic и базовой арифметики. Основной геймплей заключается в постороении арифметических выражений с помощью источников чисел и операций.

### Архитектура
#### Visual
Используется ebiten.
UX
- Клеточное поле
- Источники, конверторы, приемники, пайпы, свичи
- Истоники имеют выходные порты
- Приемники имеют входные порты
- Драг от входного порта до выходного создает пайп поверх затронутых клеток поля
  - Может не окончится на входном порту, тогда создается свитч с одним входом в точке конца
  - Только завершение драга создает пайп в Core, но рисуется и в процессе тоже
- Структуры
  - Game
    - Следит за состоянием игры, помнит что мышь нажата
  - Board
    - Само игровое поле
    - Содержит ячейки, объекты и пайпы, добавляет и рисует их
  - Cell
    - ячейка поля
  - Producer
  - Consumer
  - Convertor
  - Switch
  - Pipe
    - набор cells
  - Number

Проблемы:
- Борд это матрица ячеек и удобно в ней навигивировать в координатах номеров ячеек, но рисовать то надо в координатах экрана, причем все сущности
- Построением нового пайпа занимается Game и делает это не слишком хорошо. Должно быть нужен отдельный класс, с памятью направления.
- неэффективный поиск пересечений. Можно было бы сделать хештаблицу по координатам ячеек и так проверять
- непонятно как в эту схему с ячейками вписать порты

h5. Борд и ячейки

Board не содержит cells

ScreenPoint {x,y}
BoardPoint {r, c}

Node {
  Origin BoardPoint
  im Image
}

Node.Draw(scale)

Board{
  Nodes
  Pipes
  Ports ???

  AddNode(node)
  AddPipe(pipe)
  AddPort(port)
}

PipeBuilder {
  direction
  current_pipe
}

Game {
  VisualField {
    background []Cell
    Board
    PipeBuilder
  }
}

#### Core
Ориентированный граф связей между узлами
- Структуры
  - Producer
    - Выставляет на порты
  - Consumer
    - Читает с портов
  - Convertor
    - Consume
    - Process
    - Produce
  - Switch
    - Распределяет входы на выходы по определенному правилу
  - Pipe
    - стек чисел, возможны nil, и проталкивание на место nil
  - Number