def smallest_k(matrix, k):
    n = len(matrix)
    l = matrix[0][0]
    r = matrix[n - 1][n - 1]

    while l < r:
        mid = (l + r) // 2

        count = 0
        row, col = 0, n - 1

        while row < n and col >= 0:
            if matrix[row][col] <= mid:
                count += col + 1
                row += 1
            else:
                col -= 1

        if count < k:
            l = mid + 1
        else:
            r = mid

    return l

matrix = [
    [1, 5, 9, 11],
    [2, 6, 10, 12],
    [3, 7, 13, 14],
    [4, 8, 15, 16]
]

print(smallest_k(matrix, 8))