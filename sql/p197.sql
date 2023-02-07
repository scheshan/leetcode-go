select A.id
from Weather A
join Weather B on A.recordDate = date_add(B.recordDate, interval 1 day)
where A.Temperature > B.Temperature