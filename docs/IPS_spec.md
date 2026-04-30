# IPS File Format Specification

_By Z.e.r.o / ZeroSoft. Source: https://zerosoft.zophar.net/ips.php_

---

## Limits

- IPS files can patch any file not larger than 2^24-1 bytes (~16 MB)
- Each individual patch record must not be larger than 2^16-1 bytes (~64 KB)
- An IPS file can hold as many records as needed, as long as each offset and size stays within the above bounds

---

## File Structure

| Section         | Size (bytes)    | Description                                                    |
|-----------------|-----------------|----------------------------------------------------------------|
| Header          | 5               | The ASCII string `PATCH` (not NULL-terminated)                 |
| Record          | 3 + 2 + variable| A single patch record (see below). Repeated as many times as needed. |
| EOF marker      | 3               | The ASCII string `EOF` (not NULL-terminated)                   |

---

## Record Structure

| Section | Size (bytes) | Description                                                    |
|---------|--------------|----------------------------------------------------------------|
| Offset  | 3            | The offset in the target file where the patch data will be written |
| Size    | 2            | The number of bytes to write at the given offset               |
| Data    | `Size`       | The raw bytes to copy into the target file at `Offset`         |

> If `Size == 0`, this is an RLE-encoded record — see below.

---

## RLE Record Structure

When the `Size` field of a record reads `0`, the record is RLE (Run-Length Encoded). Instead of raw data, a single byte value is repeated N times.

| Section  | Size (bytes) | Description                                             |
|----------|--------------|---------------------------------------------------------|
| Offset   | 3            | The offset in the target file where writing begins      |
| Size     | 2            | Always `0` — signals RLE encoding                      |
| RLE_Size | 2            | The number of times to write `Value` (must be nonzero)  |
| Value    | 1            | The byte to repeat `RLE_Size` times starting at `Offset`|

---

## Endianness

Offset and Size values are stored in **big-endian** order (most significant byte first). This is sometimes called "Pascal/Basic" style. C and C++ on x86 machines are little-endian, so byte-swapping is required when reading these values.

A 16-bit value `0x6712` is stored as: `67 12`  
A 24-bit value `0x671234` is stored as: `67 12 34`

### C macro equivalents (for reference)

```c
#define BYTE3_TO_UINT(bp) \
    (((unsigned int)(bp)[0] << 16) & 0x00FF0000) | \
    (((unsigned int)(bp)[1] <<  8) & 0x0000FF00) | \
    ((unsigned int) (bp)[2]        & 0x000000FF)

#define BYTE2_TO_UINT(bp) \
    (((unsigned int)(bp)[0] << 8) & 0xFF00) | \
    ((unsigned int) (bp)[1]       & 0x00FF)
```

In Go, the equivalent reads are:

```go
offset := int(buf[0])<<16 | int(buf[1])<<8 | int(buf[2])
size   := int(buf[0])<<8  | int(buf[1])
```
