
import styles from './login.css';
import { Component } from 'react';
import { Input, Button } from 'antd';
class Login extends Component {
  render() {
    return (
      <div className={styles.normal}>
        <div>
          <h3>请输入管理员密码</h3>
          <input className={styles.input} placeholder="密码" type="password"/>
          <br/>
          <Button type="primary" className={styles.button}>登入</Button>
        </div>
      </div>
    );
  }
}

export default Login;