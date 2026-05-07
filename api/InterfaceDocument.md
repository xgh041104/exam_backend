# 学习考试平台接口文档


## 公共接口
#### 获取静态资源应该从那个路径获取
#### 获取是否开启全局人脸识别

## 学生端

### 未登录需要的接口
#### 公共列表 （不需要token就能请求）获取日期前五的公告

http://192.168.0.123:7566/Manage/NoticeList

请求类型 get

请求参数
 无

返回结果


``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "Time": "2022-05-02 10:23:20",
            "NoticeTitle": "222",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "1"
        },
        {
            "Id": 1,
            "Time": "2020-05-01 10:20:30",
            "NoticeTitle": "aaa",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "2"
        }
    ],
    "msg": "成功"
}
```




#### 学生登录
 请求地址  http://192.168.0.123:7566/Student/LoginStudent

 请求方式 post

``` json
{
    "StudentAccount":"dll",
    "StudentPwd":"123456"
}
```
返回数据
``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "StudentType": 1,
        "StudentAccount": "dll",
        "StudentPwd": "e10adc3949ba59abbe56e057f20f883e",
        "StandId": 1,
        "ExamName": "考试用户1",
        "SchoolId": 1,
        "CollegeId": 1,
        "MajorId": 1,
        "ClassId": 2,
        "TrueName": "杜拉拉2",
        "IDNumber": "111121212",
        "ExamNumber": "45454545454",
        "Birthday": "1988-02-12",
        "Phone": "15922222222",
        "Email": "284881441@qq.com",
        "IDImage": "1",
        "FaceOpen": 1,
        "Sex": 0
    },
    "msg": "登录成功",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6IuadnOaLieaLiTIiLCJleHAiOjE2OTE1NTYxNzUsImlzcyI6Imx4LWp3dCJ9.Q7L2gzkslptCKI2YdJigHBewUHqIONS92Bo47obcLVQ"
}
```



#### 学生个人信息查询
请求地址  http://192.168.0.123:7566/Student/GetStudentViewById?StudentId=1

 请求方式 get

请求参数 StudentId=1
 
返回数据
```json
{
    "code": 1,
    "data": {
        "Id": 1,
        "StudentType": 1,
        "StudentAccount": "dll",
        "StudentPwd": "e10adc3949ba59abbe56e057f20f883e",
        "StandId": 1,
        "ExamName": "考试用户1",
        "SchoolId": 1,
        "CollegeId": 1,
        "MajorId": 1,
        "ClassId": 2,
        "TrueName": "杜拉拉2",
        "IDNumber": "111121212",
        "ExamNumber": "45454545454",
        "Birthday": "1988-02-12",
        "Phone": "15922222222",
        "Email": "284881441@qq.com",
        "IDImage": "1",
        "FaceOpen": 1,
        "SchoolName": "清华大学",
        "CollegeName": "美术学院",
        "MajorName": "计算机科学与技术",
        "ClassName": "宝石1班",
        "StandName": "武汉国土1",
        "Sex": 0
    },
    "msg": "操作成功"
}

```
#### 学生个人信息修改

请求url  http://192.168.0.123:7566/Student/EditStudent

请求方式  post 

请求参数 

``` json
{
	"StudentType":1,
	"StudentAccount":"dll",
	"StudentPwd":"123",
	"StandId":1,
	"ExamName":"考试用户1",
	"SchoolId":2,
	"CollegeId":1,
	"MajorId":1,
	"ClassId":1,
	"TrueName":"杜拉拉2",
	"IDNumber":"111121212",
	"ExamNumber":"45454545454",
	"Birthday":"1988-02-12",
	"Phone":"15922222222",
	"Email":"284881441@qq.com",
	"IDImage":"1",
    "FaceOpen":"1",
    "Id":2
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

### 课程中心
#### 学生课程列表 

//正在学习 和未开始学习

请求地址  http://192.168.0.123:7566/Student/GetCurrentStudyCourse?StudentId=5

 请求方式 get

请求参数 StudentId=5

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "CourseName": "课程1",
            "Digest": "",
            "CourseId": 24,
            "TeacherId": 1,
            "TeacherName": "tea",
            "SchoolName": "清华大学",
            "CollegeName": "美术学院",
            "MajorId": 1,
            "MajorName": "计算机科学与技术",
            "ChapterSum": 1,
            "StudentSum": 1,
            "CourseStartTime": "2023-06-01",
            "CourseEndTime": "2023-06-08",
            "ChapterOrder": "0",
            "LearningRate": "0",
            "FilePath": "",
            "IsCurrentStudy": 0   //0 是未开始学习的课程  1是正在学习的课程  2已完成的课程 3是已过期的课程
        },
        {
            "CourseName": "test",
            "Digest": "",
            "CourseId": 8,
            "TeacherId": 1,
            "TeacherName": "tea",
            "SchoolName": "清华大学",
            "CollegeName": "美术学院",
            "MajorId": 1,
            "MajorName": "计算机科学与技术",
            "ChapterSum": 1,
            "StudentSum": 1,
            "CourseStartTime": "2023-06-21 16:07:00",
            "CourseEndTime": "2023-06-25 16:07:00",
            "ChapterOrder": "0",
            "LearningRate": "0",
            "FilePath": "Resources/Img/1a77d57e19a1eab2461afc4b9c4565dc.png",
            "IsCurrentStudy": 0
        }
    ],
    "msg": "成功"
}
```





#### 课程详情  章节列表（显示自己学到当前课程的某个章节的进度）


 请求地址  http://192.168.0.123:7566/Student/CourseDetails?StudentId=5&

 请求方式 get

请求参数 
StudentId=5
CourseId=25

返回数据
``` json
{
    "code": 1,
    "data": {
        "CourseName": "课程3",
        "Digest": "",
        "CourseId": 25,
        "TeacherId": 1,
        "TeacherName": "tea",
        "SchoolName": "清华大学",
        "CollegeName": "美术学院",
        "MajorId": 1,
        "MajorName": "计算机科学与技术",
        "ChapterSum": 1,
        "StudentSum": 1,
        "CourseStartTime": "2023-06-08",
        "CourseEndTime": "2023-06-15",
        "ChapterOrder": 0,
        "LearningRate": 0,
        "FilePath": "",
        "IsCurrentStudy": 0,
        "ChapterList": [
            {
                "Id": 3,
                "ChapterName": "test",
                "ChapterType": "图文课",
                "ChapterOrder": 1
            }
        ]
    },
    "msg": "成功"
}
```

#### 课程查询学生进度 根据课程id查询
请求地址  http://192.168.0.124:7566/Manage/GetCouresStudyPlanByCourseId?CourseId=25

 请求方式 get

请求参数 
CourseId=25

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 13,
            "StudentId": 4,
            "TrueName": "杜靠靠1",
            "CourseId": 25,
            "CourseName": "课程3",
            "ChapterId": 0,
            "LearningRate": 0,
            "ChapterOrder": 0,
            "ChapterNum": 5,
            "IsComplete": 0
        },
        {
            "Id": 14,
            "StudentId": 5,
            "TrueName": "杜靠靠2",
            "CourseId": 25,
            "CourseName": "课程3",
            "ChapterId": 3,
            "LearningRate": 20,
            "ChapterOrder": 1,
            "ChapterNum": 5,
            "IsComplete": 0
        }
    ],
    "msg": "操作成功"
}
```
####  章节详情  
请求url  http://192.168.0.123:7566/Student/GetChapterById?ChapterId=2

请求方式  get 

请求参数 
ChapterId=2

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 2,
        "ChapterName": "test",
        "CourseId": 25,
        "ChapterType": "1",
        "ChapterOrder": 0,
        "ChapterContent": "Resources/Video/7db48f7b4a647196b94345eea722d034/7db48f7b4a647196b94345eea722d034.m3u8",
        "FileInfo": [
            {
                "Id": 13,
                "FileType": ".jpg",
                "FileName": "",
                "FileUseTo": "",
                "SchoolId": 0,
                "FilePath": "Resources/Img/35c9580b804a56aa9cb46c332f7b4cd5.jpg"
            }
        ]
    },
    "msg": "操作成功"
}
```




#### 学习章节之后 进度上传
请求地址  http://192.168.0.123:7566/Student/StudyPlanUpload

 请求方式 post·

``` json
{
    "ChapterId": 3,
    "LearningRate": 20,
    "ChapterOrder": 1,
    "IsComplete": 0,
    "CourseId": 25,
    "StudentId": 5
}
```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 

安全问题
实时播放问题 





 
### 考试中心

#### 学生考试列表
请求url  http://192.168.0.123:7566/Student/GetStudentExamInfo?StudentId=6

请求方式  get 

请求参数 
StudentId=6

返回数据

  ExamZT
	// -- 0 已过期
	// -- 1 已完成
	// -- 2 未参加考试
	// -- 3 考试未开始
	// -- 4 未补考   
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 9,
            "StudentId": 6,
            "ExamId": 1,
            "ExamName": "测试试卷35",
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 1,
            "ExamZT": 0,
            "MajorId": 1,
            "CourseId": 0,
            "MajorName": "计算机科学与技术",
            "CourseName": "",
            "FullMarks": 100,
            "PassScore": 60,
            "SessionNum": 1,
            "ExamDuration": 60,
            "QuestionNum": 2,
            "ResetStartExamTime": "2023-07-20",
            "ResetEndExamTime": "2023-07-23",
            "SessionStartExamTime": "2021-2-3",
            "SessionEndExamTime": "2023-2-3",
            "FaceVerify":1
        },
        {
            "Id": 10,
            "StudentId": 6,
            "ExamId": 1,
            "ExamName": "测试试卷35",
            "ExamSessionId": 2,
            "StartExamTime": "",
            "EndExamTime": "",
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 1,
            "ExamZT": 0,
            "MajorId": 1,
            "CourseId": 0,
            "MajorName": "计算机科学与技术",
            "CourseName": "",
            "FullMarks": 100,
            "PassScore": 60,
            "SessionNum": 2,
            "ExamDuration": 60,
            "QuestionNum": 2,
            "ResetStartExamTime": "2023-07-20",
            "ResetEndExamTime": "2023-07-29",
            "SessionStartExamTime": "2021-2-3",
            "SessionEndExamTime": "2023-2-3",
              "FaceVerify":1
        }
    ],
    "msg": "操作成功"
}
```
#### 查看已考试的考试详情
#### 学生考试（获取该试卷所有试题按题型分类）


请求url  http://192.168.0.123:7566/Student/GetStudentExamDetails?StudentId=6&ExamId=1&ExamSessionId=2

请求方式  get 

请求参数 
StudentId=6
ExamId=1
ExamSessionId=2

返回数据
``` json
{
    "code": 1,
    "data": {
        "Id": 10,
        "StudentId": 6,
        "ExamId": 1,
        "ExamName": "测试试卷35",
        "ExamSessionId": 2,
        "TestPaperId": 4,
        "TestPaperName": "试卷测试2",
        "TestPaperType": 2,
        "ExamType": 0,
        "MajorName": "",
        "CourseName": "",
        "FullMarks": 100,
        "PassScore": 60,
        "ExamDuration": 100,
        "SessionNum": 1,
        "SchoolName": "清华大学",
        "TeacherName": "tea",
        "FaceVerify":1,
        "TestPaperQuestionTypeOver": [
            {
                "QuestionType": 1,
                "QuestionIdNum": 2,
                "QuestionScore": 70
            },
            {
                "QuestionType": 2,
                "QuestionIdNum": 1,
                "QuestionScore": 30
            },
            {
                "QuestionType": 3,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 4,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 5,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            }
        ],
        "TestPaperQuestionViewFile": [
            {
                "QuestionId": 2,
                "QuestionName": "单选题",
                "QuestionPoolId": 1,
                "QuestionType": 1,
                "QuestionContent": "1111",
                "QuestionScore": 34,
                "Digree": 1,
                "Answer": "1",
                "FileInfo": [
                    {
                        "Id": 182,
                        "FileType": ".zip",
                        "FileName": "api.zip",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/4809a7968dfceb121618db6b7cd80786.zip"
                    }
                ]
            },
            {
                "QuestionId": 6,
                "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
                "QuestionPoolId": 1,
                "QuestionType": 1,
                "QuestionContent": "[\"A. 格斯拉兄弟将自己的店开在一条横街上，这条横街座落在伦敦市的西区。\",\"B. 我不断地叩问历史的发展，产生了从中找出一条贯穿其中的脉络的愿望。\",\"C. 老王为新春写的这幅对联，寄托了美好的祝福，展现了书法大境界。\"]",
                "QuestionScore": 36,
                "Digree": 4,
                "Answer": "1",
                "FileInfo": [
                    {
                        "Id": 186,
                        "FileType": ".pdf",
                        "FileName": "9.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/1e2699ae969814aa53c585a279a1f0ec.pdf"
                    }
                ]
            },
            {
                "QuestionId": 16,
                "QuestionName": "多选题测试",
                "QuestionPoolId": 1,
                "QuestionType": 2,
                "QuestionContent": "[\"payload\",\"payload\",\"正确\",\"正确\"]",
                "QuestionScore": 30,
                "Digree": 3,
                "Answer": "[3,2]",
                "FileInfo": [
                    {
                        "Id": 193,
                        "FileType": ".pdf",
                        "FileName": "1.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/7e47b83614b5b3d4f37ad61b20d940f7.pdf"
                    },
                    {
                        "Id": 194,
                        "FileType": ".pdf",
                        "FileName": "4.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/2cb075d5b6f5c06eb8b45601616e9b5b.pdf"
                    }
                ]
            }
        ]
    },
    "msg": "操作成功"
}
```
#### 提交考试
请求地址  http://192.168.0.123:7566/Student/ExamStudentSumit

 请求方式 post

``` json
{
    "StudentId": 1,
    "ExamId": 1,
    "ExamSessionId": 1,
    "TestPaperId": 1,
    "StartExamTime": "2023-08-01 16:59:20",
    "Score": 20,
    "IsReTest":0,
    "AnswerSheetArr": [
        {
            "QuestionId": 2,
            "AnswerScore": 0.5,
            "AnswerSteps": "A",
            "IsTrue": 1
        }, {
            "QuestionId": 3,
            "AnswerScore": 0.5,
            "AnswerSteps": "A",
            "IsTrue": 1
        } 
    ]
}
```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 

#### 上传考试中的图片
请求地址  http://192.168.0.124:7566/Student/UploadExamImage

 请求方式 post
``` js
var formData1 = new FormData();
formData1.set("files",file); //file 文件名为idnumber或examnumber  匹配不上则在返回的data里面
formData1.set("data","{
    "ExamId":1,
    "ExamSessionId":1,
    "StudentId":1
}")
```
返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 

#### 查看考试详情 每题和每题的答案

http://192.168.0.124:7566/Student/GetStudentExamPaperOver?StudentId=6&ExamSessionId=1&ExamId=1
请求方式  get 

请求参数 
StudentId=6
ExamId=1
ExamSessionId=2

返回数据


``` json

