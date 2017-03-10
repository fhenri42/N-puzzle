echo "Npuzzle checker"
i=0

for i in `seq 1 30`
do
r=$(( $RANDOM % 10 + 2));
python test.py -u $r > n && ./n-puzzle n
i=$((i + 1))
#statements
done
echo "========================="
echo "========================="
echo "========================="
echo "========================="
echo "valide map"
i = 0
for i in `seq 1 30`
do
r=3
python test.py -s $r > n && time ./n-puzzle n
i=$((i + 1))
#statements
done
