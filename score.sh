#!/bin/bash

#echo "This custom script can run any bash command, and perform tests."
#echo "It needs to only output on1 line in the format 'FS_SCORE:xx%', where xx is the percentage score for the solution."
#echo "FS_SCORE:70%" 

go test github.com/b9lab/other-world/x/otherworld/types
return1=$?
if [ $return1 -ge 1 ]; then
    result1=0
else
    result1=50
fi

go test github.com/b9lab/other-world/x/otherworld/keeper
return2=$?
if [ $return2 -ge 1 ]; then
    result2=0
else
    result2=50
fi

((score=result1+result2))

echo "FS_SCORE:"$score"%"