{
    "code": 1,
    "data": {
        "Id": 9,
        "StudentId": 6,
        "ExamId": 1,
        "ExamName": "测试试卷35",
        "ExamSessionId": 1,
        "TestPaperId": 5,
        "TestPaperName": "试卷测试1",
        "TestPaperType": 1,
        "ExamType": 0,
        "MajorName": "",
        "CourseName": "",
        "FullMarks": 100,
        "PassScore": 60,
        "ExamDuration": 100,
        "SessionNum": 3,
        "SchoolName": "清华大学",
        "TeacherName": "tea",
        "Score": 1,
        "StartExamTime": "2023/8/9 15:19:04",
        "EndExamTime": "2023-08-09 15:21:36",
        "AnswerSheetarr": [
            {
                "QuestionId": 2,
                "AnswerScore": 0.5,
                "AnswerSteps": "A",
                "IsTrue": 1
            },
            {
                "QuestionId": 3,
                "AnswerScore": 0.5,
                "AnswerSteps": "A",
                "IsTrue": 1
            }
        ],
        "TestPaperQuestionTypeOver": [
            {
                "QuestionType": 1,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 2,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 3,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 4,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 5,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            }
        ],
        "TestPaperQuestionViewFile": [
            {
                "QuestionId": 2,
                "QuestionName": "单选题",
                "QuestionPoolId": 1,
                "QuestionType": 1,
                "QuestionContent": "1111",
                "QuestionScore": 34,
                "Digree": 1,
                "Answer": "1",
                "FileInfo": [
                    {
                        "Id": 182,
                        "FileType": ".zip",
                        "FileName": "api.zip",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/4809a7968dfceb121618db6b7cd80786.zip"
                    }
                ]
            },
            {
                "QuestionId": 6,
                "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
                "QuestionPoolId": 1,
                "QuestionType": 1,
                "QuestionContent": "[\"A. 格斯拉兄弟将自己的店开在一条横街上，这条横街座落在伦敦市的西区。\",\"B. 我不断地叩问历史的发展，产生了从中找出一条贯穿其中的脉络的愿望。\",\"C. 老王为新春写的这幅对联，寄托了美好的祝福，展现了书法大境界。\"]",
                "QuestionScore": 36,
                "Digree": 4,
                "Answer": "1",
                "FileInfo": [
                    {
                        "Id": 186,
                        "FileType": ".pdf",
                        "FileName": "9.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/1e2699ae969814aa53c585a279a1f0ec.pdf"
                    }
                ]
            },
            {
                "QuestionId": 16,
                "QuestionName": "多选题测试",
                "QuestionPoolId": 1,
                "QuestionType": 2,
                "QuestionContent": "[\"payload\",\"payload\",\"正确\",\"正确\"]",
                "QuestionScore": 30,
                "Digree": 3,
                "Answer": "[3,2]",
                "FileInfo": [
                    {
                        "Id": 193,
                        "FileType": ".pdf",
                        "FileName": "1.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/7e47b83614b5b3d4f37ad61b20d940f7.pdf"
                    },
                    {
                        "Id": 194,
                        "FileType": ".pdf",
                        "FileName": "4.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/2cb075d5b6f5c06eb8b45601616e9b5b.pdf"
                    }
                ]
            }
        ]
    },
    "msg": "操作成功"
}

```
### 模拟训练
#### 根据学校查询题目
请求url  http://192.168.0.123:7566/Student/GetQuestionBySchoolId?SchoolId=1

请求方式  get 

请求参数 
SchoolId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "QuestionId": 2,
            "SchoolId": 1,
            "QuestionPoolId": 1,
            "QuestionName": "实操题",
            "QuestionType": 1,
            "QuestionContent": "",
            "Digree": 1,
            "MajorID": 1,
            "CollegeId": 1,
            "CourseId": 1,
            "Answer": "1",
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": ""
        },
        {
            "QuestionId": 3,
            "SchoolId": 1,
            "QuestionPoolId": 1,
            "QuestionName": "实操题",
            "QuestionType": 5,
            "QuestionContent": "183",
            "Digree": 1,
            "MajorID": 1,
            "CollegeId": 1,
            "CourseId": 1,
            "Answer": "1",
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": ""
        }
    ],
    "msg": "操作成功"
}
```


#### 根据题目id查询题目
请求url  http://192.168.0.123:7566/Student/GetQuestionByQuestionId?QuestionId=2

请求方式  get 

请求参数 
QuestionId=2

返回数据

``` json
{
    "code": 1,
    "data": {
        "QuestionId": 3,
        "SchoolId": 1,
        "QuestionPoolId": 1,
        "QuestionName": "实操题",
        "QuestionType": 5,  //如果是5 就是实操题  fileinfo 就是实操的文件  如果是其他   里面就是附件列表
        "QuestionContent": "183",
        "Digree": 1,
        "MajorID": 1,
        "CollegeId": 1,
        "CourseId": 1,
        "Answer": "1",
        "MajorName": "计算机科学与技术",
        "CollegeName": "美术学院",
        "CourseName": "",
        "FileInfo": [
            {
                "Id": 183,
                "FileType": ".zip",
                "FileName": "api.zip",
                "FileUseTo": "",
                "SchoolId": 0,
                "FilePath": "Resources/Zip/4296e39ef8390472aa3565649d79614e.zip"
            }
        ]
    },
    "msg": "操作成功"
}
```


 
#### 试题练习
#### 归档错题集
请求地址  http://192.168.0.123:7566/Student/AddQuestionWrong

 请求方式 post

``` json
[{
    "QuestionId":6,
    "StudentId": 1,
    "AnswerSteps": "1"
},
{
    "QuestionId":6,
    "StudentId": 1,
    "AnswerSteps": "1"
}
]
```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 

#### 错题集移出
请求地址  http://192.168.0.123:7566/Student/DelQuestionWrong

 请求方式 post

``` json
[{
    "Id":1
},{
    "Id":2
}]
```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
``` 

####  个人错题集列表错题列表 根据学生id查询
请求地址  http://192.168.0.123:7566/Student/GetQuestionWrongByStudentId?StudentId=1

 请求方式 get

请求参数
StudentId=1


返回信息
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "QuestionId": 6,
            "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
            "StudentId": 1,
            "CreateTime": "2023-07-24 17:31:34",
            "AnswerSteps": "1",
            "TrueAnswer": "1",
            "QuestionViewEntity": {
                "QuestionId": 6,
                "SchoolId": 1,
                "QuestionPoolId": 1,
                "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
                "QuestionType": 1,
                "QuestionContent": "[\"A. 格斯拉兄弟将自己的店开在一条横街上，这条横街座落在伦敦市的西区。\",\"B. 我不断地叩问历史的发展，产生了从中找出一条贯穿其中的脉络的愿望。\",\"C. 老王为新春写的这幅对联，寄托了美好的祝福，展现了书法大境界。\"]",
                "Digree": 4,
                "MajorID": 0,
                "CollegeId": 0,
                "CourseId": 0,
                "Answer": "1",
                "MajorName": "",
                "CollegeName": "",
                "CourseName": "",
                "FileInfo": [
                    {
                        "Id": 186,
                        "FileType": ".pdf",
                        "FileName": "9.pdf",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/1e2699ae969814aa53c585a279a1f0ec.pdf"
                    }
                ]
            }
        }
    ],
    "msg": "操作成功"
}
``` 



## 后台

### 教员管理
#### 老师登录
 请求地址  http://192.168.0.123:7566/Manage/LoginTeacher

 请求方式 post

``` json
{
"TeacherAccount":"tea",
"TeacherPassword":"123"
}
```


返回信息
``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "SchoolId": 1,
        "TeacherAccount": "tea",
        "TeacherPassword": "",
        "TeacherName": "tea",
        "Sex": 1,
        "PhoneNumber": "138888888",
        "Email": "2848847444",
        "TeacherTitle": "讲师"
    },
    "msg": "登录成功",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6InRlYSIsImV4cCI6MTY4NTk2MjM5OCwiaXNzIjoibHgtand0In0.nnQNqVxZ1C5HkAFSSzW8BaCJqZyHWp9V5RdZ5eafZgg"
}
``` 



#### 查询所有老师


http://192.168.0.123:7566/Manage/GetTeacherList

请求类型 get

请求参数
 无

返回结果


``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "SchoolId": 1,
            "TeacherAccount": "tea",
            "TeacherName": "tea",
            "Sex": 1,
            "PhoneNumber": "138888888",
            "Email": "2848847444",
            "TeacherTitle": "讲师",
            "SchoolName": "清华大学",
            "SchoolAddress": "五道口"
        }
    ],
    "msg": "操作成功"
}
```


#### 根据老师Id 查询当前老师


http://192.168.0.123:7566/Manage/GetTeacherById?TeacherId=1

请求类型 get

请求参数
 TeacherId=1

返回结果


``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "SchoolId": 1,
        "TeacherAccount": "tea",
        "TeacherName": "tea",
        "Sex": 1,
        "PhoneNumber": "138888888",
        "Email": "2848847444",
        "TeacherTitle": "讲师",
        "SchoolName": "清华大学",
        "SchoolAddress": "五道口"
    },
    "msg": "操作成功"
}
```





#### 新增老师

 请求地址  http://192.168.0.123:7566/Manage/AddTeacher

 请求方式 post

``` json

{
        "SchoolId": 1,
        "TeacherAccount": "teacher",
        "TeacherPassword":"123",
        "TeacherName": "tea",
        "Sex": 1,
        "PhoneNumber": "1388888881",
        "Email": "2848847444",
        "TeacherTitle": "讲师"
    }

```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功",
}
``` 

#### 编辑老师

 请求地址  http://192.168.0.123:7566/Manage/EditTeacher

 请求方式 post

``` json

{
       "TeacherId":2,
        "SchoolId": 1,
        "TeacherName": "teacher",
        "Sex": 1,
        "PhoneNumber": "1388888881",
        "Email": "28488474441111111111",
        "TeacherTitle": "讲师1"
    }

```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功",
}
``` 

#### 删除老师

 请求地址  http://192.168.0.123:7566/Manage/DelTeacher

 请求方式 post

``` json
{
    "TeacherId": 2
}

```


返回信息
``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功",
}
``` 




#### 修改老师密码

请求url  http://192.168.0.123:7566/Manage/EditTeacherPassWordById

请求方式  post 

请求参数 

``` json
{
    "TeacherId": 1,
    "TeacherPassword":"123456"
}
```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "修改密码成功"
}
```
#### 个人信息修改

### 学生用户管理
#### 社会人士导入

#### execl学生表格预览
请求url  http://192.168.0.124:7566/Manage/GetStudentExeclData

请求方式  post 

请求参数 
 formdata  files file

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "TrueName": "许文亚",
            "IDNumber": "21000",
            "Sex": 1,
            "Birthday": "1990-07-04",
            "SchoolName": "武汉长江职院",
            "StudentAccount": "",
            "StudentPwd": "220016",
            "StandName": "",
            "Phone": "",
            "Email": "",
            "Status": 0
        },
        {
            "TrueName": "祝必成",
            "IDNumber": "21001",
            "Sex": 1,
            "Birthday": "1990-03-17",
            "SchoolName": "武汉长江职院",
            "StudentAccount": "",
            "StudentPwd": "220030",
            "StandName": "",
            "Phone": "",
            "Email": "",
            "Status": 0
        }
    ],
    "msg": "操作成功"
}
```



####  地大站点发来的源文件execl学生表格预览
请求url  http://192.168.0.124:7566/Manage/MatchAddSocietyStudentExecl

请求方式  post 

请求参数 
 formdata  files file

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "TrueName": "胡奇",
            "IDNumber": "429004",
            "ExamNumber": "014322220250",
            "ExamName": "",
            "NativPlace": "湖北",
            "Sex": 1,
            "Birthday": "1994-10-01",
            "SchoolName": "",
            "StudentAccount": "014322220250",
            "StudentPwd": "220250",
            "StandName": "",
            "StandId": 1,
            "Phone": "",
            "Email": "",
            "Status": 0
        },
        {
            "TrueName": "周婉银",
            "IDNumber": "421081",
            "ExamNumber": "014322220555",
            "ExamName": "",
            "NativPlace": "湖北",
            "Sex": 0,
            "Birthday": "1995-06-01",
            "SchoolName": "",
            "StudentAccount": "014322220555",
            "StudentPwd": "220555",
            "StandName": "",
            "StandId": 1,
            "Phone": "",
            "Email": "",
            "Status": 0
        }
    ],
    "msg": "操作成功"
}
```




#### 个人登记证文件夹上传
#### 社会人士导出

#### 在校生导入
#### 在校生导出

#### 学院列表 根据学校id查询学院

请求url  http://192.168.0.123:7566/Manage/GetCollegeBySchoolId?SchoolId=1

请求方式  get 

请求参数 SchoolId=1

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "CollegeName": "美术学院",
            "SchoolId": 1,
            "TeacherId": 1,
            "SchoolName": "清华大学"
        },
        {
            "Id": 3,
            "CollegeName": "计算机学院",
            "SchoolId": 1,
            "TeacherId": 1,
            "SchoolName": "清华大学"
        }
    ],
    "msg": "成功"
}
```
#### 学院修改
请求url  http://192.168.0.123:7566/Manage/EditCollege

请求方式  post 

