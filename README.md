## Difference between `uint` and `int`

The primitive data types prefixed with "u" are unsigned versions with the same bit sizes. Effectively, this means they cannot store negative numbers, but on the other hand they can store positive numbers twice as large as their signed counterparts. The signed counterparts do not have "u" prefixed.

The limits for int (32 bit) are:

```
int: –2147483648 to 2147483647
uint: 0 to 4294967295
```
And for long (64 bit):

```
long: -9223372036854775808 to 9223372036854775807
ulong: 0 to 18446744073709551615
```

Source: https://stackoverflow.com/questions/3724242/what-is-the-difference-between-int-and-uint-long-and-ulong