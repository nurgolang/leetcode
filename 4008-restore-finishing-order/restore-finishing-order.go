func recoverOrder(order []int, friends []int) []int {
    friendsSet := make(map[int]struct{})

    for _, friendID := range friends {
        friendsSet[friendID] = struct{}{}
    }

    res := make([]int, 0, len(friends))

    for _, orderID := range order {
        // Using the "comma-ok idiom" to check for key existence safely and efficiently (O(1)).
        if _, exists := friendsSet[orderID]; exists {
            res = append(res, orderID)
        }  
    }
    return res
}