请求参数 
``` json
{
"Id":1,
"CollegeName":"美术学院",
"SchoolId":1
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
#### 学院删除

请求url  http://192.168.0.123:7566/Manage/DelCollege

请求方式  post 

请求参数 
``` json
{
"Id":2
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
#### 学院新增

请求url  http://192.168.0.123:7566/Manage/AddCollege

请求方式  post 

请求参数 
``` json
{
"CollegeName":"计算机学院",
"SchoolId":1,
"TeacherId":1
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




#### 学院列表  查询所有学院 管理员使用

请求url  http://192.168.0.123:7566/Manage/GetCollegeListAll

请求方式  get 

请求参数 
 无

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "CollegeName": "美术学院",
            "SchoolId": 1,
            "TeacherId": 1
        },
        {
            "Id": 3,
            "CollegeName": "计算机学院",
            "SchoolId": 1,
            "TeacherId": 1
        },
        {
            "Id": 4,
            "CollegeName": "音乐学院",
            "SchoolId": 0,
            "TeacherId": 1
        },
        {
            "Id": 5,
            "CollegeName": "音乐学院",
            "SchoolId": 0,
            "TeacherId": 1
        },
        {
            "Id": 6,
            "CollegeName": "音乐学院",
            "SchoolId": 0,
            "TeacherId": 1
        },
        {
            "Id": 7,
            "CollegeName": "音乐学院",
            "SchoolId": 0,
            "TeacherId": 1
        },
        {
            "Id": 8,
            "CollegeName": "音乐学院",
            "SchoolId": 1,
            "TeacherId": 1
        },
        {
            "Id": 9,
            "CollegeName": "生物工程学院",
            "SchoolId": 0,
            "TeacherId": 1
        },
        {
            "Id": 10,
            "CollegeName": "生物工程学院",
            "SchoolId": 1,
            "TeacherId": 1
        },
        {
            "Id": 12,
            "CollegeName": "电子科学与工程学院",
            "SchoolId": 1,
            "TeacherId": 1
        },
        {
            "Id": 13,
            "CollegeName": "物理学院",
            "SchoolId": 1,
            "TeacherId": 1
        }
    ],
    "msg": "成功"
}

```




#### 学院列表  查询所有学院 

请求url  http://192.168.0.123:7566/Manage/GetCollegeByCollegeId?CollegeId=1

请求方式  get 

请求参数 
CollegeId=1

返回数据
``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "CollegeName": "美术学院",
        "SchoolId": 1,
        "TeacherId": 1
    },
    "msg": "成功"
}
```







#### 专业列表   根据学院id查询专业
请求url  http://192.168.0.123:7566/Manage/GetMajorByCollegeId?CollegeId=1

请求方式  get 

请求参数  CollegeId=1    

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "MajorId": 1,
            "SchoolId": 1,
            "CollegeId": 1,
            "MajorName": "计算机科学与技术",
            "TeacherId": 1
        }
    ],
    "msg": "成功"
}
```
#### 专业列表   根据学校id 查询专业列表
请求url  http://192.168.0.123:7566/Manage/GetMajorBySchoolId?SchoolId=1

请求方式  get 

请求参数  SchoolId=1

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "MajorId": 1,
            "SchoolId": 1,
            "CollegeId": 1,
            "MajorName": "计算机科学与技术",
            "TeacherId": 1,
            "CollegeName": "美术学院"
        }
    ],
    "msg": "成功"
}
```

#### 专业列表     查询所有专业列表
请求url  http://192.168.0.123:7566/Manage/GetMajorListView

请求方式  get 

请求参数   无

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "MajorId": 1,
            "SchoolId": 1,
            "CollegeId": 1,
            "MajorName": "计算机科学与技术",
            "TeacherId": 1,
            "CollegeName": "美术学院"
        }
    ],
    "msg": "成功"
}
```
#### 专业修改

请求url  http://192.168.0.123:7566/Manage/EditMajor

请求方式  post 

请求参数 
``` json
{
"SchoolId":1,
"CollegeId":1,
"MajorName":"珠宝专业2",
"TeacherId":1,
"MajorId":2
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
#### 专业新增
请求url  http://192.168.0.123:7566/Manage/AddMajor

请求方式  post 

请求参数 
``` json
{
"SchoolId":1,
"CollegeId":1,
"MajorName":"珠宝专业2",
"TeacherId":1
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
#### 专业删除
请求url  http://192.168.0.123:7566/Manage/DelMajor

请求方式  post 

请求参数 
``` json
{
"MajorId":2
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

#### 班级列表 根据专业id查询班级
请求url  http://192.168.0.123:7566/Manage/GetClassByMajorId?MajorId=1

请求方式  get 

请求参数  MajorId=1

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "MajorId": 1,
            "ClassName": "计算机一班",
            "TeacherId": 1
        }
    ],
    "msg": "成功"
}
```



#### 班级列表 根据学校id 查询班级
请求url  http://192.168.0.123:7566/Manage/GetClassBySchoolId?SchoolId=1

请求方式  get 

请求参数  SchoolId=1

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "MajorId": 1,
            "SchoolId": 1,
            "ClassName": "宝石2班",
            "TeacherId": 1,
            "SchoolName": "清华大学",
            "MajorName": "计算机科学与技术",
            "CollegeId": 1,
            "CollegeName": "美术学院"
        }
    ],
    "msg": "成功"
}
```

#### 班级列表 查询所有班级列表
请求url  http://192.168.0.123:7566/Manage/GetClassView 

请求方式  get 

请求参数   

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "MajorId": 1,
            "SchoolId": 1,
            "ClassName": "宝石2班",
            "TeacherId": 1,
            "SchoolName": "清华大学",
            "MajorName": "计算机科学与技术",
            "CollegeId": 1,
            "CollegeName": "美术学院"
        }
    ],
    "msg": "成功"
}
```


#### 查看全部班级及班级下面的学生  根据学校id查询
请求url   http://192.168.0.123:7566/Manage/GetClassStudentsBySchoolId?SchoolId=1

请求方式  get 

请求参数   
SchoolId=1
返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "ClassName": "宝石1班",
            "Students": [
                {
                    "Id": 1,
                    "TrueName": "杜拉拉2"
                },
                {
                    "Id": 3,
                    "TrueName": "杜靠靠"
                }
            ]
        },
        {
            "Id": 3,
            "ClassName": "宝石1班111",
            "Students": [
                {
                    "Id": 4,
                    "TrueName": "杜靠靠1"
                },
                {
                    "Id": 5,
                    "TrueName": "杜靠靠2"
                }
            ]
        }
    ],
    "msg": "成功"
}
```


#### 班级修改
请求url  http://192.168.0.123:7566/Manage/EditClass

请求方式  post 

请求参数 
``` json
{
"MajorId":1,
"ClassName":"宝石1班",
"TeacherId":1,
"SchoolId":1,
"Id":1
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
#### 班级新增
请求url  http://192.168.0.123:7566/Manage/AddClass

请求方式  post 

请求参数 
``` json
{
"MajorId":1,
"ClassName":"宝石1班",
"SchoolId":1,
"TeacherId":1
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
#### 班级删除
请求url  http://192.168.0.123:7566/Manage/DelClass

请求方式  post 

请求参数 
``` json
{
"Id":1
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




#### 站点列表 根据学校id查询站点列表
请求url  http://192.168.0.123:7566/Manage/GetStandListBySchoolId?SchoolId=1

请求方式  get 

请求参数  SchoolId=1

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "StandName": "武汉国土",
            "SchoolId": 1,
            "TeacherId": 1
        }
    ],
    "msg": "成功"
}
```
#### 站点修改
请求url  http://192.168.0.123:7566/Manage/EditStand

请求方式  post 

请求参数 
``` json
{
    "StandName": "武汉国土1",
    "SchoolId": 1,
    "Id": 1
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
#### 站点新增
请求url  http://192.168.0.123:7566/Manage/AddStand

请求方式  post 

请求参数 
``` json
{
    "StandName": "武汉国土1",
    "SchoolId": 1,
    "TeacherId": 1
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
#### 站点删除
请求url  http://192.168.0.123:7566/Manage/DelStand

请求方式  post 

请求参数 
``` json
{
    "Id": 3
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

#### 学生列表
请求url  http://192.168.0.123:7566/Manage/GetStudentViewList

请求方式  get 

请求参数   SchoolId=1

返回数据
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "StudentType": 1,
            "StudentAccount": "dll",
            "StudentPwd": "e10adc3949ba59abbe56e057f20f883e",
            "StandId": 1,
            "ExamName": "考试用户1",
            "SchoolId": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "ClassId": 2,
            "TrueName": "杜拉拉2",
            "IDNumber": "111121212",
            "ExamNumber": "45454545454",
            "Birthday": "1988-02-12",
            "Phone": "15922222222",
            "Email": "284881441@qq.com",
            "IDImage": "1",
            "FaceOpen": 1,
            "SchoolName": "清华大学",
            "CollegeName": "美术学院",
            "MajorName": "计算机科学与技术",
            "ClassName": "宝石1班",
            "StandName": "武汉国土1",
            "Sex": 0
        }
    ],
    "msg": "成功"
}
```
#### 学生新增
请求url  http://192.168.0.123:7566/Manage/AddStudent

请求方式  post 

请求参数 
``` json
{
	"StudentType":1,
	"StudentAccount":"dll",
	"StudentPwd":"123",
	"StandId":1,
	"ExamName":"考试用户1",
	"SchoolId":2,
	"CollegeId":1,
	"MajorId":1,
	"ClassId":1,
	"TrueName":"杜拉拉2",
	"IDNumber":"111121212",
	"ExamNumber":"45454545454",
	"Birthday":"1988-02-12",
	"Phone":"15922222222",
	"Email":"284881441@qq.com",
	"IDImage":"1",
    "Sex":0
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
#### 学生修改

请求url  http://192.168.0.123:7566/Manage/EditStudent

请求方式  post 

请求参数 
``` json
{
	"StudentType":1,
	"StudentAccount":"dll",
	"StudentPwd":"123",
	"StandId":1,
	"ExamName":"考试用户1",
	"SchoolId":2,
	"CollegeId":1,
	"MajorId":1,
	"ClassId":1,
	"TrueName":"杜拉拉2",
	"IDNumber":"111121212",
	"ExamNumber":"45454545454",
	"Birthday":"1988-02-12",
	"Phone":"15922222222",
	"Email":"284881441@qq.com",
	"IDImage":"1",
    "FaceOpen":"1",
    "Sex":0,
    "Id":2
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
#### 学生删除
请求url  http://192.168.0.123:7566/Manage/DelStudent

请求方式  post 

请求参数 
``` json
{
    "Id":2
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
#### 学生证件照上传
请求url  http://192.168.0.124:7566/Manage/MatchAddStudentIDImg

请求方式  post 

请求参数 
``` js
var formData1 = new FormData();
formData1.set("files",file); //file 文件名为idnumber或examnumber  匹配不上则在返回的data里面
```

返回数据

``` json
{
    "code": 1,
    "data": [
        "错误的文件.jpg"
    ],
    "msg": "操作成功,未导入的数据在data"
}
```
#### 修改学生人脸识别标识


请求url  http://192.168.0.123:7566/Manage/EditFaceOpenState

请求方式  post 

请求参数 

``` json
{
"FaceOpen":0,
"Id":1
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
#### 修改学生密码

请求url  http://192.168.0.123:7566/Manage/EditStudentPassWordById

请求方式  post 

请求参数 

``` json
{
    "Id": 1,
    "StudentPwd":"123456"
}
```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "修改密码成功"
}
```
#### 重置密码
请求url  http://192.168.0.123:7566/Manage/ReSetStudentPassWord

请求方式  post 

请求参数 

``` json
{
    "Id": 1
}
```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "重置密码成功" // 默认123456
}
```




### 课程列表 课程与章节管理
#### 找当前班级有哪些课程

#### 当前课程挂载哪些班级
#### 课程图片上传
请求url  http://192.168.0.123:7566/Manage/EditCourseFile

请求方式  post 

请求参数 
``` js
var formData1 = new FormData();
formData1.set("files",file);
formData1.set("data",'{"CourseId":8}');
```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 课程关系编辑
请求url  http://192.168.0.123:7566/Manage/EditClassRelation

请求方式  post 

请求参数 
``` json
{
    "Id": 5,
    "AddClass": [3],
    "AddStudent": [1],
    "RemoveClass": [2],
    "RemoveStudent": [7]
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
#### 课程新增
请求url  http://192.168.0.123:7566/Manage/AddCourse

请求方式  post 

请求参数 
``` json
{
    "CourseName": "test",
    "Digest":"1111",
    "SchoolId": 1,
    "CollegeId": 1,
    "MajorId": 1,
    "TeacherId": 1,
    "CourseStartTime": "2023-06-21 16:07:00",
    "CourseEndTime": "2023-06-25 16:07:00",
    "Status": 1
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
#### 课程修改
请求url  http://192.168.0.123:7566/Manage/EditCourse

请求方式  post 

请求参数 
``` json
{
    "CourseName": "test",
    "Digest":"1111",
    "SchoolId": 1,
    "CollegeId": 1,
    "MajorId": 1,
    "TeacherId": 1,
    "CourseStartTime": "2023-06-21 16:07:00",
    "CourseEndTime": "2023-06-25 16:07:00",
    "Status": 1,
    "Id":8
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
#### 课程删除
请求url  http://192.168.0.123:7566/Manage/DelCourse

请求方式  post 

请求参数 
``` json
{
    "CourseId":8
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

#### 课程废弃
请求url  http://192.168.0.123:7566/Manage/CourseCancel

请求方式  post 

请求参数 
``` json
{
    "CourseId":8
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

#### 获取单个课程的班级关系 ，和学生进度表
请求url  http://192.168.0.123:7566/Manage/GetCourseClassRelationById?CourseId=8

请求方式  get 

请求参数 
CourseId=8

返回数据

``` json
{
    "code": 1,
    "data": {
        "ClassArr": [
            2
        ],
        "StudentArr": [
            [
                3,
                4
            ],
            [
                3,
                5
            ]
        ]
    },
    "msg": "操作成功"
}
```







####   查看当前课程内容

请求url  http://192.168.0.123:7566/Manage/GetCourseByCourseId?CourseId=8

请求方式  get 

请求参数 
CourseId=8

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 8,
        "CourseName": "test",
        "Digest":"1111",
        "SchoolId": 1,
        "CollegeId": 1,
        "MajorId": 1,
        "FileId": 3,
        "TeacherId": 1,
        "CourseStudentJson": "",
        "CourseStartTime": "2023-06-21 16:07:00",
        "CourseEndTime": "2023-06-25 16:07:00",
        "Status": 1,
        "SchoolName": "清华大学",
        "CollegeName": "美术学院",
        "MajorName": "计算机科学与技术",
        "FilePath": "Resources/Img/bf7c4f089fa96235dbd8d06d516372a3.png",
        "TeacherName": ""
    },
    "msg": "操作成功"
}
```
 
####   看所有课程 带老师Id

请求url  http://192.168.0.123:7566/Manage/GetCourseListByTeacherId?TeacherId=1

请求方式  get 

请求参数 
TeacherId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "Id": 8,
            "CourseName": "test",
            "Digest":"1111",
            "SchoolId": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "FileId": 3,
            "TeacherId": 1,
            "CourseStudentJson": "",
            "CourseStartTime": "2023-06-21 16:07:00",
            "CourseEndTime": "2023-06-25 16:07:00",
            "Status": 1,
            "SchoolName": "清华大学",
            "CollegeName": "美术学院",
            "MajorName": "计算机科学与技术",
            "FilePath": "Resources/Img/bf7c4f089fa96235dbd8d06d516372a3.png",
            "TeacherName": ""
        }
    ],
    "msg": "成功"
}
```


#### 章节列表  根据课程查询

请求url  http://192.168.0.123:7566/Manage/GetChapterByCourseId?CourseId=25

请求方式  get 

请求参数 
CourseId=25

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "ChapterName": "test",
            "CourseId": 25,
            "ChapterType": "图文课1",
            "ChapterOrder": 0,
            "ChapterContent": ""
        }
    ],
    "msg": "操作成功"
}
```


#### 查询单个章节
请求url  http://192.168.0.123:7566/Manage/GetChapterById?ChapterId=2

请求方式  get 

请求参数 
ChapterId=2

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 2,
        "ChapterName": "test",
        "CourseId": 25,
        "ChapterType": "图文课1",
        "ChapterOrder": 0,
        "ChapterContent": "aaas111ss",
        "FileInfo": [
            {
                "Id": 13,
                "FileType": ".jpg",
                "FileName": "",
                "FileUseTo": "",
                "SchoolId": 0,
                "FilePath": "Resources/Img/35c9580b804a56aa9cb46c332f7b4cd5.jpg"
            }
        ]
    },
    "msg": "操作成功"
}
```

