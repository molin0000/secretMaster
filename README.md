# 序列战争--诡秘之主粉丝序列QQ群插件

欢迎大家给我打赏和发电~~ https://afdian.net/@molin

欢迎大家使用文字游戏编辑器，编写自己的文字副本： https://mission-editor.now.sh/


[《⭐️序列战争 - 玩家指南⭐️》](./playGuide.md)

## 更新记录

2020年3月20日 调高了宠物被投喂的概率，宠物迷路也有经验收入，提高10级以下的宠物pk的胜率。v3.0.2

2020年3月19日 宠物系统上线，回复宠物店体验。v3.0.1

2020年3月17日 紧急修复了退出阵营后还可以挑战的bug。 v2.8.6

2020年3月17日 修复加入阵营立刻退出后可能得到收益的bug。v2.8.5

2020年3月16日 修复GM指令中luck和level的bug。v2.8.4

2020年3月16日 提高竞赛灵性消耗到25点，提高竞赛奖金到150金镑。v2.8.3

2020年3月16日 修复了金镑存银行没有利息的bug  v2.8.2

2020年3月16日 限制了红包功能，序列7以下不可以发，序列7以上限制单次红包金额P=100*2^n, n∈[0,9]  v2.8.1

2020年3月16日 在生活菜单中增加了红包功能，现在可以使用红包来转账了。v2.8.0

2020年3月16日 修复了阵营加入后即休整状态的bug，修复了异常提示教会达到最大人数无法加入的bug；v2.7.9

2020年3月15日 修复了：1）女装直播等普通工作无法开始的问题；2）竞赛答题时间在对战中数值异常的问题；v2.7.7

2020年3月15日 增加了.supermaster全局管理员选项，限制自杀次数每天3次； v2.7.6

2020年3月14日 最近老是被封号，于是删除了彩票功能，增加了许愿功能。v2.7.4

2020年3月13日 阵营对战系统正式上线。使用方法请点击帮助菜单下面的详细帮助并阅读3.6章节。希望大家喜欢~ v2.7.3

2020年3月11日 增加了分群定时静默功能，可定时开启插件回复。v2.6.1

2020年3月11日 在生活菜单中增加“竞赛”功能，回复“竞赛”开启。v2.5.0

2020年3月10日 序列战争更新v2.4.0版本了，在人物卡位置新增了勋章🎖功能。勋章颁发给那些对序列战争游戏有特殊贡献的人，用以感谢他们的默默付出。贡献类型包括但不限于：1）提交bug；2）贡献副本；3）贡献探险或其它文案；4）提出有效建议并被采纳；5）⚡️⚡️⚡️(隐藏了)；勋章只能作者颁发，在机器人全局有效，即每个群，只要是同一个机器人，都可看到勋章。每种类型贡献最多得1枚勋章。谢谢大家~

2020年3月10日 修复了：1）教会入会费不打给教主的问题；2）尝试修复属性昵称为空的问题；以及一些其它bug和优化；v2.3.4

2020年3月8日 支持副本了 v2.3.0

2020年3月10日 修复bug，增加游玩指南文档 v2.3.4

2020年3月7日 修复许多的bug v2.2.8

2020年3月5日 修复自杀后时间异常，工作没清空的bug v2.2.4

2020年3月5日 修复了好几个bug v2.2.3

2020年3月4 序列战争v2.2.0版本发布。新版本改善了私聊体验，不用再@输入群号了。加机器人好友，然后发送：序列战争，开始体验吧~~  

2020年3月2 全新梳理和设计的重要更新版本，5大系统，新的功能等待你探索  v2.0.0

2020年2月21 增加了红剧场门票，增加了货币对接其它插件ini功能 v1.0.7

2020年2月16 增加分群开关，指令为`序列战争开`或`序列战争关` v1.0.3

2020年2月16 感谢flak.更新好多探险文案 v1.0.2

2020年2月14 更新部分序列名称，正式更名为序列战争

2020年2月11 增加查询排行榜指定人物属性功能，增加了大量探险文案。优化了文案添加方法。

2020年2月1 增加私聊查询功能，只需指令@群号私聊即可；在热心书友flak.的建议下，新增部分探险文案。

2020年1月16 增加命运途径默认幸运加25%特质，增加灵性属性，灵性枯竭则不加经验和金镑，修复删除人物的bug

2020年1月16 整理代码分布，修复水群bug

2020年1月15 尊名系统上线，现在序列3可以设置个性化的尊名了

2020年1月14 增加购买探险卷轴功能和删除人物功能

2020年1月13 基本功能完成

2020年1月10 开始开发并完成第一版

## 介绍

序列战争是一款在QQ群中酷Q机器人使用的聊天游戏插件。故事背景是《诡秘之主》的世界观。

策划群：1028799086

游玩群：1030551041，466238445

《诡秘之主》粉丝群：731419992

开发日志（三）

提前恭喜本群第一位真神即将诞生。愿众生在祂的指引下前进。

之前的几次大版本中我们逐步添加了金镑系统，尊名设置和探险系统，甜姜茶的滋味想必大家一定很喜欢。

在现行的0.4.5版本之中，我们发现由于金镑用途单一，导致大家花费大量的金镑在探险卷轴上，所以我们会增加商店系统，增添一些骗取金镑的新物品，海量超值货物，正在东拜朗中转。

随着真神的诞生，需要有新的教会引导群里迷茫的人们，所以我们准备开放组织建设功能，此功能会逐步完善，尽情期待，成为神灵的眷者或者是主教牧首，全看你的选择。

很多非凡者反应需要一些能够体现非凡能力的设定，于是在上次更新中增加了命运途径（怪物途径）的幸运，那么接下来我们会逐步完善各途径的技能，而技能的开发也是战斗系统的先行版。

希望大家能够健康聊天，合理水群，毕竟这只是个拥有统计功能的工具，我们只不过尽量让这个工具更有趣。

待更新内容

1）商店系统

售卖商品包括：

a. 探险卷轴

b. 灵性药剂

c. 权杖

d. 仪式材料

2）教会组织系统

必须序列3以上才能创建教会，通过购买权杖，获得创建教会资格。

眷者可以加入教会享受被动增益

3）技能系统

分为主动技能和被动技能

主动技能PK时使用

被动技能可由序列获得或由教会赐予。
