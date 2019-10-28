def printAllKLength(set, k): 
    n = len(set)  
    printAllKLengthRec(set, "", n, k) 
  
def printAllKLengthRec(set, prefix, n, k): 
    if (k == 0) : 
        f.write(prefix+"\n")
        #print(prefix)
        return
    for i in range(n): 
        newPrefix = prefix + set[i] 
        printAllKLengthRec(set, newPrefix, n, k - 1) 
  
if __name__ == "__main__": 
    f= open("all_combinations.txt","w+")
    ttt_set = ['x', 'o', '-'] 
    k = 9
    printAllKLength(ttt_set, k)