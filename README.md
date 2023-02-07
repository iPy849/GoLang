# Algunos consejos para cuando vuelvas a ver esto:
En Linux existe un comando que se llama "nc" y que es para trabajar con sockets (TCP/IP, UDP, UNIX Domain Sockets). Para levantar un servidor sencillo esta herramienta funciona más que bien, usa el comando:
```bash
nc -lp <Port>
```
este comando te va a permitir levantar un server tranquilamente donde la información se puede enviar se forma interactiva.

# Orden de lectura
1. [tcpC.go](tcpC.go)
2. [tcpS.go](tcpS.go)
3. [udpC.go](udpC.go)
4. [udpS.go](udpS.go)
5. [socketClient.go](socketClient.go)
6. [socketServer.go](socketServer.go)
7. [otherTCPclient.go](otherTCPclient.go)
8. [otherTCPserver.go](otherTCPserver.go)
9. [concTCP.go](concTCP.go)
10. [ws.go](ws.go)

# Mis propios scripts
- [m_tcpC.go](m_tcpC.go) -> Arregla la salida y separa la lectura y escritura en dos gorrutinas
- [m_tcpS.go](m_tcpS.go) -> Atiende las conexiones en gorrutinas y responde a cada llamada que se hace

# Tareas del capítulo (la verdad no entendí esta vez)
- 
