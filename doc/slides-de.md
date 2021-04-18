% Communicating Services
% Alexander Gunkel
% 17.04.2021
---

# Communicating Services
* Golang hat seine Stärken gerade im Bereich asynchroner Programmierung
* Die beiden wichtigsten Hilfsmittel sind:
    * Goroutines
    * Channels

----
    
# Goroutines als Services

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

* Auf diese Weise wird quasi ein eigenes Programm mit eigener main-loop in einer Goroutine erstellt.
* Event-basiert: Die Groutine kann verschiedene Ereignisse koordinieren, z.B.
  * Eingehende Nachrichten,
  * Shutdown-Signale
  * Update-Anfragen
* Vorteile durch quasi serialisierten Workflow
  * keine Synchronisierung nötig durch Mutex, Lock etc.
  * keine Data Races
  * keine Deadlocks

----

# Channels I: Grundlagen
````
ch := make(chan string, 1) // [1]
ch <- "Hello World" // [2]
v, open := <- ch // [3]
<- ch // [4]

close(ch) // [5]
v, open = <- ch [6]
````

* [1] erstellt einen Channel `ch` für Strings mit einer Bufferlänge von 1
* [2] sendet den String "Hello World" an den Channel `ch`
* [3] liest aus dem Channel `ch` und speichert "Hello World" in `v` sowie `true` in `open`.
* [4] wartet auf den nächsten Eintrag aus `ch`;
  in diesem Fall blockiert der Aufruf die Ausführung.
* [5] schließt den Channel `ch`. Weiteres Senden an `ch` 
  sowie ein weiterer Aufruf von `close(ch)` führen zu einem fatalen Fehler.
* [6] empfängt von `ch`. Da `ch` in [5] geschlossen wurde, blockiert dieser Aufruf nicht.
`v` enthält nun einen leeren String und `open` den Wert `false`. Dies signalisiert, dass der
  Channel geschlossen wurde.


----

# Channels II: Best Practice

```
ch := make(chan string)
var send chan<- string = ch
var receive <-chan string = ch

send <- "hello" // ok
receive <- "hello" // doesn't compile

<- receive // ok
<- send // doesn't compile

close(send) // ok
close(send) // fatal error
send <- "send to closed channel" // fatal error
close(receive) // doesn't compile
```
* Jeder Channel sollte genau einen Owner (eine `struct` oder eine Goroutine) haben,
  der als einziger senden (und schließen) darf.
* Beliebig vielen anderen wird der Channel als receive-only zur Verfügung gestellt.
* Der Owner signalisiert durch Schließen, dass alle Aufgaben abgearbeitet wurden.

----

# Service-Pipelines

```
type Node interface {
	Messages() <-chan Message
}

type StartNode interface {
	Node
	Stop()
}

type EndNode interface {
	Done() <-chan struct{}
}
```

* Alles, was ein einzelner Node unserer Service-Pipeline können muss, ist weitere Tasks auszugeben,
die dann von anderen Nodes ausgeführt werden.
* Ausschließlich die StartNode, die von keinen weiteren Nodes abhängen, brauchen eine Funktionalität
zum Anhalten der Ausführung.
* Ausschließlich die EndNodes benötigen eine Methode, die anderen anzeigt, ob sie noch Tasks erwarten.
Alle anderen signalisieren dies durch Schließen der Messages()-Channels.
* Die Services kennen voneinander ausschließlich die unidirektionales Messages-Channels.
* Die Anwendung schließt nur die Start-Nodes und wartet auf das Signal der End-Nodes.
  
----
# Beispiel

```

```
