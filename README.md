# COPYU Protocol 1.0

This is a simplifed gopher-type protocol using a TCP server.

So far this just fetches (1) a top level directory and (2) some small
text files (poetry by Emily Dickinson, from Project Gutenberg).

The environment is another simple system in which all requests are 64
non-zero bytes and all replies are 1024 non-zero bytes.  By non-zero we
mean there are no NUL '\0' characters in the So we've designed an easy
protocol that fits these requirements.

To keep the client trival, some work has been moved into the server,
or into a script that compiles the `demo/text` directory of text files
into the `demo/coypu` directory for serving.

A trivial client can talk to just one TCP server, and ignore the server IP
and port numbers in the directory.

# Request

The request is 64 bytes, all of which should be printable ASCII or spaces.

Bytes 0..1: a two-byte type.  "30" for directory, "31" for plaintext.

Bytes 2..9: Eight hex characters for the numeric value of the IP address
of the server.  NOT CURRENTLY USED.  This field should be ignored by
the server.

Bytes 10..13:  Four hex characters for the TCP port number of the server.
NOT CURRENTLY USED.  This field should be ignored by the server.

Bytes 14..31:  A "Selector" to choose a file on the server.  Pad with
spaces on the right.

Bytes 32..63:  A title to display for the file.  NOT CURRENTLY USED.
Pad with spaces on the right.  This field should be ignored by the server.

# Reply for Type 30: Directory

1024 bytes, which are 16 records.

Each records is a 64-byte Request, defined above.

Unused records are all ^Z characters (char 26) at the end.

# Reply for Type 31: Text File.

1024 bytes, which are ASCII text, using LF (char 10) for newlines.
The payload text must be shorter than 1024 bytes.  Use ^Z characters
(char 26) to fill the end.
