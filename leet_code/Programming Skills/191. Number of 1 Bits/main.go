func hammingWeight(num uint32) int {
    result:=0
    s:=strconv.FormatUint(uint64(num),2)
    for _,v:=range s{
        if v=='1' {
            result++
        }  
    }
    return result
    
}