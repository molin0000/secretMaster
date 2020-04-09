import { Divider, Button, Switch, Table, Select, Input } from 'antd';
import styles from './config.css';
import { Component } from 'react';
import { apiGet, apiAsyncGet } from './utils/utils.js';
import router from 'umi/router'


const { Option } = Select;
const { TextArea } = Input;

const groupColumns = [
  {
    title: '序号',
    dataIndex: 'key',
    key: 'key',
    width: 50,
  },
  {
    title: '群号',
    dataIndex: 'group',
    key: 'group',
  },
  {
    title: '人数',
    dataIndex: 'member',
    key: 'member',
  },
  {
    title: '管理员',
    dataIndex: 'master',
    key: 'master',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    key: 'operation',
    render: (text, record) => (
      <span>
        <a href="/#">退群</a>
      </span>
    ),
  },
  {
    title: '开关',
    dataIndex: 'switch',
    key: 'switch',
    render: (text, record) => (
      <span>
        <Switch checkedChildren="开" unCheckedChildren="关" checked={text} />
      </span>
    ),
  },
  {
    title: '静默',
    dataIndex: 'silence',
    key: 'silence',
    render: (text, record) => (
      <span>
        <Switch checkedChildren="开" unCheckedChildren="关" checked={text} />
      </span>
    ),
  },
]

const activitiesColumns = [
  {
    title: '序号',
    dataIndex: 'key',
    key: 'key',
    width: 50,
  },
  {
    title: '关键字',
    dataIndex: 'word',
    key: 'word',
    ellipsis: true,
  },
  {
    title: '回复信息',
    dataIndex: 'reply',
    key: 'reply',
    ellipsis: true,
  },
  {
    title: '类型',
    dataIndex: 'type',
    key: 'type',
  },
  {
    title: '经验',
    dataIndex: 'exp',
    key: 'exp',
    width: 70,
  },
  {
    title: '金镑',
    dataIndex: 'money',
    key: 'money',
    width: 70,
  },
  {
    title: '灵力',
    dataIndex: 'magic',
    key: 'magic',
    width: 70,
  },
  {
    title: '操作',
    dataIndex: 'operation',
    key: 'operation',
    render: (text, record) => (
      <span>
        <a href="/#">删除</a>
      </span>
    ),
  },
  {
    title: '开关',
    dataIndex: 'switch',
    key: 'switch',
    render: (text, record) => (
      <span>
        <Switch checkedChildren="开" unCheckedChildren="关" defaultChecked />
      </span>
    ),
  },
]

const activitiesData = [
  {
    key: 1,
    word: "新年快乐",
    reply: "亚米收到你的祝福啦~红包给你！",
    type: "每年",
    exp: 0,
    money: 100,
    magic: 0,
  },
  {
    key: 2,
    word: "签到",
    reply: "签到成功！",
    type: "每天",
    exp: 6,
    money: 6,
    magic: 6,
  },
]

class Config extends Component {
  constructor(props) {
    super(props);

    this.state = {
      supermaster: 0,
      delay: 300,
      count: {
        "MissionCnt": 0,
        "QuestionCnt": 0,
        "RealPetCnt": 0,
        "SpiritPetCnt": 0,
        "PetFoodCnt": 0,
      },
      moneyMap: {
        "group": 0,
        "iniPath": "/home/user/coolq/data/app/com.qmt.demo/玩家数据.ini",
        "iniSection": "cQQ",
        "iniKey": "金币数量",
        "hasUpdate": false,
        "encode": ""
      },
      imageMode: {
        "enable": false,
        "lines": 0
      },
      textSegment: {
        "enable": false,
        "lines": 0
      },
      group: {
        globalSwitch: true,
        globalSilence: false,
        groups: [],
      },
      activities: [],
    }
  }

  async componentWillMount(props) {
    let locked = await apiGet('locked');
    if (!locked.data.data) {
      router.push("/register");
      return;
    }

    if (!global.unlocked) {
      router.push("/login");
      return;
    }

    apiAsyncGet('count', (res) => {
      console.log(res.data.data);
      this.setState({ count: res.data.data });
    });

    apiAsyncGet('group', (res) => {
      console.log(res.data.data);
      this.setState({ group: res.data.data });
    });

    apiAsyncGet('supermaster', (res) => {
      console.log(res.data.data);
      this.setState({ supermaster: res.data.data });
    });

    apiAsyncGet('delay', (res) => {
      console.log(res.data.data);
      this.setState({ delay: res.data.data.DelayMs });
    });

    apiAsyncGet('moneyMap', (res) => {
      console.log(res.data.data);
      let ret = res.data.data;
      ret.group = this.state.moneyMap.group;
      this.setState({ moneyMap: ret });
    });

    apiAsyncGet('imageMode', (res) => {
      console.log(res.data.data);
      this.setState({ imageMode: res.data.data });
    });

    apiAsyncGet('textSegment', (res) => {
      console.log(res.data.data);
      this.setState({ textSegment: res.data.data });
    });

    apiAsyncGet('activities', (res) => {
      console.log(res.data.data);
      this.setState({ activities: res.data.data });
    });

  }

