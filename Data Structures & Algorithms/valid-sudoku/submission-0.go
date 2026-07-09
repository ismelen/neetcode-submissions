func isValidSudoku(board [][]byte) bool {
    for row := range 3 {
        for column := range 3 {
            seen := make(map[byte]bool)

            for i:=row*3; i < (row+1)*3; i++ {
                for j:=column*3; j < (column+1)*3; j++ {
                    value := board[i][j]
                    if string(value) != "." {
                        if seen[value] {
                            return false
                        }
                        seen[value] = true
                    }
                }
            } 
        }
    }
    
    for i := range 9 {
        seenInRow := make(map[byte]bool)
        seenInColumn := make(map[byte]bool)

        for j := range 9 {
            value := board[i][j]
            if string(value) != "." {
                if seenInRow[value] {
                    return false
                }
                seenInRow[value] = true
            }

            value = board[j][i]
            if string(value) != "." {
                if seenInColumn[value] {
                    return false
                }
                seenInColumn[value] = true
            }
        }
    }

    return true
}

