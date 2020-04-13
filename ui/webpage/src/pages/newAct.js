import { Card, Button, Input, InputNumber, Select, message } from 'antd';
import styles from './newAct.css';
import { Component } from 'react';
import router from 'umi/router';
import { apiPost } from './utils/utils.js';

const { TextArea } = Input;
const { Option } = Select;


class NewAct extends Component {
  constructor(props) {
    super(props);

    this.state = {
      keyWord: "签到",
      reply: "签到成功",
      type: "每日",
      exp: 0,
      money: 0,
      magic: 0
    }
  }

  onSave = async ()=>{
    let data = global.activities;

    for(let i=0; i<data.activities.length; i++) {
      if (data.activities[i].keyWord === this.state.keyWord) {
        message.error("关键字重复！")
        return;
      }
    }

    let key = data.activities.length > 0 ? Number(data.activities[data.activities.length-1].key) + 1 : 0;
    let value = {
      key, 
      keyWord: this.state.keyWord, 
      reply: this.state.reply,
      type: this.state.type,
      exp: Number(this.state.exp),
      money: Number(this.state.money),
      magic: Number(this.state.magic),
    };

    data.activities.push(value);
    console.log('data', data);

    let ret = await apiPost('activities', data);
    console.log('ret', ret.data.data);
    if (ret.data.data === true) {
      message.success("保存成功");
      router.push("/config")
    } else {
      message.error("保存失败");
    }
  }

  render() {
    return (
      <div className={styles.normal}>
        <Card title="新建活动事件" style={{ maxWidth: "600px" }}>
          <div className={styles.title}>关键字：</div>
          <Input className={styles.input} value={this.state.keyWord} onChange={e => this.setState({ keyWord: e.target.value })} />
          <div className={styles.title}>回复：</div>
          <TextArea row={3} className={styles.input} value={this.state.reply} onChange={e => this.setState({ reply: e.target.value })} />
          <div className={styles.title}>类型：</div>
          <Select className={styles.input} defaultValue={"每天"} value={this.state.type} onChange={e => this.setState({ type: e })} >
            <Option value="每天">每天</Option>
            <Option value="每周">每周</Option>
            <Option value="每月">每月</Option>
            <Option value="每年">每年</Option>
            <Option value="单次">单次</Option>
            <Option value="每人一次">每人一次</Option>
          </Select>
          <div className={styles.title}>金镑：</div>
          <InputNumber className={styles.input} value={this.state.money} onChange={e => this.setState({ money: e })} min={0}/>
          <div className={styles.title}>经验：</div>
          <InputNumber className={styles.input} value={this.state.exp} onChange={e => this.setState({ exp: e })} min={0}/>
          <div className={styles.title}>灵力：</div>
          <InputNumber className={styles.input} value={this.state.magic} onChange={e => this.setState({ magic: e })} min={0}/>
          <br />
          <Button type='primary' style={{ margin: "20px 5px 5px 5px" }} onClick={() => { this.onSave(); }}>确定</Button>
          <Button type='primary' style={{ margin: "20px 5px 5px 5px" }} onClick={() => { router.push("/config") }}>取消</Button>
        </Card>
      </div>
    );
  }
}

export default NewAct;
