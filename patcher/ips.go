package patcher

import (
	"fmt"
	"io"
	"os"
)

// IPS represents the IPS patching format
type IPS struct{}

// Apply creates a patched copy of baseFile by applying the IPS patch in patchFile,
// and writes the result to outFile. The original baseFile is never modified.
// Returns an error if any file cannot be opened, if the patch header is invalid,
// or if reading or writing fails at any point during patching.
func (IPS) Apply(baseFile, patchFile, outFile string) error {

	patchReader, err := os.Open(patchFile)

	if err != nil {
		return fmt.Errorf("failed to open %s: %w", patchFile, err)
	}

	defer patchReader.Close()

	buf := make([]byte, 5)

	_, err = io.ReadFull(patchReader, buf)

	if err != nil {
		return fmt.Errorf("failed to read %s: %w", patchReader.Name(), err)
	}

	if string(buf) != "PATCH" {
		return fmt.Errorf("invalid IPS file: missing PATCH header")
	}

	base, err := os.Open(baseFile)

	if err != nil {
		return fmt.Errorf("failed to open %s: %w", baseFile, err)
	}

	defer base.Close()

	out, err := os.Create(outFile)

	if err != nil {
		return fmt.Errorf("failed to create %s: %w", outFile, err)
	}

	defer out.Close()

	// Copy the entire base ROM into the output file before applying patches
	_, err = io.Copy(out, base)

	if err != nil {
		return fmt.Errorf("failed to copy %s to %s: %w", base.Name(), out.Name(), err)
	}

	// Read and apply patch records one by one until the EOF marker is reached.
	// Each record contains a 3-byte offset (where to write in the output file),
	// a 2-byte size (how many bytes to write), and the data bytes themselves.
	// If size is 0, it is an RLE record: instead of raw data, a single byte
	// value is repeated N times — a compact way to encode large runs of the
	// same byte without storing them all individually in the patch file.
	for {
		offsetBuf := make([]byte, 3)

		_, err = io.ReadFull(patchReader, offsetBuf)

		if err != nil {
			return fmt.Errorf("failed to read offset: %w", err)
		}

		if string(offsetBuf) == "EOF" {
			break
		}

		offset := int(offsetBuf[0])<<16 | int(offsetBuf[1])<<8 | int(offsetBuf[2])

		sizeBuf := make([]byte, 2)
		_, err = io.ReadFull(patchReader, sizeBuf)
		if err != nil {
			return fmt.Errorf("failed to read size: %w", err)
		}
		size := int(sizeBuf[0])<<8 | int(sizeBuf[1])

		if size == 0 {
			// RLE record: size=0 signals that the next 2 bytes are the repeat
			// count, and the byte after that is the value to repeat
			rleBuf := make([]byte, 3)
			_, err = io.ReadFull(patchReader, rleBuf)
			if err != nil {
				return fmt.Errorf("failed to read RLE data: %w", err)
			}
			rleSize := int(rleBuf[0])<<8 | int(rleBuf[1])
			value := rleBuf[2]

			rleData := make([]byte, rleSize)
			for i := range rleData {
				rleData[i] = value
			}

			_, err = out.Seek(int64(offset), io.SeekStart)
			if err != nil {
				return fmt.Errorf("failed to seek to offset %d: %w", offset, err)
			}
			_, err = out.Write(rleData)
			if err != nil {
				return fmt.Errorf("failed to write RLE data: %w", err)
			}
		} else {
			data := make([]byte, size)
			_, err = io.ReadFull(patchReader, data)
			if err != nil {
				return fmt.Errorf("failed to read patch data: %w", err)
			}

			_, err = out.Seek(int64(offset), io.SeekStart)
			if err != nil {
				return fmt.Errorf("failed to seek to offset %d: %w", offset, err)
			}

			_, err = out.Write(data)
			if err != nil {
				return fmt.Errorf("failed to write patch data: %w", err)
			}
		}
	}

	return nil
}
