import { Divider, Button, Switch, Table, message, Card, Popconfirm } from 'antd';
import styles from './config.css';
import { Component } from 'react';
import { apiGet, apiAsyncGet, apiPost } from './utils/utils.js';
import router from 'umi/router'

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
      activities: {
        globalSwitch: true,
        activites: [],
      },
      logs: '',
    }
  }

  groupColumns = [
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
      title: '开关',
      dataIndex: 'switch',
      key: 'switch',
      render: (text, record, index) => (
        <span>
          <Switch checkedChildren="开" unCheckedChildren="关" checked={text} onChange={
            async (e) => {
              console.log(e);
              console.log('record', record);
              let ret = this.onSave('groupSwitch', { group: record.group, value: e, password: global.adminPassword });
              if (ret) {
                let group = Object.assign({}, this.state.group);
                for (let i = 0; i < group.groups.length; i++) {
                  if (group.groups[i].group === record.group) {
                    group.groups[i].switch = e;
                    break;
                  }
                }
                this.setState({ group });
              }
            }
          } />
        </span>
      ),
    },
    {
      title: '静默',
      dataIndex: 'silence',
      key: 'silence',
      render: (text, record, index) => (
        <span>
          <Switch checkedChildren="开" unCheckedChildren="关" checked={text} onChange={
            async (e) => {
              console.log(e);
              console.log('record', record);
              let ret = this.onSave('groupSilent', { group: record.group, value: e, password: global.adminPassword });
              if (ret) {
                let group = Object.assign({}, this.state.group);
                for (let i = 0; i < group.groups.length; i++) {
                  if (group.groups[i].group === record.group) {
                    group.groups[i].silence = e;
                    break;
                  }
                }
                this.setState({ group });
              }
            }
          } />
        </span>
      ),
    },
  ]

  activitiesColumns = [
    {
      title: '序号',
      dataIndex: 'key',
      key: 'key',
      width: 50,
    },
    {
      title: '关键字',
      dataIndex: 'keyWord',
      key: 'keyWord',
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
      render: (text, record) =>
        this.state.activities.activities.length >= 1 ? (
          <Popconfirm title="确认要删除?" onConfirm={async () => {
            console.log(text, record);
            let activities = Object.assign({}, this.state.activities);
            for (let i=0; i<activities.activities.length; i++) {
              if (activities.activities[i].keyWord === record.keyWord) {
                activities.activities.splice(i, 1);
                let ret = await apiPost('activities', activities);
                if (ret.data.data === true) {
                  message.success("删除成功");
                  this.setState({activities});
                  return;
                } else {
                  message.error('删除失败');
                  return;
                }
              }
            }
          }}>
            <a href="# ">删除</a>
          </Popconfirm>
        ) : null,
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

    apiAsyncGet('imageMode', (res) => {
      console.log(res.data.data);
      this.setState({ imageMode: res.data.data });
    });

    apiAsyncGet('textSegment', (res) => {
      console.log(res.data.data);
      this.setState({ textSegment: res.data.data });
    });

    apiAsyncGet('activities', (res) => {
      if (res.data.data.activities == null) {
        res.data.data.activities = [];
      }
      console.log(res.data.data);
      this.setState({ activities: res.data.data });
    });

  }

  appendLog = (msg) => {
    let log = this.state.logs;
    log += "\n" + (new Date()).toString() + "\n" + msg;
    this.setState({ logs: log });
  }

  onSaveSuperMaster = () => {
    this.onSave('supermaster', { qq: Number(this.state.supermaster), password: global.adminPassword });
  }

  onSaveDelay = async () => {
    this.onSave('supermaster', { delay: Number(this.state.delay), password: global.adminPassword });
  }

  onSave = async (path, data) => {
    console.log(data);
    let ret = await apiPost(path, data);
    if (ret.data.data === true) {
      message.success("保存成功");
      return true;
    }
    message.error("保存失败");
    console.log(ret);
    return false;
  }

  render() {
    return (
      <div className={styles.normal}>
        <div className={styles.body}>
          <Card style={{ width: "840px" }}>
            <div className={styles.msg}>更多设置选项，请查看使用手册，或私聊机器人发送“设置”查询和使用。</div>
            <div className={styles.inline}>
              <div className={styles.title}>插件主人（超级管理员）QQ：</div>
              <input className={styles.input} value={this.state.supermaster} onChange={e => this.setState({ supermaster: e.target.value })} />
              <Button type="primary" style={{ marginLeft: "40px" }} onClick={this.onSaveSuperMaster}>保存</Button>
            </div>
            <Divider className={styles.divide} />
            <div className={styles.inline}>
              <div className={styles.title}>消息回复延迟（毫秒）：</div>
              <input className={styles.input} style={{ marginLeft: "35px" }} value={this.state.delay} onChange={e => this.setState({ delay: e.target.value })} />
              <Button type="primary" style={{ marginLeft: "40px" }} onClick={this.onSaveDelay}>保存</Button>
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
              <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.group.globalSwitch} onChange={
                async (e) => {
                  console.log(e);
                  let ret = this.onSave('globalSwitch', { group: 0, value: e, password: global.adminPassword });
                  if (ret) {
                    let group = Object.assign({}, this.state.group);
                    group.globalSwitch = e;
                    this.setState({ group });
                  }
                }
              } />
              <div className={styles.text} style={{ marginLeft: "30px" }}>全局静默：</div>
              <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.group.globalSilence} onChange={
                async (e) => {
                  console.log(e);
                  let ret = this.onSave('globalSilent', { group: 0, value: e, password: global.adminPassword });
                  if (ret) {
                    let group = Object.assign({}, this.state.group);
                    group.globalSilence = e;
                    this.setState({ group });
                  }
                }
              } />
            </div>
            <Table columns={this.groupColumns} dataSource={this.state.group.groups} size="small" />
            <Divider className={styles.divide} />
            <div className={styles.inline}>
              <div className={styles.title}>图片模式：</div>
              <div className={styles.text} style={{ marginLeft: "40px" }}>触发行数：</div>
              <input className={styles.input} style={{ width: "200px" }} value={this.state.imageMode.lines} onChange={e => {
                let imageMode = Object.assign({}, this.state.imageMode);
                imageMode.lines = e.target.value;
                this.setState({ imageMode });
              }
              } />
              <div className={styles.text} style={{ marginLeft: "210px" }}>启用：</div>
              <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.imageMode.enable} onChange={e => {
                let ret = this.onSave('imageMode', { enable: e, lines: Number(this.state.imageMode.lines) });
                if (ret) {
                  let imageMode = Object.assign({}, this.state.imageMode);
                  imageMode.enable = e;
                  this.setState({ imageMode });
                }
              }
              } />
            </div>

            <Divider className={styles.divide} />
            <div className={styles.inline}>
              <div className={styles.title}>文字分段：</div>
              <div className={styles.text} style={{ marginLeft: "40px" }}>触发行数：</div>
              <input className={styles.input} style={{ width: "200px" }} value={this.state.textSegment.lines} onChange={e => {
                let textSegment = Object.assign({}, this.state.textSegment);
                textSegment.lines = e.target.value;
                this.setState({ textSegment });
              }
              } />
              <div className={styles.text} style={{ marginLeft: "210px" }}>启用：</div>
              <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.textSegment.enable} onChange={e => {
                let ret = this.onSave('textSegment', { enable: e, lines: Number(this.state.textSegment.lines) });
                if (ret) {
                  let textSegment = Object.assign({}, this.state.textSegment);
                  textSegment.enable = e;
                  this.setState({ textSegment });
                }
              }
              } />
            </div>
            <Divider className={styles.divide} />
            <div className={styles.inline}>
              <div className={styles.title}>活动管理：</div>
              <Button type="primary" style={{ marginLeft: "390px" }} onClick={() => {
                global.activities = Object.assign({}, this.state.activities);
                router.push("/newAct");
              }
              }>新增</Button>
              <div className={styles.text} style={{ marginLeft: "25px" }}>全局开关：</div>
              <Switch checkedChildren="开" unCheckedChildren="关" checked={this.state.activities.globalSwitch} />
            </div>
            <Table columns={this.activitiesColumns} dataSource={this.state.activities.activities} size="small" />
          </Card>
        </div>
      </div>
    );
  }
}

export default Config
