import styles from './index.css';
import { Component } from 'react';
import { Button, Statistic, Row, Col, Input, Divider, Card } from 'antd';
import router from 'umi/router'
import { apiAsyncPost } from './utils/utils.js';

const { TextArea } = Input;

class Home extends Component {
  constructor(props) {
    super(props)
    this.state = {
      qq: 0,
      password: "",
      msg: "",
      reply: "",
    }
  }

  formatNumber = n => {
    n = n.toString()
    return n[1] ? n : '0' + n
  }
  // 时间格式化
  formatTime = date => {
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = date.getHours()
    const minute = date.getMinutes()
    const second = date.getSeconds()

    return [year, month, day].map(this.formatNumber).join('-') + ' ' + [hour, minute, second].map(this.formatNumber).join(':')
  }

  componentWillMount(props) {
    if (!global.qq) {
      router.push('/qqlogin');
      return;
    }

    this.setState({ qq: global.qq, password: global.password });
  }

  sendMsg = (msg) => {
    let data = {
      qq: Number(this.state.qq),
      password: this.state.password,
      msg
    }
    console.log("msg:", data);
    apiAsyncPost('chat', data, (res) => {
      console.log(res);
      let info = this.state.reply;
      this.setState({ reply: info + "\n---------" + this.formatTime(new Date()) + "----------\n" + res.data.data.Msg });
      var ta = document.getElementById('textArea');
      ta.scrollTop = ta.scrollHeight;
    });
  }

  render() {
    return (
      <div className={styles.body}>
        <Card style={{ maxWidth: "600px" }}>
          <Row gutter={16}>
            <Row>
              <Col span={12}>
                <Button type='primary' style={{ width: '20vw', marginTop: '10px', maxWidth: "100px" }}
                  onClick={() => { router.push('/qqlogin'); }}>QQ登录</Button>
              </Col>
              <Col span={12}>
                {/* <Statistic title="当前QQ号码" value={this.state.qq} groupSeparator={""} /> */}
                <Row style={{color:"white", padding:"15px"}}>
                  <Col span={12}>当前QQ号码:</Col>
                  <Col span={12}>{this.state.qq}</Col>
                </Row>

              </Col>
              {/* <Col span={9}>
              <Statistic title="昵称" value={"空想之喵"} />
            </Col> */}
            </Row>
            <Row>
              <Col span={1}></Col>
              <Col span={22}>
                <TextArea id="textArea" rows={16} style={{ marginTop: "10px", marginBottom: "10px", background: "#AE8AFF", borderRadius: "20px", color: "white" }} readOnly={true} value={this.state.reply} />
              </Col>
              <Col span={1}></Col>
            </Row>
            <Row>
              <Row gutter={16}>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('探险') }}>探险</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('钓鱼') }}>钓鱼</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('副本') }}>副本</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('许愿') }}>许愿</Button></Col>
              </Row>
              <Row gutter={16}>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('速算') }}>速算</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('学识') }}>学识</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('阵营') }}>阵营</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('道具') }}>道具</Button></Col>
              </Row>
              <Row gutter={16}>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('祈祷') }}>祈祷</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('商店') }}>商店</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('属性') }}>属性</Button></Col>
                <Col span={6}><Button type='primary' style={{ marginBottom: "10px", width: "20vw", maxWidth: "100px" }} onClick={() => { this.sendMsg('排行') }}>排行</Button></Col>
              </Row>
            </Row>
            <Divider style={{ width: "80%" }} />
            <Row>
              <Col span={2}></Col>
              <Col span={16}><Input value={this.state.msg} onChange={(e) => { this.setState({ msg: e.target.value }) }} onKeyDown={e => { if (e.keyCode === 13) { this.sendMsg(this.state.msg); this.setState({ msg: "" }) } }} /></Col>
              <Col span={4}><Button type='primary' style={{ width: "20vw", marginLeft: "10px", maxWidth: "100px" }} onClick={() => { this.sendMsg(this.state.msg); this.setState({ msg: "" }) }}>发送</Button></Col>
            </Row>
          </Row>
        </Card>
      </div>
    );
  }
}

export default Home;