#### 新增章节

请求url  http://192.168.0.123:7566/Manage/AddChapter

请求方式  post 

请求参数 
``` js
 
var formData1 = new FormData();
formData1.set("files",file);
formData1.set("data",'{
    "ChapterName":"test",
    "CourseId": 25,
    "ChapterType":"图文课",
    "ChapterContent":"aaasss"
}');

```
返回数据

``` json
{
    "code": 1,
    "data": "{'ChapterId':1}",
    "msg": "操作成功"
}
```
#### 修改章节

请求url http://192.168.0.123:7566/Manage/EditChapter

请求方式  post 

请求参数 
``` js
 
var formData1 = new FormData();
formData1.set("files",file);// 可多文件
formData1.set("data",'{
    "ChapterName": "test",
    "CourseId": 25,
    "ChapterType": "图文课1",
    "ChapterContent": "aaas111ss",
    "RemoveFile": [
       
    ],
    "Id": 2
}');

```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除章节
请求url http://192.168.0.123:7566/Manage/DelChapter

请求方式  post 

请求参数 
``` json
{
    "ChapterId": 2
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

#### 上传章节内容文件
请求url http://192.168.0.123:7566/Manage/UploadChapterFile

请求方式  post 

请求参数 
``` js
 
var formData1 = new FormData();
formData1.set("files",file);// 可多文件
formData1.set("data",'{
    "SchoolId": 8
}');

```

返回数据

``` json
{
    "code": 1,
    "data": "文件路径",
    "msg": "操作成功"
}
```


#### 章节顺序修改
请求url http://192.168.0.123:7566/Manage/UpdateChapterOrder

请求方式  post 

请求参数 
``` json
[
    {
        "ChapterOrder": [
            3,
            1
        ]//索引0 是章节id  1是顺序
    }
]


```

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```

#### 上传视频课接口

请求url http://192.168.0.123:7566/Manage/UploadChapterVideoFile

请求方式  post 

请求参数 
``` js
 
var formData1 = new FormData();
formData1.set("files",file);// 可多文件
formData1.set("data",'{
    "SchoolId": 8,
     "ChapterId": 1,
}');

```

返回数据

``` json
{
    "code": 1,
    "data": "文件路径",
    "msg": "{}"
}
```


#### 章节预览


### 题库与试卷管理
#### 题库列表（查询当前学校的题目）（管理员查询所有学校  ，可以以学校分类）
#### 新增题目

请求url  http://192.168.0.123:7566/Manage/AddQuestion

请求方式  post 

请求参数 
``` js
 
var formData1 = new FormData();
formData1.set("files",file);
formData1.set("data",'{
    "SchoolId": 1,
    "QuestionPoolId": 1,
    "QuestionName": "实操题",
    "QuestionType": 1,
    "QuestionContent": "",
    "Digree": 1,
    "MajorID": 1,
    "CollegeId": 1,
    "CourseId": 1,
    "Answer": "1",
    "TeacherId": 1
}');

```
返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 修改题目

请求url  http://192.168.0.123:7566/Manage/EditQuestion

请求方式  post 

请求参数 
``` js
 
var formData1 = new FormData();
formData1.set("files",file);
formData1.set("data",'{
    "SchoolId": 1,
    "QuestionPoolId": 1,
    "QuestionName": "实操题",
    "QuestionType": 5,
    "QuestionContent": "",
    "Digree": 1,
    "MajorID": 1,
    "CollegeId": 1,
    "CourseId": 1,
    "Answer": "1",
    "QuestionId": "1",
    "RemoveFile":[]
}');

```
返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除题目
请求url http://192.168.0.123:7566/Manage/DelQuestion

请求方式  post 

请求参数 
``` json
{
    "QuestionId": 1
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

#### 根据学校查询题目
请求url  http://192.168.0.123:7566/Manage/GetQuestionBySchoolId?SchoolId=1

请求方式  get 

请求参数 
SchoolId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "QuestionId": 2,
            "SchoolId": 1,
            "QuestionPoolId": 1,
            "QuestionName": "实操题",
            "QuestionType": 1,
            "QuestionContent": "",
            "Digree": 1,
            "MajorID": 1,
            "CollegeId": 1,
            "CourseId": 1,
            "Answer": "1",
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": ""
        },
        {
            "QuestionId": 3,
            "SchoolId": 1,
            "QuestionPoolId": 1,
            "QuestionName": "实操题",
            "QuestionType": 5,
            "QuestionContent": "183",
            "Digree": 1,
            "MajorID": 1,
            "CollegeId": 1,
            "CourseId": 1,
            "Answer": "1",
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": ""
        }
    ],
    "msg": "操作成功"
}
```



####  社会人士查询练习题  
请求url  http://192.168.0.124:7566/Student/GetSocietyQuestionByStudentId?StudentId=6

请求方式  get 

请求参数 
StudentId=6

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "QuestionId": 2,
            "SchoolId": 1,
            "QuestionPoolId": 1,
            "QuestionName": "单选题",
            "QuestionType": 1,
            "QuestionContent": "1111",
            "Digree": 1,
            "MajorID": 1,
            "CollegeId": 1,
            "CourseId": 25,
            "Answer": "1",
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": "课程3"
        },
        {
            "QuestionId": 6,
            "SchoolId": 1,
            "QuestionPoolId": 1,
            "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
            "QuestionType": 1,
            "QuestionContent": "[\"A. 格斯拉兄弟将自己的店开在一条横街上，这条横街座落在伦敦市的西区。\",\"B. 我不断地叩问历史的发展，产生了从中找出一条贯穿其中的脉络的愿望。\",\"C. 老王为新春写的这幅对联，寄托了美好的祝福，展现了书法大境界。\"]",
            "Digree": 4,
            "MajorID": 0,
            "CollegeId": 0,
            "CourseId": 25,
            "Answer": "1",
            "MajorName": "",
            "CollegeName": "",
            "CourseName": "课程3"
        }
    ],
    "msg": "操作成功"
}
```


#### 根据题目id查询题目
请求url  http://192.168.0.123:7566/Manage/GetQuestionByQuestionId?QuestionId=2

请求方式  get 

请求参数 
QuestionId=2

返回数据

``` json
{
    "code": 1,
    "data": {
        "QuestionId": 3,
        "SchoolId": 1,
        "QuestionPoolId": 1,
        "QuestionName": "实操题",
        "QuestionType": 5,  //如果是5 就是实操题  fileinfo 就是实操的文件  如果是其他   里面就是附件列表
        "QuestionContent": "183",
        "Digree": 1,
        "MajorID": 1,
        "CollegeId": 1,
        "CourseId": 1,
        "Answer": "1",
        "MajorName": "计算机科学与技术",
        "CollegeName": "美术学院",
        "CourseName": "",
        "FileInfo": [
            {
                "Id": 183,
                "FileType": ".zip",
                "FileName": "api.zip",
                "FileUseTo": "",
                "SchoolId": 0,
                "FilePath": "Resources/Zip/4296e39ef8390472aa3565649d79614e.zip"
            }
        ]
    },
    "msg": "操作成功"
}
```


#### 导入题目预览
请求url  http://192.168.0.124:7566/Manage/GetQuestionExeclData

请求方式  post 

请求参数 
 formdata  files file

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试1",
            "QuestionContent": [
                "1",
                "2",
                "4"
            ],
            "QuestionType": 1,
            "Digree": 2,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0    
        
        },
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试2",
            "QuestionContent": [
                "1",
                "3",
                "4",
                "5"
            ],
            "QuestionType": 2,
            "Digree": 2,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0    
        },
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试3",
            "QuestionContent": [
                "12",
                "3435",
                "23",
                "3"
            ],
            "QuestionType": 1,
            "Digree": 3,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0    
        }
    ],
    "msg": "操作成功"
}
```

#### 批量导入题目
请求url  http://192.168.0.124:7566/Manage/MatchAddQuestion

请求方式  post 

请求参数 
 ``` json
 [
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试1",
            "QuestionContent": [
                "1",
                "2",
                "4"
            ],
            "QuestionType": 1,
            "Digree": 2,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0    
        
        },
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试2",
            "QuestionContent": [
                "1",
                "3",
                "4",
                "5"
            ],
            "QuestionType": 2,
            "Digree": 2,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0    
        },
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试3",
            "QuestionContent": [
                "12",
                "3435",
                "23",
                "3"
            ],
            "QuestionType": 1,
            "Digree": 3,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0    
        }
    ]
 ```

返回数据

 
``` json
{
    "code": 1,
    "data": [
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试1",
            "QuestionContent": [
                "1",
                "2",
                "4"
            ],
            "QuestionType": 1,
            "Digree": 2,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":0     //0 标识未导入  1 标识表示导入成功
        
        },
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试2",
            "QuestionContent": [
                "1",
                "3",
                "4",
                "5"
            ],
            "QuestionType": 2,
            "Digree": 2,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":1  
        },
        {
            "QuestionPoolId": 1,
            "QuestionName": "测试3",
            "QuestionContent": [
                "12",
                "3435",
                "23",
                "3"
            ],
            "QuestionType": 1,
            "Digree": 3,
            "Answer": "[1,2]",
            "CourseId":1,  
            "TeacherId":1,
            "Status":1     
        }
    ],
    "msg": "操作成功"
}
```
 
#### 学生提交练习记录


请求url  http://192.168.0.124:7566/Student/AddOperPractice

请求方式  post 

请求参数 
``` json
 {
    "QuestionId": 1,
    "StudentId": 1,
    "PracticeStep": "",
    "PracticeAnswer": "",
    "PracticeScore": 1.0
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
#### 管理端查看练习记录

请求url  http://192.168.0.124:7566/Manage/GetOperPracticeByQuestionId?QuestionId=1

请求方式  get 

请求参数 
QuestionId=1
 
返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "StudentId": 1,
            "TrueName": "杜拉拉2",
            "StudentNum": 2,
            "MaxTime": "2018-05-07 11:00:00"
        },
        {
            "StudentId": 3,
            "TrueName": "杜靠靠",
            "StudentNum": 1,
            "MaxTime": "2018-05-07 11:00:00"
        }
    ],
    "msg": "操作成功"
}
```

#### 试卷列表  根据学校id查询试卷
http://192.168.0.123:7566/Manage/GetTestPaperBySchoolId?SchoolId=1


请求方式  get 

请求参数 
SchoolId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "TestPaperName": "测试试卷1",
            "ExamDuration": 60,
            "FullMarks": 100,
            "PassScore": 60,
            "TestPaperType": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "CourseId": 1,
            "TeacherId": 1,
            "SchoolId": 1,
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": "",
        },
        {
            "Id": 2,
            "TestPaperName": "测试试卷1",
            "ExamDuration": 60,
            "FullMarks": 100,
            "PassScore": 60,
            "TestPaperType": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "CourseId": 1,
            "TeacherId": 1,
            "SchoolId": 1,
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": "",
        },
        {
            "Id": 3,
            "TestPaperName": "测试试卷3",
            "ExamDuration": 60,
            "FullMarks": 100,
            "PassScore": 60,
            "TestPaperType": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "CourseId": 1,
            "TeacherId": 1,
            "SchoolId": 1,
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": "",
        }
    ],
    "msg": "操作成功"
}
```
#### 查看已删除试卷列表

http://192.168.0.123:7566/Manage/GetDelTestPaperByTeacherId?TeacherId=1


请求方式  get 

请求参数 
TeacherId=1

如果TeacherId==0 那就需要传SchoolId  如果不等于0  则不需要传学校id
SchoolId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "TestPaperName": "测试试卷1",
            "ExamDuration": 60,
            "FullMarks": 100,
            "PassScore": 60,
            "TestPaperType": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "CourseId": 0,
            "TeacherId": 1,
            "SchoolId": 1,
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": "",
        },
        {
            "Id": 3,
            "TestPaperName": "测试试卷3",
            "ExamDuration": 60,
            "FullMarks": 100,
            "PassScore": 60,
            "TestPaperType": 1,
            "CollegeId": 1,
            "MajorId": 1,
            "CourseId": 0,
            "TeacherId": 1,
            "SchoolId": 1,
            "MajorName": "计算机科学与技术",
            "CollegeName": "美术学院",
            "CourseName": "",
        }
    ],
    "msg": "操作成功"
}
```
#### 查看试卷（试卷预览）获取该试卷所有题目
//根据试卷id查看试卷

