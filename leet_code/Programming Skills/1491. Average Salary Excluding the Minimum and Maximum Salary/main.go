func average(salary []int) float64 {
    sum:=0
    min:=10000000
    max:=100
    for _,v:= range salary{
        sum= sum+v
        if v<min{
            min=v
        }
        if v>max{
            max=v
        }
    }
    return float64(sum-min-max)/float64(len(salary)-2)
    
}