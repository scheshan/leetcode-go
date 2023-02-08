select A.product_name, B.year, B.price
from Product A
join Sales B on A.product_id = B.product_id