http://192.168.0.123:7566/Manage/GetTestPaperByTestPaperId?TestPaperId=2


请求方式  get 

请求参数 
TestPaperId=2

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 2,
        "TestPaperName": "测试试卷1",
        "ExamDuration": 60,
        "FullMarks": 100,
        "PassScore": 60,
        "TestPaperType": 1,
        "CollegeId": 1,
        "MajorId": 1,
        "CourseId": 0,
        "TeacherId": 1,
        "SchoolId": 1,
        "MajorName": "计算机科学与技术",
        "CollegeName": "美术学院",
        "CourseName": "",
        "SchoolName": "清华大学",
        "TeacherName": "tea",
        "TestPaperQuestionTypeOver": [
            {
                "QuestionType": 1,
                "QuestionIdNum": 1,
                "QuestionScore": 50
            },
            {
                "QuestionType": 2,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 3,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 4,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            },
            {
                "QuestionType": 5,
                "QuestionIdNum": 0,
                "QuestionScore": 0
            }
        ],
        "TestPaperQuestionViewFile": [
            {
                "QuestionId": 3,
                "QuestionName": "",
                "QuestionPoolId": 0,
                "QuestionType": 0,
                "QuestionContent": "",
                "QuestionScore": 50,
                "Digree": 0,
                "Answer": "0",
                "FileInfo": []
            },
            {
                "QuestionId": 2,
                "QuestionName": "单选题",
                "QuestionPoolId": 1,
                "QuestionType": 1,
                "QuestionContent": "1111",
                "QuestionScore": 50,
                "Digree": 1,
                "Answer": "1",
                "FileInfo": [
                    {
                        "Id": 182,
                        "FileType": ".zip",
                        "FileName": "api.zip",
                        "FileUseTo": "",
                        "SchoolId": 0,
                        "FilePath": "Resources/Annex/4809a7968dfceb121618db6b7cd80786.zip"
                    }
                ]
            }
        ]
    },
    "msg": "操作成功"
}
```


#### 根据id查询试卷 修改时候用
//根据试卷id查看试卷

http://192.168.0.123:7566/Manage/GetTestPaperByPaperId?TestPaperId=2


请求方式  get 

请求参数 
TestPaperId=2

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 8,
        "TestPaperName": "试卷测试1",
        "ExamDuration": 100,
        "FullMarks": 100,
        "PassScore": 60,
        "TestPaperType": 1,
        "CollegeId": 0,
        "MajorId": 0,
        "CourseId": 0,
        "QuestionScoreJson": [
            {
                "QuestionType": 1,
                "QuestionIdArr": [
                    {
                        "QuestionId": 2,
                        "QuestionName": "单选题",
                        "Score": 10
                    },
                    {
                        "QuestionId": 6,
                        "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
                        "Score": 10
                    }
                ],
                "FullMarksRatio": 20,
                "QuestionIdNum": 2
            },
            {
                "QuestionType": 2,
                "QuestionIdArr": [
                    {
                        "QuestionId": 16,
                        "QuestionName": "多选题测试",
                        "Score": 20
                    }
                ],
                "FullMarksRatio": 20,
                "QuestionIdNum": 1
            },
            {
                "QuestionType": 3,
                "QuestionIdArr": [
                    {
                        "QuestionId": 17,
                        "QuestionName": "判断题测试",
                        "Score": 20
                    }
                ],
                "FullMarksRatio": 20,
                "QuestionIdNum": 1
            },
            {
                "QuestionType": 4,
                "QuestionIdArr": [
                    {
                        "QuestionId": 18,
                        "QuestionName": "填空题测试",
                        "Score": 20
                    }
                ],
                "FullMarksRatio": 20,
                "QuestionIdNum": 1
            },
            {
                "QuestionType": 5,
                "QuestionIdArr": [],
                "FullMarksRatio": 20,
                "QuestionIdNum": 0
            }
        ],
        "TeacherId": 1,
        "SchoolId": 1,
    },
    "msg": "操作成功"
}
```

#### 添加试卷


请求url  http://192.168.0.123:7566/Manage/AddTestPaper

请求方式  post 

请求参数 
``` json
 {
    "TestPaperName": "测试试卷3",
    "ExamDuration": 60,
    "FullMarks": 100.00,
    "PassScore": 60.00,
    "TestPaperType": 1,
    "CollegeId": 1,
    "MajorId": 1,
    "CourseId": 1,
    "TeacherId": 1,
    "SchoolId": 1,
    "QuestionScoreJson": [
        {
            "QuestionType": 1,
            "FullMarksRatio": 70,
            "QuestionArr": [
                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}
            ]
        },
        {
            "QuestionType": 2,
            "FullMarksRatio": 0,
            "QuestionArr": [
                                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}
            ]
        },
        {
            "QuestionType": 3,
            "FullMarksRatio": 0,
            "QuestionArr": [                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}]
        },
        {
            "QuestionType": 4,
            "FullMarksRatio": 0,
            "QuestionArr": [                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}]
        },
        {
            "QuestionType": 5,
            "FullMarksRatio": 30,
            "QuestionArr": [
                            {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}
            ]
        }
    ]
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
#### 修改试卷


请求url  http://192.168.0.123:7566/Manage/EditTestPaper

请求方式  post 

请求参数 
``` json
 {
    "Id":1,
    "TestPaperName": "测试试卷3",
    "ExamDuration": 60,
    "FullMarks": 100.00,
    "PassScore": 60.00,
    "TestPaperType": 1,
    "CollegeId": 1,
    "MajorId": 1,
    "CourseId": 1,
    "TeacherId": 1,
    "SchoolId": 1,
    "QuestionScoreJson": [
        {
            "QuestionType": 1,
            "FullMarksRatio": 70,
            "QuestionArr": [
                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}
            ]
        },
        {
            "QuestionType": 2,
            "FullMarksRatio": 0,
            "QuestionArr": [
                                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}
            ]
        },
        {
            "QuestionType": 3,
            "FullMarksRatio": 0,
            "QuestionArr": [                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}]
        },
        {
            "QuestionType": 4,
            "FullMarksRatio": 0,
            "QuestionArr": [                {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}]
        },
        {
            "QuestionType": 5,
            "FullMarksRatio": 30,
            "QuestionArr": [
                            {"QuestionId":7,"QuestionName":"实操题测试","Score":10},{"QuestionId":7,"QuestionName":"实操题测试","Score":10}
            ]
        }
    ]
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
#### 删除试卷


请求url  http://192.168.0.123:7566/Manage/DelTestPaper

请求方式  post 

请求参数 
``` json
 {
    "Id":1
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

### 考试与成绩管理
#### 考试与考场列表

#### 将execl批量添加学生到考试中
请求url  http://192.168.0.124:7566/Manage/AddExamBatchSessionStudent

请求方式  post 

请求参数 
 formdata  files file

 目前只支持xlsx

返回数据

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}


//err
{
    "code": 0,
    "data": "{}",
    "msg":  "操作失败,该execl 第1行， 课程代码找不到考试场次",
}


```

#### 新增考试主题

请求url  http://192.168.0.123:7566/Manage/AddExam

请求方式  post 

请求参数 
``` json
 
 {
    "SchoolId": 1,
    "ExamName": "测试试卷3",
    "ExamDescribe": "试卷描述2",
    "ExamStatus": 1,
    "FaceVerify": 1,
    "TeacherId": 1
}
```
返回数据

``` json
{
    "code": 1,
    "data": "{\"ExamId\":3}",
    "msg": "操作成功"
}
```
#### 删除考试主题
请求url  http://192.168.0.123:7566/Manage/DelExam 

请求方式  post 

请求参数 
``` json
{
    "ExamId": 1
}
```
返回数据

``` json
{
    "code": 0,
    "data": "{}",
    "msg": "考试已归档不能删除"
}
```

####  废弃考试
请求url  http://192.168.0.123:7566/Manage/ExamCancel

请求方式  post 

请求参数 
``` json
{
    "ExamId": 1
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


#### 编辑考试主题
请求url  http://192.168.0.123:7566/Manage/EditExam

请求方式  post 

请求参数 
``` json
 
