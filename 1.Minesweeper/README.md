# Task
**Input**: - A 2D list (matrix) representing a Minesweeper board - Each cell must be either 0
(empty space) or 1 (mine)
**Output**: - A 2D list (matrix) of the same dimensions as the input - Each cell contains the count
of neighbouring mines (0-8), or 9 if the cell contains a mine



Code:

```
def count_neighbouring_mines(matrix: list) -> list:
    
    """
    Counts neighbouring mines for each cell in a Minesweeper board.
    Parameters:
    board (list): A 2D list where 0 represents an empty space and 1
    represents a mine
    Returns:
        list: A 2D list where each cell contains the count of neighbouring mines,
    or 9 if the cell contains a mine
    """
    # Your implementation here
    
    # obtiene el número de filas y columnas de la matriz
    rows = len(matrix)
    cols = len(matrix[0]) if rows > 0 else 0
    # se crea una matriz de resultados con el mismo tamaño que la matriz de entrada, inicializada con ceros
    result = [[0 for i in range(cols)] for j in range(rows)]
    # se itera sobre cada celda de la matriz de entrada
    # para cada fila i y cada columna j, se verifica si la celda contiene una mina (valor 1)
    for i in range(rows):
        for j in range(cols):
            if matrix[i][j] == 1:
                result[i][j] = 9
            else:
                # conteo de minas cuando la celda no contiene una mina
                count = 0
                # se itera sobre las celdas vecinas (incluyendo la celda actual) 
                # utilizando un rango que va desde i-1 hasta i+1 y j-1 hasta j+1, 
                # asegurándose de no salir de los límites de la matriz
                for x in range(max(0, i - 1), min(rows, i + 2)):
                    for y in range(max(0, j - 1), min(cols, j + 2)):
                        if matrix[x][y] == 1:
                            count += 1
                result[i][j] = count
                
    return result
```
Additionally, the solution to this task is in the file called: **solution_minesweeper-py**
