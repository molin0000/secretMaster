import { Card, Input, Button, message, Row, Modal } from 'antd';
import styles from './qqlogin.css';
import router from 'umi/router'
import { Component } from 'react';
import { apiPost } from './utils/utils.js';

class QQLogin extends Component {
  constructor(props) {
    super(props)
    this.state = {
      qq: '',
      password: '',
    }
  }

  onOk = async () => {
    if (this.state.qq.length === 0 || this.state.password.length === 0) {
      message.error("QQ号码和口令都不能为空");
      return;
    }

    let ret = await apiPost('password', { qq: Number(this.state.qq), password: this.state.password });
    if (ret.data.data !== true) {
      message.error("口令错误！");
      return;
    }

    message.success("登陆成功！");
    global.qq = this.state.qq;
    global.password = this.state.password;
    router.push('/');
  }

  render() {
    return (
      <div className={styles.normal}>
        <Card title="QQ登录" style={{ maxWidth: "600px" }}>
          <Row>
            <Input placeholder={"请输入QQ号码"} style={{ width: "200px", textAlign: "center", margin: "10px" }} value={this.state.qq} onChange={(e) => { this.setState({ qq: e.target.value }); }} />
          </Row>
          <Row>
            <Input placeholder={"请输入口令(不是QQ密码！)"} style={{ width: "200px", textAlign: "center", margin: "10px" }} value={this.state.password} onChange={(e) => { this.setState({ password: e.target.value }); }} onKeyDown={e => { if (e.keyCode === 13) this.onOk() }} />
          </Row>
          <Row>
            <Button type='primary' onClick={this.onOk} style={{marginTop:"20px", marginBottom:"10px"}}>登录</Button>

          </Row>
          <Row>
            <a href="# "
            onClick={() => {
              Modal.success({
                content: '使用QQ号码登录来绑定游玩账号,请注意！！口令不是你的QQ密码！是通过私聊机器人配置的口令。必须先设定口令才能登陆。随时可以私聊机器人修改口令，配置指令格式为：口令;新口令，例如：口令;cat。需要绑定私聊QQ群，绑定方法为：私聊机器人发送：序列战争。更换绑定群请私聊机器人发送一个@，网页版目前不支持绑定群聊功能，需要私聊机器人完成绑定',
              });}}>登录说明?</a>
          </Row>
          {/* */}
        </Card>
      </div>
    );
  }
}

export default QQLogin;

