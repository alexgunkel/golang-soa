% Communicating Services
% Alexander Gunkel
% 17.04.2021
-----

# Communicating Services
* Golang hat seine Stärken gerade im Bereich asynchroner Programmierung
* Die beiden wichtigsten Hilfsmittel sind:
    * Goroutines
    * Channels

---
    
# Goroutines

~~~
    go func() {
        for true {
          select {
            case <- event1:
              // handle event1
            case <- event2:
              // handle event2
          }
        }
    }()
~~~

* können zur Service-orientierten Programmierung genutzt werden:
* Auf diese Weise wird quasi ein eigenes Programm mit eigener main-loop in einer Goroutine erstellt.
* Event-basiert: Die Groutine kann verschiedene Ereignisse koordinieren, z.B.
  * Eingehende Nachrichten,
  * Shutdown-Signale
  * Update-Anfragen


---
# Channels
~~~
v, open := <- ch
~~~
* liest aus Channel `ch`
* speichert den gelesenen Wert in `v`
* speichert in `open` (boolean), ob der Channel noch offen ist, also nicht mit `close(ch)` geschlossen wurde.