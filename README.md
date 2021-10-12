# workchatapp　golang版企业微信应用sdk

## 安装使用

    go get github.com/go-laoji/workchatapp

## 使用样例

样例中配置为测试使用,请根据实际情况修改，但切记不要公开secret配置

    package main

    import (
        "github.com/go-laoji/workchatapp"
        "log"
    )
    
    func main() {
        defer func() {
            if e := recover(); e != nil {
                log.Println("recover", e)
            }
        }()
        var testConfig = workchatapp.WorkChatConfig{
            CorpId:        "ww190690c489d2eb53",
            ContactSecret: "08tnu5LGrsbKwvEDfTGlBMFMw3CsUCwRMavxvkLZSH8",
            AppId:         "1000002",
            AppSecret:     "pedn4nqraARPFOG_A-aVFz1F9pp1sdR-3K1fsCpTwg0",
        }
        weworkApp := workchatapp.NewWorkChatApp(testConfig)
        resp := weworkApp.CorpTagList([]string{}, []string{})
        log.Println(resp)
    }



## 接口列表(更新ing...)

- 通讯录管理
    - 成员管理
      - [x] 创建成员
      - [x] 读取成员
      - [x] 更新成员
      - [x] 删除成员
      - [x] 批量删除成员
      - [x] 获取部门成员
      - [x] 获取部门成员详情
      - [x] userid与openid互换
      - [ ] 二次验证
      - [ ] 邀请成员
      - [ ] 获取加入企业二维码
      - [ ] 获取企业活跃成员数
    - 部门管理
      - [x] 创建部门
      - [x] 更新部门
      - [x] 删除部门
      - [x] 获取部门列表
    - 标签管理
      - [x] 创建标签
      - [x] 更新标签名字
      - [x] 删除标签
      - [x] 获取标签成员
      - [x] 删除标签成员
      - [x] 获取标签列表
    - 异步批量接口
      - [ ] 增量更新成员
      - [ ] 全量覆盖成员
      - [ ] 全量覆盖部门
      - [ ] 获取异步任务结果
    - 通讯录回调通知
      - [ ] 成员变更通知
      - [ ] 部门变更通知
      - [ ] 标签变更通知
      - [ ] 异步任务完成通知
    - 互联企业
      - [ ] 获取应用的可见范围
      - [ ] 获取互联企业成员详细信息
      - [ ] 获取互联企业部门成员
      - [ ] 获取互联企业部门成员详情
      - [ ] 获取互联企业部门列表
    - 异步导出接口
      - [ ] 导出成员
      - [ ] 导出成员详情
      - [ ] 导出部门
      - [ ] 导出标签成员
      - [ ] 获取导出结果
      - [ ] 导出任务完成通知

- 客户联系
    - 企业服务人员管理
      - [x] 获取配置了客户联系功能的成员列表
    - 客户联系「联系我」管理
      - [x] 配置客户联系「联系我」方式
      - [x] 获取企业已配置的「联系我」方式
      - [x] 获取企业已配置的「联系我」列表
      - [x] 更新企业已配置的「联系我」方式
      - [x] 删除企业已配置的「联系我」方式
      - [x] 结束临时会话
    - 客户管理
      - [x] 获取客户列表
      - [x] 获取客户详情
      - [x] 批量获取客户详情
      - [x] 修改客户备注信息
    - 客户联系规则组管理
      - [x] 获取规则组列表
      - [x] 获取规则组详情
      - [x] 获取规则组管理范围
      - [x] 创建新的规则组
      - [x] 编辑规则组及其管理范围
      - [x] 删除规则组
    - 客户标签管理
        - 管理企业标签
          - [x] 获取企业标签库
          - [x] 添加企业客户标签
          - [x] 编辑企业客户标签
          - [x] 删除企业客户标签
        - 编辑客户企业标签
    - 在职继承
      - [x] 分配在职成员的客户
      - [x] 查询客户接替状态
    - 离职继承
      - [x] 获取待分配的离职成员列表
      - [x] 分配离职成员的客户
      - [x] 查询客户接替状态
      - [x] 分配离职成员的客户群