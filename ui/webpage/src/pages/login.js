
import styles from './login.css';
import { Component } from 'react';
import { Button, message, Modal } from 'antd';
import { apiPost, apiGet } from './utils/utils';
import router from 'umi/router';
class Login extends Component {
  constructor(props) {
    super(props);

    this.state = {
      password: ""
    }
  }

  async componentWillMount() {
    let locked = await apiGet('locked');
    if (!locked.data.data) {
      router.push("/register");
      return;
    }
  }

  onOk = async () => {
    if (this.state.password.length === 0) {
      message.error("密码不能为空！");
      return;
    }

    let ret = await apiPost('password', { qq: 0, password: this.state.password });
    if (ret.data.data !== true) {
      message.error("密码错误!");
      return;
    }

    global.unlocked = true;
    message.success("登录成功！");
    global.adminPassword = this.state.password;
    router.push('/config');
  }

  render() {
    return (
      <div className={styles.normal}>
        <div>
          <h3>请输入管理员密码</h3>
          <input className={styles.input} placeholder="密码" type="password" value={this.state.password} onChange={(e) => { this.setState({ password: e.target.value }) }} onKeyDown={e => { if (e.keyCode === 13) this.onOk() }} />
          <br />
          <a href="# "
            onClick={() => {
              Modal.success({
                content: '使用超级管理员QQ私聊机器人发送：GM;clearpassword;0;0 可清除密码',
              });
            }}>忘记密码?</a>
          <br />
          <Button type="primary" className={styles.button} onClick={this.onOk}>登入</Button>
        </div>
      </div>
    );
  }
}

export default Login;