  render() {
    return (
      <div className={styles.normal}>
        <div className={styles.body}>
          <div className={styles.msg}>更多设置选项，请查看使用手册，或私聊机器人发送“设置”查询和使用。</div>
          <div className={styles.inline}>
            <div className={styles.title}>插件主人（超级管理员）QQ：</div>
            <input className={styles.input} value={this.state.supermaster} />
            <Button type="primary" style={{ marginLeft: "40px" }}>保存</Button>
          </div>
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>消息回复延迟（毫秒）：</div>
            <input className={styles.input} style={{ marginLeft: "35px" }} value={this.state.delay} />
            <Button type="primary" style={{ marginLeft: "40px" }}>保存</Button>
          </div>
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>附件扫描：</div>
            <div className={styles.title} style={{ marginLeft: "20px" }}>副本：{this.state.count.MissionCnt} 个， 学识：{this.state.count.QuestionCnt} 个，现世宠物：{this.state.count.RealPetCnt} 个，灵界宠物：{this.state.count.SpiritPetCnt} 个，宠物投食：{this.state.count.PetFoodCnt} 个</div>
          </div>
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>分群管理：</div>
            <div className={styles.text} style={{ marginLeft: "360px" }}>全局开关：</div>
            <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.group.globalSwitch} />
            <div className={styles.text} style={{ marginLeft: "30px" }}>全局静默：</div>
            <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.group.globalSilence} />
          </div>
          <Table columns={groupColumns} dataSource={this.state.group.groups} size="small" />
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>货币映射：</div>
          </div>
          <div className={styles.inline}>
            <div className={styles.text2} >QQ群号：</div>
            <input className={styles.input} value={this.state.moneyMap.group} style={{ margin: "4px 0px 4px 104px", width: "120px", textAlign: "center" }} />
            <Button style={{ marginLeft: "20px", marginTop: "2px" }}>加载</Button>
            <div className={styles.text} style={{ marginLeft: "185px", marginTop: "6px" }}>启用：</div>
            <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.moneyMap.hasUpdate} style={{ marginTop: "6px"}} />
          </div>
          <div className={styles.inline}>
            <div className={styles.text2}>ini文件路径：</div>
            <input className={styles.input} style={{ margin: "4px 0px 4px 84px", width: "520px", textAlign: "left", paddingLeft: "10px" }} value={this.state.moneyMap.iniPath} />
          </div>
          <div className={styles.inline}>
            <div className={styles.text2}>节点为QQ号，关键字：</div>
            <input className={styles.input} style={{ margin: "6px 0px 6px 20px", width: "100px" }} value={this.state.moneyMap.iniKey} />
            <div className={styles.text2} style={{ marginLeft: "40px" }}>编码：</div>
            <Select defaultValue={this.state.moneyMap.encode?this.state.moneyMap.encode:"GB2312"} style={{ width: 124, margin: "4px 0" }}>
              <Option value="GB2312">GB2312</Option>
              <Option value="UTF8">UTF8</Option>
            </Select>
          </div>
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>图片模式：</div>
            <div className={styles.text} style={{ marginLeft: "40px" }}>触发行数：</div>
            <input className={styles.input} style={{ width: "200px" }} value={this.state.imageMode.lines} />
            <div className={styles.text} style={{ marginLeft: "210px" }}>启用：</div>
            <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.imageMode.enable} />
          </div>

          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>文字分段：</div>
            <div className={styles.text} style={{ marginLeft: "40px" }}>触发行数：</div>
            <input className={styles.input} style={{ width: "200px" }} value={this.state.textSegment.lines} />
            <div className={styles.text} style={{ marginLeft: "210px" }}>启用：</div>
            <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.textSegment.enable} />
          </div>
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>活动管理：</div>
            <Button type="primary" style={{ marginLeft: "390px" }}>新增</Button>
            <div className={styles.text} style={{ marginLeft: "25px" }}>全局开关：</div>
            <Switch checkedChildren="开" unCheckedChildren="关" defaultChecked />
          </div>
          <Table columns={activitiesColumns} dataSource={activitiesData} size="small" />
          <Divider className={styles.divide} />
          <div className={styles.inline}>
            <div className={styles.title}>操作日志：</div>
          </div>
          <div style={{ margin: "10px 40px 40px 40px" }}>
            <TextArea rows={6} />
          </div>
        </div>
      </div>
    );
  }
}

export default Config
