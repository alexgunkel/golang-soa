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
            // service
        }
    }()
~~~

* können zur Service-orientierten Programmierung genutzt werden:
* Auf diese Weise wird quasi ein eigenes Programm mit eigener main-loop
in einer Goroutine erstellt.


---
# Channels