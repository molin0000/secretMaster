
import styles from './register.css';
import { Component } from 'react';
import { Input, Button } from 'antd';
class Register extends Component {
  render() {
    return (
      <div className={styles.normal}>
        <div>
          <h3>首次使用请设置管理员密码</h3>
          <Input className={styles.input} placeholder="密码" type="password"/>
          <br/>
          <Input className={styles.input} placeholder="密码确认" type="password"/>
          <br/>
          <Button type="primary" className={styles.button}>确定</Button>
        </div>
      </div>
    );
  }
}

export default Register;