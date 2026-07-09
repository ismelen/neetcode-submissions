func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    boxes := make([]map[byte]bool, 9)

    for i := range 9 {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := range 9 {
        for c := range 9 {
            val := board[r][c]

            if val == '.' {
                continue
            }

            boxIdx := (r/3)*3 + c/3
            if rows[r][val] || cols[c][val] || boxes[boxIdx][val] {
                return false
            } 

            rows[r][val] = true
            cols[c][val] = true
            boxes[boxIdx][val] = true
        }
    }

    return true
}
