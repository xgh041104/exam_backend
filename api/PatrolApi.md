
### 巡考账号登录
请求url  http://0.0.0.0:7566/Patrol/LoginPatrolUser
请求方式  post 

请求参数  post
 ``` json
{
    "patrolUserAccount": "admin",
    "patrolUserPwd": "123"
}

 ```

 ``` json
{
    "code": 1,
    "data": {
        "patrolUserId": 1,
        "patrolUserAccount": "admin",
        "patrolUserPwd": "202cb962ac59075b964b07152d234b70"
    },
    "msg": "登录成功",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6ImFkbWluIiwiVXNlclR5cGUiOjAsImV4cCI6MTcyNDkyMTE4OSwiaXNzIjoibHgtand0In0.Wqt_niFNDdRqlfZyW9N-P_SLJrLbDRE1T0hKQtDSrrY"
}
```
### 查询学生登录情况
请求url  http://0.0.0.0:7566/Patrol/GetLoginLog

请求方式  get 

请求参数 
无
返回数据

``` json
{
    "code": 1,
    "data": {
        "nowDay": [
            {
                "hour": "00",
                "day": "",
                "total": 1,
                "dateType": 1
            },
            {
                "hour": "01",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "02",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "03",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "04",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "05",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "06",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "07",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "08",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "09",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "10",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "11",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "12",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "13",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "14",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "15",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "16",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "17",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "18",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "19",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "20",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "21",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "22",
                "day": "",
                "total": 0,
                "dateType": 1
            },
            {
                "hour": "23",
                "day": "",
                "total": 0,
                "dateType": 1
            }
        ],
        "yesterday": [
            {
                "hour": "00",
                "day": "",
                "total": 1,
                "dateType": 2
            },
            {
                "hour": "01",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "02",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "03",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "04",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "05",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "06",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "07",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "08",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "09",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "10",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "11",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "12",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "13",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "14",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "15",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "16",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "17",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "18",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "19",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "20",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "21",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "22",
                "day": "",
                "total": 0,
                "dateType": 2
            },
            {
                "hour": "23",
                "day": "",
                "total": 0,
                "dateType": 2
            }
        ],
        "sevenday": [
            {
                "hour": "",
                "day": "2024-09-06",
                "total": 0,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-07",
                "total": 0,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-08",
                "total": 0,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-09",
                "total": 6,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-10",
                "total": 0,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-11",
                "total": 0,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-12",
                "total": 1,
                "dateType": 3
            },
            {
                "hour": "",
                "day": "2024-09-13",
                "total": 1,
                "dateType": 3
            }
        ]
    },
    "msg": "操作成功"
}
``` 


### 添加人脸监视 

请求url  http://0.0.0.0:7566/Patrol/AddFaceMonitor

请求方式  post 

请求参数  formdata

```json

files   图片

data 
{
    "studentId": 2
}

```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 
 
### 查询人脸监视图片  



请求url  http://0.0.0.0:7566/Patrol/QueryFaceMonitor?studentId=2

请求方式  get 

请求参数  studentId

返回数据

``` json
{
    "code": 1,
    "data": "Resources/FaceMonitor/2/1.jpg",
    "msg": "操作成功"
}
``` 


### 添加作弊预警


请求url  http://0.0.0.0:7566/Patrol/AddCheatWarning
请求方式  post 

请求参数  formdata

```json

files   图片

data 
{
    "studentId": 1,
    "studentName": "aaa"
}

```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 








### 查询作弊预警
请求url  http://0.0.0.0:7566/Patrol/QueryCheatWarning?studentId=1

请求方式  get 

请求参数  studentId

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "cheatWarningId": 4,
            "studentId": 1,
            "studentName": "aaa",
            "studentImgPath": "Resources/CheatWarning/1/85c56182b7d9f217573aa74d8e50a96c.png",
            "studentImgNum": 0,
            "createTime": 1725347502
        },
        {
            "cheatWarningId": 5,
            "studentId": 1,
            "studentName": "aaa",
            "studentImgPath": "Resources/CheatWarning/1/1072e9e32481344a10425954ff3ab937.png",
            "studentImgNum": 0,
            "createTime": 1725347503
        },
        {
            "cheatWarningId": 6,
            "studentId": 1,
            "studentName": "aaa",
            "studentImgPath": "Resources/CheatWarning/1/941b9fbc796a4d264cafbe1cb7af8973.png",
            "studentImgNum": 0,
            "createTime": 1725347504
        },
        {
            "cheatWarningId": 7,
            "studentId": 1,
            "studentName": "aaa",
            "studentImgPath": "Resources/CheatWarning/1/2118c755cf10f720103991383279b920.png",
            "studentImgNum": 0,
            "createTime": 1725347510
        },
        {
            "cheatWarningId": 8,
            "studentId": 1,
            "studentName": "aaa",
            "studentImgPath": "Resources/CheatWarning/1/96ef13d6a8a84460296114e767505994.png",
            "studentImgNum": 0,
            "createTime": 1725347535
        }
    ],
    "msg": "操作成功"
}
``` 

### 删除作弊预警
  请求url http://0.0.0.0:7566/Patrol/DelCheatWarning?studentId=1

请求方式  delete 

请求参数  studentId

返回数据

```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}

