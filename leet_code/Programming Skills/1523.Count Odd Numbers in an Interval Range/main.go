func countOdds(low int, high int) int {
    diff := high - low
    if (low%2==0 && high%2==0){
        return diff/2
    }
    return diff/2+1
}