{
    "SchoolId": 1,
    "ExamName": "测试试卷35",
    "ExamDescribe": "试卷描述2",
    "ExamStatus": 1,
    "FaceVerify": 1,
    "Id":1
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
#### 新增考试场次

请求url  http://192.168.0.123:7566/Manage/AddExamSession

请求方式  post 

请求参数 
``` json
 
{
    "ExamId": 1,
    "StartTime": "2021-2-3",
    "EndTime": "2023-2-3",
    "TestPaperId": 2
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

#### 编辑考试场次
请求url  http://192.168.0.123:7566/Manage/EditExamSession

请求方式  post 

请求参数 
``` json
{
    "ExamId": 1,
    "StartTime": "测试试卷35",
    "EndTime": "试卷描述2",
    "TestPaperId": 3,
    "Id": 1
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


#### 删除考试场次

请求url  http://192.168.0.123:7566/Manage/DelExamSession

请求方式  post 

请求参数 
``` json
{
    "Id": 1,
    "ExamId":1
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



#### 新增考试学生
请求url  http://192.168.0.123:7566/Manage/AddExamStudent

请求方式  post 

请求参数 
``` json
 
{
    "ExamId": 2,
    "AddStudentIdArr": [8,9,10],
    "RemoveStudentIdArr": []
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

#### 修改考试学生
请求url  http://192.168.0.123:7566/Manage/EditExamStudent

请求方式  post 

请求参数 
``` json
 
{
    "ExamId": 1,
    "AddStudentIdArr": [8,9,10],
    "RemoveStudentIdArr": [7]
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



#### 查看考试信息列表 根据学校id查询  考试主题列表
请求url  http://192.168.0.123:7566/Manage/GetExamBySchoolId?SchoolId=1

请求方式  get 

请求参数 
SchoolId=1
返回数据

##### 备注
    状态（state） 未开始  已开始  已完成
                     0	    1	2
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "SchoolId": 1,
            "ExamName": "测试试卷35",
            "ExamDescribe": "灌灌灌灌灌",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 2,
            "ReviewFlag":1,
            "ExamSessionArr": [
                {
                    "Id": 1,
                    "ExamId": 1,
                    "StartTime": "2023-8-9 9:43:16",
                    "EndTime": "2023-8-11 9:43:23",
                    "TestPaperId": 5,
                    "TestPaperName": "试卷测试1",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                },
                {
                    "Id": 2,
                    "ExamId": 1,
                    "StartTime": "2023-8-5 0:0:0",
                    "EndTime": "2023-8-6 0:0:0",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                },
                {
                    "Id": 12,
                    "ExamId": 1,
                    "StartTime": "2023-8-8 9:54:46",
                    "EndTime": "2023-8-9 9:54:54",
                    "TestPaperId": 5,
                    "TestPaperName": "试卷测试1",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": [
                {
                    "StudentId": 6,
                    "TrueName": "五五"
                },
                {
                    "StudentId": 8,
                    "TrueName": "轻轻巧巧"
                },
                {
                    "StudentId": 9,
                    "TrueName": "学1"
                },
                {
                    "StudentId": 10,
                    "TrueName": "学2"
                }
            ]
        },
        {
            "Id": 2,
            "SchoolId": 1,
            "ExamName": "测试试卷2",
            "ExamDescribe": "试卷描述2",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": [
                {
                    "StudentId": 8,
                    "TrueName": "轻轻巧巧"
                },
                {
                    "StudentId": 9,
                    "TrueName": "学1"
                },
                {
                    "StudentId": 10,
                    "TrueName": "学2"
                }
            ]
        },
        {
            "Id": 3,
            "SchoolId": 1,
            "ExamName": "测试试卷3",
            "ExamDescribe": "试卷描述2",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 4,
            "SchoolId": 1,
            "ExamName": "开学考试",
            "ExamDescribe": "考试测试",
            "ExamStatus": 0,
            "FaceVerify": 0,
            "State": 2,
            "ExamSessionArr": [
                {
                    "Id": 19,
                    "ExamId": 4,
                    "StartTime": "2023-8-7 17:10:45",
                    "EndTime": "2023-8-8 17:10:47",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": [
                {
                    "StudentId": 8,
                    "TrueName": "轻轻巧巧"
                },
                {
                    "StudentId": 15,
                    "TrueName": "学7"
                },
                {
                    "StudentId": 11,
                    "TrueName": "学3"
                },
                {
                    "StudentId": 12,
                    "TrueName": "学4"
                },
                {
                    "StudentId": 13,
                    "TrueName": "学5"
                },
                {
                    "StudentId": 14,
                    "TrueName": "学6"
                }
            ]
        },
        {
            "Id": 5,
            "SchoolId": 1,
            "ExamName": "开学考试测试1",
            "ExamDescribe": "考试测试1",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 6,
            "SchoolId": 1,
            "ExamName": "开学考试测试2",
            "ExamDescribe": "",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 7,
            "SchoolId": 1,
            "ExamName": "开学考试测试2",
            "ExamDescribe": "",
            "ExamStatus": 0,
            "FaceVerify": 0,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 8,
            "SchoolId": 1,
            "ExamName": "开学考试测试3",
            "ExamDescribe": "",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 9,
            "SchoolId": 1,
            "ExamName": "开学考试测试4",
            "ExamDescribe": "sdfafasdf",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 10,
            "SchoolId": 1,
            "ExamName": "开学考试测试6",
            "ExamDescribe": "fvdstreyrefcx6yujhtrew",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 11,
            "SchoolId": 1,
            "ExamName": "开学考试测试7",
            "ExamDescribe": "fvdstreyrefcx6yujhtrew",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 12,
            "SchoolId": 1,
            "ExamName": "开学考试测试8",
            "ExamDescribe": "sad阿达撒大大驱蚊器我认为",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 13,
            "SchoolId": 1,
            "ExamName": "开学考试测试8",
            "ExamDescribe": "手动阀打发奥迪啊发顺丰撒村是在是防守打法三十",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 14,
            "SchoolId": 1,
            "ExamName": "开学考试测试9",
            "ExamDescribe": "肯德基发电房接口科技反馈哈科技花开减肥的had设计费华快递费",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 15,
            "SchoolId": 1,
            "ExamName": "开学考试测试10",
            "ExamDescribe": "开学考试测试10开学考试测试10开学考试测试10开学考试测试10开学考试测试10开学考试测试10",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 16,
            "SchoolId": 1,
            "ExamName": "开学考试测试11",
            "ExamDescribe": "开学考试测试11开学考试测试11开学考试测试11开学考试测试11开学考试测试11",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 17,
            "SchoolId": 1,
            "ExamName": "开学考试测试12",
            "ExamDescribe": "开学考试测试12开学考试测试12开学考试测试12开学考试测试12开学考试测试12开学考试测试12",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 18,
            "SchoolId": 1,
            "ExamName": "开学考试",
            "ExamDescribe": "森岛帆高冻干粉",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 2,
            "ExamSessionArr": [
                {
                    "Id": 6,
                    "ExamId": 18,
                    "StartTime": "2023-8-7 9:41:7",
                    "EndTime": "2023-8-9 9:41:9",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": []
        },
        {
            "Id": 19,
            "SchoolId": 1,
            "ExamName": "bug测试考试",
            "ExamDescribe": "测试测试",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 2,
            "ExamSessionArr": [
                {
                    "Id": 20,
                    "ExamId": 19,
                    "StartTime": "2023-8-9 14:56:9",
                    "EndTime": "2023-8-11 14:56:13",
                    "TestPaperId": 2,
                    "TestPaperName": "测试试卷1",
                    "ExamDuration": 60,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                },
                {
                    "Id": 21,
                    "ExamId": 19,
                    "StartTime": "2023-8-17 14:56:46",
                    "EndTime": "2023-8-19 14:56:48",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": [
                {
                    "StudentId": 32,
                    "TrueName": "学24"
                }
            ]
        }
    ],
    "msg": "操作成功"
}
```



#### 查看未审核的考试
请求url  http://192.168.0.123:7566/Manage/GetExamReViewBySchoolId?SchoolId=1

请求方式  get 

请求参数 
SchoolId=1
返回数据

##### 备注
    状态（state） 未开始  已开始  已完成
                     0	    1	2
``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "SchoolId": 1,
            "ExamName": "测试试卷35",
            "ExamDescribe": "灌灌灌灌灌",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 2,
            "ReviewFlag":1,
            "ExamSessionArr": [
                {
                    "Id": 1,
                    "ExamId": 1,
                    "StartTime": "2023-8-9 9:43:16",
                    "EndTime": "2023-8-11 9:43:23",
                    "TestPaperId": 5,
                    "TestPaperName": "试卷测试1",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                },
                {
                    "Id": 2,
                    "ExamId": 1,
                    "StartTime": "2023-8-5 0:0:0",
                    "EndTime": "2023-8-6 0:0:0",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                },
                {
                    "Id": 12,
                    "ExamId": 1,
                    "StartTime": "2023-8-8 9:54:46",
                    "EndTime": "2023-8-9 9:54:54",
                    "TestPaperId": 5,
                    "TestPaperName": "试卷测试1",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": [
                {
                    "StudentId": 6,
                    "TrueName": "五五"
                },
                {
                    "StudentId": 8,
                    "TrueName": "轻轻巧巧"
                },
                {
                    "StudentId": 9,
                    "TrueName": "学1"
                },
                {
                    "StudentId": 10,
                    "TrueName": "学2"
                }
            ]
        },
        {
            "Id": 2,
            "SchoolId": 1,
            "ExamName": "测试试卷2",
            "ExamDescribe": "试卷描述2",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
              "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": [
                {
                    "StudentId": 8,
                    "TrueName": "轻轻巧巧"
                },
                {
                    "StudentId": 9,
                    "TrueName": "学1"
                },
                {
                    "StudentId": 10,
                    "TrueName": "学2"
                }
            ]
        },
        {
            "Id": 3,
            "SchoolId": 1,
            "ExamName": "测试试卷3",
            "ExamDescribe": "试卷描述2",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 4,
            "SchoolId": 1,
            "ExamName": "开学考试",
            "ExamDescribe": "考试测试",
            "ExamStatus": 0,
            "FaceVerify": 0,
            "State": 2,
            "ReviewFlag":1,
            "ExamSessionArr": [
                {
                    "Id": 19,
                    "ExamId": 4,
                    "StartTime": "2023-8-7 17:10:45",
                    "EndTime": "2023-8-8 17:10:47",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": [
                {
                    "StudentId": 8,
                    "TrueName": "轻轻巧巧"
                },
                {
                    "StudentId": 15,
                    "TrueName": "学7"
                },
                {
                    "StudentId": 11,
                    "TrueName": "学3"
                },
                {
                    "StudentId": 12,
                    "TrueName": "学4"
                },
                {
                    "StudentId": 13,
                    "TrueName": "学5"
                },
                {
                    "StudentId": 14,
                    "TrueName": "学6"
                }
            ]
        },
        {
            "Id": 5,
            "SchoolId": 1,
            "ExamName": "开学考试测试1",
            "ExamDescribe": "考试测试1",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 6,
            "SchoolId": 1,
            "ExamName": "开学考试测试2",
            "ExamDescribe": "",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 7,
            "SchoolId": 1,
            "ExamName": "开学考试测试2",
            "ExamDescribe": "",
            "ExamStatus": 0,
            "FaceVerify": 0,
            "State": 0,
            "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 8,
            "SchoolId": 1,
            "ExamName": "开学考试测试3",
            "ExamDescribe": "",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 9,
            "SchoolId": 1,
            "ExamName": "开学考试测试4",
            "ExamDescribe": "sdfafasdf",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 10,
            "SchoolId": 1,
            "ExamName": "开学考试测试6",
            "ExamDescribe": "fvdstreyrefcx6yujhtrew",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 11,
            "SchoolId": 1,
            "ExamName": "开学考试测试7",
            "ExamDescribe": "fvdstreyrefcx6yujhtrew",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 12,
            "SchoolId": 1,
            "ExamName": "开学考试测试8",
            "ExamDescribe": "sad阿达撒大大驱蚊器我认为",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 13,
            "SchoolId": 1,
            "ExamName": "开学考试测试8",
            "ExamDescribe": "手动阀打发奥迪啊发顺丰撒村是在是防守打法三十",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 14,
            "SchoolId": 1,
            "ExamName": "开学考试测试9",
            "ExamDescribe": "肯德基发电房接口科技反馈哈科技花开减肥的had设计费华快递费",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 15,
            "SchoolId": 1,
            "ExamName": "开学考试测试10",
            "ExamDescribe": "开学考试测试10开学考试测试10开学考试测试10开学考试测试10开学考试测试10开学考试测试10",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 16,
            "SchoolId": 1,
            "ExamName": "开学考试测试11",
            "ExamDescribe": "开学考试测试11开学考试测试11开学考试测试11开学考试测试11开学考试测试11",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
             "ReviewFlag":1,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 17,
            "SchoolId": 1,
            "ExamName": "开学考试测试12",
            "ExamDescribe": "开学考试测试12开学考试测试12开学考试测试12开学考试测试12开学考试测试12开学考试测试12",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 0,
            "ExamSessionArr": [],
            "ExamStudentArr": []
        },
        {
            "Id": 18,
            "SchoolId": 1,
            "ExamName": "开学考试",
            "ExamDescribe": "森岛帆高冻干粉",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 2,
            "ExamSessionArr": [
                {
                    "Id": 6,
                    "ExamId": 18,
                    "StartTime": "2023-8-7 9:41:7",
                    "EndTime": "2023-8-9 9:41:9",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": []
        },
        {
            "Id": 19,
            "SchoolId": 1,
            "ExamName": "bug测试考试",
            "ExamDescribe": "测试测试",
            "ExamStatus": 1,
            "FaceVerify": 1,
            "State": 2,
            "ExamSessionArr": [
                {
                    "Id": 20,
                    "ExamId": 19,
                    "StartTime": "2023-8-9 14:56:9",
                    "EndTime": "2023-8-11 14:56:13",
                    "TestPaperId": 2,
                    "TestPaperName": "测试试卷1",
                    "ExamDuration": 60,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                },
                {
                    "Id": 21,
                    "ExamId": 19,
                    "StartTime": "2023-8-17 14:56:46",
                    "EndTime": "2023-8-19 14:56:48",
                    "TestPaperId": 4,
                    "TestPaperName": "试卷测试2",
                    "ExamDuration": 100,
                    "FullMarks": 100,
                    "PassScore": 60,
                    "State": 2
                }
            ],
            "ExamStudentArr": [
                {
                    "StudentId": 32,
                    "TrueName": "学24"
                }
            ]
        }
    ],
    "msg": "操作成功"
}
```

#### 管理员 未审核改为已审核
请求url  http://192.168.0.123:7566/Manage/EditReViewExamByExamId


请求方式  post 

请求参数 
``` json
{
    "ExamId": 1
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


#### 查看考试信息列表 根据考试id查询
请求url  http://192.168.0.123:7566/Manage/GetExamByExamId?ExamId=1

请求方式  get 

请求参数 
ExamId=1

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "SchoolId": 1,
        "ExamName": "测试试卷35",
        "ExamDescribe": "试卷描述2",
        "ExamStatus": 1,
        "FaceVerify": 1,
        "ExamSessionArr": [
            {
                "Id": 1,
                "ExamId": 1,
                "StartTime": "2021-2-3",
                "EndTime": "2023-2-3",
                "TestPaperId": "3",
                "TestPaperName": "测试试卷3",
                "ExamDuration": 60,
                "FullMarks": 100,
                "PassScore": 60
            },
            {
                "Id": 2,
                "ExamId": 1,
                "StartTime": "2021-2-3",
                "EndTime": "2023-2-3",
                "TestPaperId": "2",
                "TestPaperName": "测试试卷1",
                "ExamDuration": 60,
                "FullMarks": 100,
                "PassScore": 60
            }
        ],
        "ExamStudentArr": [
            {
                "StudentId": 6,
                "TrueName": "五五"
            },
            {
                "StudentId": 8,
                "TrueName": "轻轻巧巧"
            },
            {
                "StudentId": 9,
                "TrueName": "学1"
            },
            {
                "StudentId": 10,
                "TrueName": "学2"
            }
        ]
    },
    "msg": "操作成功"
}
```

#### 删除考试（所有当前考试主题所关联的任何信息）

#### 设置考试主题是否开启人脸
#### 设置考试是否开启
#### 新增补考
请求url  http://192.168.0.123:7566/Manage/AddExamReset

请求方式  post 

请求参数 
``` json
{
    "OldExamSessionId": 1,
    "StudentIdArr": [ // 学生id数组
        6,
        9
    ],
    "OldExamId": 1,
    "StartExamTime": "",
    "EndExamTime": ""
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



#### 补考列表  根据场次id 查询
请求url  http://192.168.0.123:7566/Manage/GetExamResetByExamSessionId?ExamSessionId=2

请求方式  get 

请求参数 
ExamSessionId=2

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "StudentId": 6,
            "TrueName": "五五",
            "TestPaperName": "测试试卷1",
            "OldExamSessionId": 2,
            "OldExamId": 1,
            "StartExamTime": "2023-07-20",
            "EndExamTime": "2023-07-24",
            "Score": -1,
            "Status": 0
        },
        {
            "Id": 2,
            "StudentId": 9,
            "TrueName": "学1",
            "TestPaperName": "测试试卷1",
            "OldExamSessionId": 2,
            "OldExamId": 1,
            "StartExamTime": "2023-07-20",
            "EndExamTime": "2023-07-24",
            "Score": -1,
            "Status": 0
        }
    ],
    "msg": "操作成功"
}
```

#### 查看需要补考人员 -- 根据场次查询  查询考试未及格、缺考、补考缺考的学生

请求url  http://192.168.0.123:7566/Manage/GetCurrentExamSessionBKStudent?ExamSessionId=1

请求方式  get 

请求参数 
ExamSessionId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "StudentId": 8,
            "TrueName": "轻轻巧巧"
        },
        {
            "StudentId": 9,
            "TrueName": "学1"
        },
        {
            "StudentId": 10,
            "TrueName": "学2"
        }
    ],
    "msg": "操作成功"
}
```



#### 成绩查询--  考试场次列表 显示每个场次的平均分



请求url  http://192.168.0.123:7566/Manage/GetExamOverView?SchoolId=1

请求方式  get 

请求参数 
SchoolId=1

返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "ExamSessionId": 1,
            "ExamId": 1,
            "ExamSession": "测试试卷35/试卷测试1",
            "MajorName": "",
            "Score": "",
            "StartTime": "",
            "AvgScore": 0,
            "UnexaminedNum": 0,
            "ExaminedNum": 0,
            "ExamDuration": 0,
            "ExamPeopleSumNum": 0
        },
        {
            "ExamSessionId": 2,
            "ExamId": 1,
            "ExamSession": "测试试卷35/试卷测试2",
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
#### 成绩查询-- 单个场次的学生分数列表

请求url  http://192.168.0.123:7566/Manage/GetExamSessionResult?ExamSessionId=1&ExamId=1

请求方式  get 

请求参数 
ExamSessionId=1
ExamId=1

返回数据

``` json
 {
    "code": 1,
    "data": [
        {
            "Id": 9,
            "TrueName": "五五",
            "IDNumber": "422202199402131313",
            "StudentId": 6,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "UseTime": 0,
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
             "ExamImageArr": [
                {
                    "Id": 27,
                    "ExamId": 24,
                    "ExamSessionId": 25,
                    "StudentId": 113,
                    "ImagePath": "Resources/ExamImage/25/113/dd7ea816137334210ed2b33fc54f9001.png",
                    "CreateTime": "2023-11-17 14:33:59"
                }]
        },
        {
            "Id": 13,
            "TrueName": "轻轻巧巧",
            "IDNumber": "422202199202131313",
            "StudentId": 8,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "UseTime": 0,
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
             "ExamImageArr": [
                {
                    "Id": 27,
                    "ExamId": 24,
                    "ExamSessionId": 25,
                    "StudentId": 113,
                    "ImagePath": "Resources/ExamImage/25/113/dd7ea816137334210ed2b33fc54f9001.png",
                    "CreateTime": "2023-11-17 14:33:59"
                }]
        },
        {
            "Id": 15,
            "TrueName": "学1",
            "IDNumber": "111121212",
            "StudentId": 9,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "UseTime": 0,
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
             "ExamImageArr": [
                {
                    "Id": 27,
                    "ExamId": 24,
                    "ExamSessionId": 25,
                    "StudentId": 113,
                    "ImagePath": "Resources/ExamImage/25/113/dd7ea816137334210ed2b33fc54f9001.png",
                    "CreateTime": "2023-11-17 14:33:59"
                }]
        },
        {
            "Id": 17,
            "TrueName": "学2",
            "IDNumber": "111121212",
            "StudentId": 10,
            "ExamId": 1,
            "ExamSessionId": 1,
            "StartExamTime": "",
            "EndExamTime": "",
            "UseTime": 0,
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
             "ExamImageArr": [
                {
                    "Id": 27,
                    "ExamId": 24,
                    "ExamSessionId": 25,
                    "StudentId": 113,
                    "ImagePath": "Resources/ExamImage/25/113/dd7ea816137334210ed2b33fc54f9001.png",
                    "CreateTime": "2023-11-17 14:33:59"
                }]
        }
    ],
    "msg": "操作成功"
}
```



#### 站点用户登录
 请求地址  http://192.168.0.123:7566/StandUser/LoginStandUser

 请求方式 post
``` json
{
    "StandUserAccount":"stand123",
    "StandUserPwd":"123"
}
```
返回数据
``` json
{
    "code": 1,
    "data": {
        "StandUserId": 1,
        "StandUserName": "stand123",
        "StandUserAccount": "stand123",
        "StandUserPwd": "",
        "StandId": 1
    },
    "msg": "登录成功",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6InN0YW5kMTIzIiwiVXNlclR5cGUiOjIsImV4cCI6MTcwNjg0ODQzMSwiaXNzIjoibHgtand0In0.4MRB2bpUn3YZaTj7v_qqGct3Gn1th9v7UjE3yqKjRVo"
}
```

#### 根据站点查询成绩
http://192.168.0.124:7566/StandUser/QueryStudentResultByStandId?StandId=1&CourseId=&IsPassScore=

请求方式  get 

请求参数 
StandId=1   //必填 数字
CourseId=1  //可为空  不为空的时候只能为数字
IsPassScore=0 //0或者空值 都查  1 查及格 2 查不及格
返回数据

``` json
{
    "code": 1,
    "data": [
        {
            "RowNumber": 1,
            "TrueName": "胡奇",
            "CourseName": "工程检测技术",
            "CourseCode": "01538",
            "ExamNumber": "014322220250",
            "SchoolName": "清华大学",
            "Score": 30
        }
    ],
    "msg": "操作成功"
}
```

#### 根据站点输出execl
http://192.168.0.124:7566/StandUser/QueryStudentResultExeclByStandId?StandId=1&CourseId=&IsPassScore=
请求方式  get 

请求参数 
StandId=1   //必填 数字
CourseId=1  //可为空  不为空的时候只能为数字
IsPassScore=0 //0或者空值 都查  1 查及格 2 查不及格


返回数据
 成功直接返回一个execl 前端可直接下载

#### 成绩查询-- 单个考场的统计图表
#### 按考试主题导出成绩
#### 按站点，分考试主题 导出成绩



#### 考试通知新增
http://192.168.0.124:7566/Manage/AddExamNotice

请求类型 post

请求参数
``` json
{
    "ExamId":1,
    "Title":"",
    "SchoolName":"1",
    "CourseName":"1",
    "CourseCode":"",
    "Context":""
}
```

返回结果
``` json 
 {
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 考试通知修改
http://192.168.0.124:7566/Manage/EditExamNotice

请求类型 post

请求参数
``` json
{
    "ExamId":1,
    "Title":"2222222",
    "SchoolName":"1",
    "CourseName":"22221",
    "CourseCode":"222323",
    "Context":"123123"
}
```

返回结果
``` json 
 {
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 考试通知删除
http://192.168.0.124:7566/Manage/DelExamNotice

请求类型 post

请求参数
``` json
{
    "ExamId":1
}
```

返回结果
``` json 
 {
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 考试通知查询
请求url  http://192.168.0.124:7566/Manage/GetExamNoticeByExamId?ExamId=1

请求方式  get 

请求参数 
ExamId=1

返回数据

``` json
{
    "code": 1,
    "data": {
        "Id": 2,
        "ExamId": 1,
        "Title": "",
        "SchoolName": "1",
        "CourseName": "1",
        "CourseCode": "",
        "Context": ""
    },
    "msg": "成功"
}
```
#### 管理员登录
 请求地址  http://192.168.0.123:7566/Manage/LoginAdmin

 请求方式 post
``` json
{
    "AdminAccount":"admin",
    "AdminPassword":"123"
}
```
返回数据
``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "AdminName": "admin",
        "AdminAccount": "admin",
        "AdminPassword": ""
    },
    "msg": "登录成功",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6ImFkbWluIiwiZXhwIjoxNjg2MDQyODM0LCJpc3MiOiJseC1qd3QifQ.g6lcyOhg9Q7xOUZNVPuzZXB12ph25d9I5rxgP_S1h2Q"
}
```

### 学校管理
#### 新增学校

http://192.168.0.123:7566/Manage/AddSchool

请求类型 post

请求参数
``` json
{
"SchoolName":"武汉大学",
"SchoolAddress":"珞珈山职业技术学院"
}
```

返回结果
``` json 
 {
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 学校列表
http://192.168.0.123:7566/Manage/SchoolList

请求类型 get

请求参数
 无

返回结果

``` json
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "SchoolName": "清华大学",
            "SchoolAddress": "五道口"
        },
        {
            "Id": 4,
            "SchoolName": "武汉大学",
            "SchoolAddress": "珞珈山职业技术学院"
        }
    ],
    "msg": "成功"
}
```
#### 修改学校
http://192.168.0.123:7566/Manage/EditSchool


请求类型 post

请求参数
``` json
{
"SchoolName":"清华大学",
"SchoolAddress":"五道口地铁站",
"Id":1
}
```

返回结果
``` json 
 {
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除学校
http://192.168.0.123:7566/Manage/AddSchool


请求类型 post

请求参数
``` json
{
"Id":3
}
```

返回结果
``` json 
 {
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```




### 学习计划
#### 查询当前学习计划


#### 添加学习计划
url http://192.168.0.113:7566/Manage/AddPlan
请求类型 post

请求参数
``` json
{
    "CourseId": 24,
    "PlanName": "学习计划1",
    "CourseRatio": 20,
    "ExamRatio": 30,
    "TrainRatio": 50,
    "TeacherId": 1
}
```

返回结果
``` json 
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```


#### 修改学习计划
url http://192.168.0.113:7566/Manage/EditPlan

请求类型 post

请求参数
``` json
{
    "PlanId":1,
    "PlanName": "学习计划1",
    "CourseRatio": 20,
    "ExamRatio":50,
    "TrainRatio": 30
}
```

返回结果
``` json 
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除学习计划
url http://192.168.0.113:7566/Manage/DelPlan
请求类型 post

请求参数
``` json
{
    "PlanId":1
}
```

返回结果
``` json 
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 查询学习计划
url http://192.168.0.113:7566/Manage/QueryPlan

请求类型  get 

无参数

返回值

```json
{
    "code": 1,
    "data": [
        {
            "PlanId": 2,
            "CourseId": 1,
            "PlanName": "学习计划1",
            "CourseRatio": 20,
            "ExamRatio": 30,
            "TrainRatio": 40,
            "TeacherId": 0,
            "CourseName": 0,
            "PlanStudentCount": 0,
            "PlanExamCount": 0,
            "PlanTrainCount": 0
        },
        {
            "PlanId": 3,
            "CourseId": 24,
            "PlanName": "学习计划1",
            "CourseRatio": 20,
            "ExamRatio": 30,
            "TrainRatio": 50,
            "TeacherId": 0,
            "CourseName": 0,
            "PlanStudentCount": 0,
            "PlanExamCount": 0,
            "PlanTrainCount": 0
        }
    ],
    "msg": "操作成功"
}

```
 
#### 根据学习计划id 和学习计划中的学生查询课程进度
url http://192.168.0.109:7566/Manage/QueryPlanCourseProgress?PlanId=2&StudentId=1

请求类型  get 

参数
PlanId=2
StudentId=1
返回值

```json
{
    "code": 1,
    "data": {
        "CourseName": "课程3",
        "Digest": "",
        "CourseId": 25,
        "TeacherId": 1,
        "TeacherName": "tea",
        "SchoolName": "清华大学",
        "CollegeName": "美术学院",
        "MajorId": 1,
        "MajorName": "计算机科学与技术",
        "ChapterSum": 5,
        "StudentSum": 4,
        "CourseStartTime": "2023-06-08",
        "CourseEndTime": "2023-06-15",
        "ChapterOrder": 0,
        "LearningRate": 0,
        "FilePath": "Resources/Img/27411e0136569b2d318a166307ff5a91.png",
        "IsCurrentStudy": 3
    },
    "msg": "操作成功"
}
```
 
#### 根据学习计划id 和学习计划中的学生查询考试进度

url http://192.168.0.109:7566/Manage/QueryPlanExamProgress?PlanId=2&StudentId=1

请求类型  get 

参数
PlanId=2
StudentId=1
返回值

```json
{
    "code": 1,
    "data": [
        {
            "Id": 940841,
            "StudentId": 1,
            "ExamId": 49,
            "ExamName": "工程地质学基础",
            "ExamSessionId": 60,
            "StartExamTime": "",
            "EndExamTime": "",
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
            "ExamZT": 2,
            "MajorId": 0,
            "CourseId": 63,
            "MajorName": "",
            "CourseName": "工程地质学基础",
            "FullMarks": 100,
            "PassScore": 60,
            "SessionNum": 1,
            "ExamDuration": 60,
            "QuestionNum": 1,
            "ResetStartExamTime": "",
            "ResetEndExamTime": "",
            "SessionStartExamTime": "2024-2-28 10:34:38",
            "SessionEndExamTime": "2024-11-30 10:34:43",
            "FaceVerify": 0
        },
        {
            "Id": 940842,
            "StudentId": 1,
            "ExamId": 48,
            "ExamName": "地理信息系统",
            "ExamSessionId": 59,
            "StartExamTime": "",
            "EndExamTime": "",
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
            "ExamZT": 2,
            "MajorId": 0,
            "CourseId": 62,
            "MajorName": "",
            "CourseName": "地理信息系统",
            "FullMarks": 100,
            "PassScore": 60,
            "SessionNum": 1,
            "ExamDuration": 60,
            "QuestionNum": 1,
            "ResetStartExamTime": "",
            "ResetEndExamTime": "",
            "SessionStartExamTime": "2024-2-28 10:34:11",
            "SessionEndExamTime": "2024-11-30 10:34:15",
            "FaceVerify": 0
        }
    ],
    "msg": "操作成功"
}
```

#### 根据学习计划id 和学习计划中的学生查询训练进度

url http://192.168.0.109:7566/Manage/QueryPlanTrainProgress?PlanId=2&StudentId=1

请求类型  get 

参数
PlanId=2
StudentId=1
返回值

```json
{
    "code": 1,
    "data": [
        {
            "PlanTrainId": 1,
            "PlanId": 2,
            "QuestionId": 6,
            "QuestionType": 1,
            "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
            "QuestionContent": "[\"A. 格斯拉兄弟将自己的店开在一条横街上，这条横街座落在伦敦市的西区。\",\"B. 我不断地叩问历史的发展，产生了从中找出一条贯穿其中的脉络的愿望。\",\"C. 老王为新春写的这幅对联，寄托了美好的祝福，展现了书法大境界。\"]",
            "AnswerSteps": "1",
            "TrainScore": 100
        },
        {
            "PlanTrainId": 2,
            "PlanId": 2,
            "QuestionId": 66,
            "QuestionType": 1,
            "QuestionName": "测试1",
            "QuestionContent": "[\"1\",\"2\",\"4\"]",
            "AnswerSteps": "",
            "TrainScore": 0
        }
    ],
    "msg": "操作成功"
}
```

#### 添加学习计划的学生
url http://192.168.0.114:7566/Manage/AddPlanStudent
请求方式 post
请求参数
```json
[{
    "PlanId": 2,
    "StudentId": 1
}
,{
    "PlanId": 2,
    "StudentId":3
}]
```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除学习计划学生
url http://192.168.0.114:7566/Manage/DelPlanStudent
请求方式 post
请求参数
```json
    {
        "PlanStudentId": 4
    }

```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```


#### 删除学习计划全部学生
url http://192.168.0.114:7566/Manage/DelAllPlanStudent
请求方式 post
请求参数
```json
 
    {
        "PlanId": 1
    }

```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 查询学习计划学生
url http://192.168.0.114:7566/Manage/QueryPlanStudent?PlanId=1

请求方式 get 
参数 PlanId=1

返回值
```json
{
    "code": 1,
    "data": [
        {
            "PlanStudentId": 4,
            "PlanId": 1,
            "StudentId": 1,
            "TrueName": "杜拉拉2",
            "StudentAccount": "dll"
        },
        {
            "PlanStudentId": 5,
            "PlanId": 1,
            "StudentId": 3,
            "TrueName": "杜靠靠",
            "StudentAccount": "dkk"
        }
    ],
    "msg": "操作成功"
}
```

#### 往学习计划添加考场
url http://192.168.0.114:7566/Manage/AddPlanExam
请求方式 post
请求参数
```json
[
    {
        "PlanId": 1,
        "ExamSessionId":60,
        "ExamId":49
    },
      {
        "PlanId": 1,
        "ExamSessionId":59,
        "ExamId":48
    }
]
```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除学习计划考场
url http://192.168.0.114:7566/Manage/DelPlanExam
请求方式 post
请求参数
```json
 
    {
        "PlanexamId": 23
    }

```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```

#### 删除学习计划全部训考场
url http://192.168.0.114:7566/Manage/DelAllPlanExam
请求方式 post
请求参数
```json
 
    {
        "PlanId": 1
    }

```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```

#### 查询学习计划考场
url http://192.168.0.114:7566/Manage/QueryPlanExam?PlanId=1

请求方式 get 
参数 PlanId=1

返回值
```json
{
    "code": 1,
    "data": [
        {
            "PlanexamId": 23,
            "PlanId": 1,
            "ExamSessionId": 60,
            "ExamId": 49,
            "ExamName": "工程地质学基础"
        },
        {
            "PlanexamId": 24,
            "PlanId": 1,
            "ExamSessionId": 59,
            "ExamId": 48,
            "ExamName": "地理信息系统"
        }
    ],
    "msg": "操作成功"
}
```

#### 往学习计划添加训练题
url http://192.168.0.114:7566/Manage/AddPlanTrain
请求方式 post
请求参数
```json
[
    {
        "PlanId": 1,
        "QuestionId": 2,
        "QuestionType": 1
    },
        {
        "PlanId": 1,
        "QuestionId": 62,
        "QuestionType": 1
    }
]
```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```


#### 删除学习计划训练题
url http://192.168.0.114:7566/Manage/DelPlanTrain
请求方式 post
请求参数
```json
 
    {
        "PlanTrainId": 1
    }

```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```


#### 删除学习计划全部训练题
url http://192.168.0.114:7566/Manage/DelAllPlanTrain
请求方式 post
请求参数
```json
 
    {
        "PlanId": 1
    }

```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```

#### 查询学习计划训练题
url http://192.168.0.114:7566/Manage/QueryPlanTrain?PlanId=1

请求方式 get 
参数 PlanId=1

返回值
```json
{
    "code": 1,
    "data": [
        {
            "PlanTrainId": 1,
            "PlanId": 1,
            "QuestionId": 2,
            "QuestionType": 1,
            "QuestionName": "单选题"
        },
        {
            "PlanTrainId": 2,
            "PlanId": 1,
            "QuestionId": 62,
            "QuestionType": 0,
            "QuestionName": ""
        }
    ],
    "msg": "操作成功"
}
```

### 学生学习计划相关接口

#### 学生训练计划内的训练提交

url http://127.0.0.1:7566/Student/AddQuestionRecord
请求方式 post
请求参数
```json
[
    {
        "QuestionId": 6,
        "StudentId": 3,
        "CreateTime": "2022-11-23 18:00:00",
        "AnswerSteps": "1",
        "TrueAnswer": "1",
        "TrainScore": 100,
        "PlanId": 2
    }
]
```
返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```





#### 查询当前学生的所有关联的学习计划
url http://127.0.0.1:7566/Student/QuertStudentPlan?StudentId=1
请求方式 get
请求参数
学生id StudentId=1



返回值
```json
{
    "code": 1,
    "data": [
        {
            "PlanId": 2,
            "CourseId": 25,
            "PlanName": "学习计划1",
            "CourseRatio": 20,
            "ExamRatio": 30,
            "TrainRatio": 40,
            "TeacherId": 1,
            "CourseName": 0,
            "PlanStudentCount": 2,
            "PlanExamCount": 5,
            "PlanTrainCount": 5
        }
    ],
    "msg": "操作成功"
}
```




#### 上传照片到当前学习计划
http://127.0.0.1:7566/Student/UploadPlanImage
请求方式 post
请求参数
 files 照片内容

 data 数据 ：
 ```json
 {
    "PlanId": 2,
    "PlanStudentId": 1
}
 ```



返回值
```json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```


#### 根据学习计划id 和学习计划中的学生查询课程进度
url http://192.168.0.109:7566/Student/QueryPlanCourseProgress?PlanId=2&StudentId=1

请求类型  get 

参数
PlanId=2
StudentId=1
返回值

```json
{
    "code": 1,
    "data": {
        "CourseName": "课程3",
        "Digest": "",
        "CourseId": 25,
        "TeacherId": 1,
        "TeacherName": "tea",
        "SchoolName": "清华大学",
        "CollegeName": "美术学院",
        "MajorId": 1,
        "MajorName": "计算机科学与技术",
        "ChapterSum": 5,
        "StudentSum": 4,
        "CourseStartTime": "2023-06-08",
        "CourseEndTime": "2023-06-15",
        "ChapterOrder": 0,
        "LearningRate": 0,
        "FilePath": "Resources/Img/27411e0136569b2d318a166307ff5a91.png",
        "IsCurrentStudy": 3
    },
    "msg": "操作成功"
}
```
 
#### 根据学习计划id 和学习计划中的学生查询考试进度

url http://192.168.0.109:7566/Student/QueryPlanExamProgress?PlanId=2&StudentId=1

请求类型  get 

参数
PlanId=2
StudentId=1
返回值

```json
{
    "code": 1,
    "data": [
        {
            "Id": 940841,
            "StudentId": 1,
            "ExamId": 49,
            "ExamName": "工程地质学基础",
            "ExamSessionId": 60,
            "StartExamTime": "",
            "EndExamTime": "",
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
            "ExamZT": 2,
            "MajorId": 0,
            "CourseId": 63,
            "MajorName": "",
            "CourseName": "工程地质学基础",
            "FullMarks": 100,
            "PassScore": 60,
            "SessionNum": 1,
            "ExamDuration": 60,
            "QuestionNum": 1,
            "ResetStartExamTime": "",
            "ResetEndExamTime": "",
            "SessionStartExamTime": "2024-2-28 10:34:38",
            "SessionEndExamTime": "2024-11-30 10:34:43",
            "FaceVerify": 0
        },
        {
            "Id": 940842,
            "StudentId": 1,
            "ExamId": 48,
            "ExamName": "地理信息系统",
            "ExamSessionId": 59,
            "StartExamTime": "",
            "EndExamTime": "",
            "Score": -1,
            "ExamStatus": 0,
            "ExamType": 0,
            "ExamZT": 2,
            "MajorId": 0,
            "CourseId": 62,
            "MajorName": "",
            "CourseName": "地理信息系统",
            "FullMarks": 100,
            "PassScore": 60,
            "SessionNum": 1,
            "ExamDuration": 60,
            "QuestionNum": 1,
            "ResetStartExamTime": "",
            "ResetEndExamTime": "",
            "SessionStartExamTime": "2024-2-28 10:34:11",
            "SessionEndExamTime": "2024-11-30 10:34:15",
            "FaceVerify": 0
        }
    ],
    "msg": "操作成功"
}
```

#### 根据学习计划id 和学习计划中的学生查询训练进度

url http://192.168.0.109:7566/Student/QueryPlanTrainProgress?PlanId=2&StudentId=1

请求类型  get 

参数
PlanId=2
StudentId=1
返回值

```json
{
    "code": 1,
    "data": [
        {
            "PlanTrainId": 1,
            "PlanId": 2,
            "QuestionId": 6,
            "QuestionType": 1,
            "QuestionName": "1 ．下列句子中没有错别字的一项是 ( )",
            "QuestionContent": "[\"A. 格斯拉兄弟将自己的店开在一条横街上，这条横街座落在伦敦市的西区。\",\"B. 我不断地叩问历史的发展，产生了从中找出一条贯穿其中的脉络的愿望。\",\"C. 老王为新春写的这幅对联，寄托了美好的祝福，展现了书法大境界。\"]",
            "AnswerSteps": "1",
            "TrainScore": 100
        },
        {
            "PlanTrainId": 2,
            "PlanId": 2,
            "QuestionId": 66,
            "QuestionType": 1,
            "QuestionName": "测试1",
            "QuestionContent": "[\"1\",\"2\",\"4\"]",
            "AnswerSteps": "",
            "TrainScore": 0
        }
    ],
    "msg": "操作成功"
}
```





###  修改工具表
#### 设置静态资源类型\

1：OSS和CDN都不开
2：OSS和CDN都开
3：只开OSS
4：只开CDN
#### 设置oss路径
#### 设置CDN链接
#### 设置本地路径配置
#### 设置全局人脸识别

#### 系统功能设置
http://192.168.0.123:7566/Manage/EditTool

请求类型 post

请求参数
 
 ```json
{
    "StaticResourcesType": 1,
    "FaceVerify": 0
}
 ```

返回结果

``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```



#### 获取系统设置
http://192.168.0.123:7566/Manage/GetTool

请求类型 get

请求参数
 无

返回结果

``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "StaticResourcesType": 1,
        "FaceVerify": 0
    },
    "msg": "操作成功"
}
```

 
### 公告管理
#### 添加公告
http://192.168.0.123:7566/Manage/AddNotice

请求类型 post

请求参数
``` json

