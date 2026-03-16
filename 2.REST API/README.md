# Task
Given a genre, find the series with the highest imdb_rating . If there is a tie, return the alphabetically lower name.
# Sample
**Input**: Action 
**Output**: Game of Thrones


# Solution
important fields:
* total_pages
* page
* data: it contains data for the series in the current page.
  * genre
  * imdb_rating
  * name

code
```
mport urllib.request
import json

def bestInGenre(genre: str) -> str:
    """
    Finds the highest-rated TV series in the given genre.
    Parameters:
    genre (str): The genre to search for (e.g., 'Action', 'Comedy','Drama')
    Returns:
        str: The name of the highest-rated show in the genre. 
        If there is a tie, returns the alphabetically lower name. Returns the name as a string.
    Notes:
        - Ties are broken by alphabetical order of the show name
        - Genre matching is case-insensitive
        - Shows can have multiple genres (comma-separated)
    """
    # Your implementation here
    # Estandarizamos el género buscado
    genero_buscado = genre.lower().strip()
    
    mejor_serie = ""
    puntaje_mas_alto = -1.0
    
    url_base = "https://jsonmock.hackerrank.com/api/tvseries"
    
    # Paso 1: Hacer una primera consulta para descubrir cuántas páginas existen en total
    try:
        with urllib.request.urlopen(f"{url_base}?page=1") as respuesta:
            datos_iniciales = json.loads(respuesta.read().decode('utf-8'))
            total_paginas = datos_iniciales.get('total_pages', 1)
    except Exception:
        # Si la API falla o no hay conexión, devolvemos una cadena vacía
        return ""
    # Paso 2: Iterar a través de todas las páginas para encontrar la mejor serie en el género
    for pagina in range(1, total_paginas + 1):
        try:
            with urllib.request.urlopen(f"{url_base}?page={pagina}") as respuesta:
                datos = json.loads(respuesta.read().decode('utf-8'))
                series = datos.get('data', [])
                
                for serie in series:
                    # Extraemos el nombre, género y puntaje de la serie
                    nombre = serie.get('name', '').strip()
                    generos = serie.get('genre', '').lower().split(',')
                    puntaje = float(serie.get('imdb_rating', 0))
                    
                    # Verificamos si el género buscado está presente en los géneros de la serie
                    if genero_buscado in [g.strip() for g in generos]:
                        # Si el puntaje es mayor que el puntaje más alto encontrado hasta ahora
                        if puntaje > puntaje_mas_alto:
                            mejor_serie = nombre
                            puntaje_mas_alto = puntaje
                        # Si hay un empate en el puntaje, se desempata por orden alfabético
                        elif puntaje == puntaje_mas_alto and nombre < mejor_serie:
                            mejor_serie = nombre
        except Exception:
            # Si la API falla o no hay conexión, continuamos con la siguiente página
            continue
    return mejor_serie
```

Also, the solution is in the file: **solution_best_in_genre.py**
