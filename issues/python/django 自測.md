
```python
python  manage.py shell < test_local.py
```

內容可以如下:

leapy core 測試 ip
```
from gis.common.location import location_service

a = location_service.get_province_city("127.0.0.1", True)
print(a)
```