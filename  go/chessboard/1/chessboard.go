package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f := cb[file]
	occupied := 0
	for _, square := range f {
		if square {
			occupied++
		}
	}
	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	rank-- // convert from 1 to 0-based indexing
	occupied := 0
	if 0 <= rank && rank <= 7 {
		for _, f := range cb {
			if f[rank] {
				occupied++
			}
		}
	}
	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squares := 0
	for range cb {
		squares += 8
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0
	for _, f := range cb {
		for _, rank := range f {
			if rank {
				occupied++
			}
		}
	}
	return occupied
}
