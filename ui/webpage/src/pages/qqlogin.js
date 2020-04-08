import { Card, Input, Button } from 'antd';
import styles from './qqlogin.css';
import router from 'umi/router'

export default function () {
  return (
    <div className={styles.normal}>
      <Card title="QQ登录">
        <p>您可以使用QQ号码登录来绑定游玩账号</p>
        <p>请注意！！这里的口令不是你的QQ密码！！</p>
        <p>而是通过私聊机器人配置的独立口令</p>
        <p>默认口令为空，可输入QQ号码直接登录</p>
        <p>随时可以私聊机器人修改口令</p>
        <p>配置指令格式为：口令;新口令</p>
        <p>例如：口令;cat</p>
        <Input placeholder={"请输入QQ号码"} style={{width:"200px", textAlign:"center", margin:"10px"}}/>
        <Input placeholder={"请输入口令，默认为空"} style={{width:"200px", textAlign:"center", margin:"10px"}}/>
        <Button type='primary' onClick={()=>{router.push("/?qq=67939461")}}>登录</Button>
      </Card>
    </div>
  );
}

