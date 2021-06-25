# tcp

## Remark
> the simplest way to implement a tls with out authentication
Generating a certificate directly from a private key(no matter it is a server side key or ca key),  which means it just use that private key to sign an input certificate, whether or not
perform an authentication of that signature is decided by a client who actually communicate with the server with handshake protocol.

A client who decide to establish tls without authenticate still have the untility of an encrypted transport layer.