{
    "NoticeTitle": "test",
    "NoticeContent": "test",
    "SendUser": "test",
    "NoticeLevel": 1,
    "NoticeType": "2"
}
```

返回结果


``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 修改公告
http://192.168.0.123:7566/Manage/EditNotice

请求类型 post

请求参数
``` json

{
    "NoticeTitle": "test1",
    "NoticeContent": "test",
    "SendUser": "test",
    "NoticeLevel": 1,
    "NoticeType": "2",
    "Id":1
}
```

返回结果


``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 删除公告
http://192.168.0.123:7566/Manage/DelNotice

请求类型 post

请求参数
``` json
{
    "Id":3
}
```

返回结果


``` json
{
    "code": 1,
    "data": "{}",
    "msg": "操作成功"
}
```
#### 公共列表 （不需要token就能请求）获取日期前五的公告  后台

http://192.168.0.123:7566/Manage/NoticeList

请求类型 get

请求参数
 无

返回结果


``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "Time": "2022-05-02 10:23:20",
            "NoticeTitle": "222",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "1"
        },
        {
            "Id": 1,
            "Time": "2020-05-01 10:20:30",
            "NoticeTitle": "aaa",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "2"
        }
    ],
    "msg": "成功"
}
```







#### 公共列表 （不需要token就能请求）获取日期前五的公告  学生

