select A.symbol as metal , B.symbol as nonmetal
from Elements A
join Elements B
where A.type = 'Metal' and B.type = 'Nonmetal'