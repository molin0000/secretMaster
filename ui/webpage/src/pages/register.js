
import styles from './register.css';
import { Component } from 'react';
import { Button, message } from 'antd';
import {apiPost} from './utils/utils.js';
import router from 'umi/router'


class Register extends Component {
  constructor(props) {
    super(props);

    this.state = {
      password: '',
      passwordConfirm: '',
    }
  }

  onChangeP1 = (e) => {
    this.setState({password: e.target.value})
  }

  onChangeP2 = (e) => {
    this.setState({passwordConfirm: e.target.value})
  }

  onOk = async ()=>{
    if (this.state.password !== this.state.passwordConfirm) {
      message.error("两次输入的密码不一致！");
      return;
    }

    if (this.state.password.length === 0) {
      message.error("密码不能为空！");
      return;
    }

    let ret = await apiPost('password', {
      qq: 0,
      password: this.state.password
    });
    message.success("密码设置成功！");
    router.push('/config');
  }

  render() {
    return (
      <div className={styles.normal}>
        <div>
          <h3>首次使用请设置管理员密码</h3>
          <input className={styles.input} placeholder="密码" type="password" onChange={this.onChangeP1} value={this.state.password}/>
          <br/>
          <input className={styles.input} placeholder="密码确认" type="password" onChange={this.onChangeP2} value={this.state.passwordConfirm} onKeyDown={e => { if (e.keyCode === 13) this.onOk() }}/>
          <br/>
          <Button type="primary" className={styles.button} onClick={this.onOk}>确定</Button>
        </div>
      </div>
    );
  }
}

export default Register;