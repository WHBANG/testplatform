
### 镜像接口

1. 新增镜像   
POST: v1/image  

2. 修改镜像  
PUT: v1/image/{image_id}    

3. 删除镜像  
DEL: v1/image/{image_id}  

4. 查看镜像  
GET: v1/image/{image_id}  
 
5. 镜像列表   
GET: v1/image?page=&size=&query=...    
区分analyzer-flow镜像和analyzer-io镜像   

### 引擎镜像接口   
有关引擎启动停止接口   

1. 部署引擎
POST: v1/image/{id}/deploy  

2. 停止引擎
POST: v1/image/{id}/stop  

### 测试用例接口CRUD


### 测试接口  
新增测试，查看测试列表等

1. 新增测试
POST: v1/test   

2. 查看测试  
GET: v1/test/{id} 

3. 删除测试  
DELETE: v1/test/{id}  

4. 获取测试列表  
GET: v1/test?query=  


