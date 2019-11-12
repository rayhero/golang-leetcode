func distanceBetweenBusStops(distance []int, start int, destination int) int {
    //clock
    var sIndex, lIndex, c_sum, cc_sum, sum int = 0, 0, 0, 0, 0;
    if destination > start {
        sIndex = start
        lIndex = destination
    } else {
        sIndex = destination
        lIndex = start
    }
    for i:=sIndex; i < lIndex; i++ {
        c_sum = c_sum + distance[i]
    }
    // total
    for _,v := range distance {
        sum = sum + v
    }
    
    cc_sum = sum - c_sum
    
    
    if c_sum < cc_sum {
        return c_sum
    }
    return cc_sum
}