http://192.168.0.123:7566/Student/NoticeList

请求类型 get

请求参数
 无

返回结果


``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "Time": "2022-05-02 10:23:20",
            "NoticeTitle": "222",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "1"
        },
        {
            "Id": 1,
            "Time": "2020-05-01 10:20:30",
            "NoticeTitle": "aaa",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "2"
        }
    ],
    "msg": "成功"
}
```


#### 公共列表所有数据 后台

http://192.168.0.123:7566/Manage/NoticeListAll


请求类型 get

请求参数
 无

返回结果


``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "Time": "2022-05-02 10:23:20",
            "NoticeTitle": "222",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "1"
        },
        {
            "Id": 1,
            "Time": "2020-05-01 10:20:30",
            "NoticeTitle": "aaa",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "2"
        }
    ],
    "msg": "成功"
}
```


#### 公共列表所有数据 学生

http://192.168.0.123:7566/Manage/NoticeListAll


请求类型 get

请求参数
 无

返回结果


``` json
{
    "code": 1,
    "data": [
        {
            "Id": 2,
            "Time": "2022-05-02 10:23:20",
            "NoticeTitle": "222",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "1"
        },
        {
            "Id": 1,
            "Time": "2020-05-01 10:20:30",
            "NoticeTitle": "aaa",
            "NoticeContent": "11",
            "SendUser": "11",
            "NoticeLevel": 1,
            "NoticeType": "2"
        }
    ],
    "msg": "成功"
}
```

#### 获取某一个公告 根据公告id 后台

http://192.168.0.123:7566/Manage/GetNoticeById?NoticeId=1

请求类型 get

请求参数
NoticeId=1

返回结果


``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "Time": "2023-06-02 15:23:00",
        "NoticeTitle": "test1",
        "NoticeContent": "test1",
        "SendUser": "test1",
        "NoticeLevel": 2,
        "NoticeType": "3"
    },
    "msg": "成功"
}
```

#### 获取某一个公告 根据公告id 学生
http://192.168.0.123:7566/Student/GetNoticeById?NoticeId=1

请求类型 get

请求参数
 NoticeId=1 

返回结果


``` json
{
    "code": 1,
    "data": {
        "Id": 1,
        "Time": "2023-06-02 15:23:00",
        "NoticeTitle": "test1",
        "NoticeContent": "test1",
        "SendUser": "test1",
        "NoticeLevel": 2,
        "NoticeType": "3"
    },
    "msg": "成功"
}
```

#### 获取全局是否开启人脸识别
http://192.168.0.123:7566/Student/GetFaceVerify

请求类型 get

请求参数
无

返回结果


``` json
{
    "code": 1,
    "data": {
        "FaceVerify": 0,
        "SeparateFaceVerify":0,
        "MidiFlag":0
    },
    "msg": "操作成功"
}
```

#### 滑动条图片获取
 请求 url http://ebarotech.cn/StudyExamPlatformAPI/File/SlideImg

 请求方式  get  


####  获取系统时间

 请求 url  http://192.168.0.123:7566/Student/GetTime
 请求方式  get  
返回
2023-08-03 16:20:41





#### 获取文件存放类型

 请求 url  http://192.168.0.124:7566/Student/GetHttpUrl
 请求方式  get  
返回
``` json
{
    "code": 1,
    "data": {
        "OSShttp": "https://studyexamplatform.oss-cn-beijing.aliyuncs.com/",
        "StaticResourcesType": 3
    },
    "msg": "操作成功"
}
```