```
 





### 查看考试清单1 get
http://0.0.0.0:7566/Patrol/GetPatrolExamPlan 
无参数

返回值
``` json
{
    "code": 1,
    "data": [
        {
            "id": 1,
            "startTime": "2024-6-16 11:10:41",
            "endTime": "2025-6-16 11:10:54",
            "testPaperName": "计划测试1-试卷",
            "examName": "计划测试1-考试",
            "status": "正在进行",
            "personCount": 5
        }
    ],
    "msg": "操作成功"
}
```

### 查看考试清单2  get
http://0.0.0.0:7566/Patrol/GetPatrolExamOverView 
无参数

返回值
``` json
{
    "code": 1,
    "data": [
        {
            "ExamSessionId": 1,
            "ExamId": 1,
            "ExamSession": "计划测试1-考试/计划测试1-试卷",
            "MajorName": "",
            "Score": "",
            "StartTime": "",
            "AvgScore": 0,
            "UnexaminedNum": 0,
            "ExaminedNum": 0,
            "ExamDuration": 0,
            "ExamPeopleSumNum": 0
        }
    ],
    "msg": "操作成功"
}
```
### 查看考试清单详情 get
http://0.0.0.0:7566/Patrol/GetPatrolExamSessionResult?ExamSessionId=1

请求参数 
ExamSessionId 考场id

返回值

``` json

{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "TrueName": "林琛",
            "IDNumber": "12345",
            "StudentId": 2,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "UseTime": 0,
            "Score": -1,
            "ExamStatus": 0,
            "ExamImageArr": [],
            "ExamType": 0
        },
        {
            "Id": 4,
            "TrueName": "殷宏",
            "IDNumber": "1",
            "StudentId": 4,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "UseTime": 0,
            "Score": -1,
            "ExamStatus": 0,
            "ExamImageArr": [],
            "ExamType": 0
        },
        {
            "Id": 3,
            "TrueName": "宋坤",
            "IDNumber": "song123456",
            "StudentId": 3,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "2024/6/16 20:52:57",
            "EndExamTime": "2024-06-16 20:53:07",
            "UseTime": 0,
            "Score": 0,
            "ExamStatus": 1,
            "ExamImageArr": [],
            "ExamType": 0
        },
        {
            "Id": 1,
            "TrueName": "杨双权",
            "IDNumber": "421023199208281011",
            "StudentId": 1,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "2024/6/16 20:22:38",
            "EndExamTime": "2024-06-16 20:22:54",
            "UseTime": 0,
            "Score": 10,
            "ExamStatus": 1,
            "ExamImageArr": [],
            "ExamType": 0
        },
        {
            "Id": 5,
            "TrueName": "祁博",
            "IDNumber": "422202199110123456",
            "StudentId": 5,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "2024/6/16 16:01:08",
            "EndExamTime": "2024-06-16 16:04:54",
            "UseTime": 0,
            "Score": 10,
            "ExamStatus": 1,
            "ExamImageArr": [],
            "ExamType": 0
        }
    ],
    "msg": "操作成功"
}

```

###  查询登记照片  get
http://0.0.0.0:7566/Patrol/GetStudentIdImg?studentId=1

请求参数   
 studentId 学生id

 返回值 
 ```json
{
    "code": 1,
    "data": "Resources/IDImg/cb65fc22d9f381df88f7cf15c9efdd4e.jpg",
    "msg": "操作成功"
}

 ```



### 实时考试情况人次查询 get

http://0.0.0.0:7566/Patrol/GetRealTimeExamSituation

无参数  

 返回值 
 ```json
{
    "code": 1,
    "data": {
        "sumCount": 5,
        "doneCount": 3,
        "unDoneCount": 2
    },
    "msg": "操作成功"
}
 ```

### 查询各个站点的考试信息  get

http://0.0.0.0:7566/Patrol/GetStandGL

无参数

 返回值 
 ```json
{
    "code": 1,
    "data": [
        {
            "id": 1,
            "standName": "测试站点",
            "schoolId": 1,
            "teacherId": 1,
            "postion": "22,11",
            "personCount": 5,
            "doneCount": 3,
            "unDoneCount": 2
        }
    ],
    "msg": "操作成功"
}
 ```